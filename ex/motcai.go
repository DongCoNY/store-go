package tutorial

import (
	"context"
	"fmt"
	"io"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/stellar/go/ingest"
	backends "github.com/stellar/go/ingest/ledgerbackend"
	"github.com/stellar/go/support/log"
)

const (
	step = 64
)

// 555391, 555327
var fromSeq = uint32(557400)

type txInfo struct {
	txHash string
}

type Aggregation struct {
	txQueue chan txInfo
	data    []txInfo
	index   int64
}

func start() error {
	as := Aggregation{
		txQueue: make(chan txInfo),
		data:    []txInfo{},
		index:   0,
	}

	// Only log errors from the backend to keep output cleaner.
	lg := log.New()
	lg.SetLevel(logrus.ErrorLevel)
	config.Log = lg
	ctx := context.Background()

	backend, err := backends.NewCaptive(config)
	panicIf(err)
	defer backend.Close()
	for {

		ledgerRange := backends.BoundedRange(fromSeq, fromSeq+step)
		err = backend.PrepareRange(ctx, ledgerRange)
		if err != nil {
			//"is greater than max available in history archives"
			err = pauseWaitLedger(err)
			if err != nil {
				return err
			}
			continue
		}

		for seq := fromSeq; seq < fromSeq+step; seq++ {
			fmt.Println(seq)
			txReader, err := ingest.NewLedgerTransactionReader(
				ctx, backend, config.NetworkPassphrase, seq,
			)
			panicIf(err)
			defer txReader.Close()

			// Read each transaction within the ledger, extract its operations, and
			// accumulate the statistics we're interested in.
			for {
				tx, err := txReader.Read()
				if err == io.EOF {
					break
				}

				if err != nil {
					return err
				}

				if tx.Result.Successful() {
					newTxInfo := txInfo{
						txHash: tx.Result.TransactionHash.HexString(),
					}

					go func(tx txInfo) {
						// add txInfo chan txQueue <- tx
						as.txQueue <- tx
					}(newTxInfo)
				}
			}
			go as.process()
		}

		fromSeq += step
	}
}

func (as *Aggregation) process() {
	// fmt.Println("index:", as.index)
	as.index++
	tx := <-as.txQueue
	// fmt.Println(tx)
	as.data = append(as.data, tx)
	// fmt.Println("=====")
	// fmt.Println(as.GetData())
}

func (as *Aggregation) GetData() []txInfo {
	return as.data
}

// to limit computational resources
func pauseWaitLedger(err error) error {
	if !strings.Contains(err.Error(), "is greater than max available in history archives") {
		// if not err by LatestLedger: xxx is greater than max available in history archives yyy
		return err
	}

	re := regexp.MustCompile(`(\d+)`)
	matches := re.FindAllString(err.Error(), -1)
	seqHistoryArchives, err := strconv.Atoi(matches[1])

	if err != nil {
		return err
	}
	estimateSeqNext := int64(seqHistoryArchives) + 64

	latestLedger, err := GetLatestLedger()
	if err != nil {
		return err
	}

	numLedgerWait := estimateSeqNext - int64(latestLedger) + 1

	if numLedgerWait < 0 {
		fmt.Println("amm")
		return nil
	}
	// Ledger closing time is ~4s/ledger
	ledgerClosingTime := 4 * time.Second
	estimateTimeWait := numLedgerWait * ledgerClosingTime.Nanoseconds()

	fmt.Println(" Waiting....", time.Duration(estimateTimeWait))
	time.Sleep(time.Duration(estimateTimeWait))
	return nil
}
