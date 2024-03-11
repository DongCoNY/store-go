package tutorial

import (
	"context"
	"fmt"
	"io"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/stellar/go/ingest"
	backends "github.com/stellar/go/ingest/ledgerbackend"
	"github.com/stellar/go/support/log"
)

func example() {
	var fromSeq, step uint32
	// 485951 to 486079, 486143, 486207
	fromSeq = 556898
	step = 64

	ctx := context.Background()
	// Only log errors from the backend to keep output cleaner.
	lg := log.New()
	lg.SetLevel(logrus.ErrorLevel)
	config.Log = lg

	backend, err := backends.NewCaptive(config)

	panicIf(err)
	defer backend.Close()

	for {
		ledgerRange := backends.BoundedRange(fromSeq, fromSeq+step)
		err = backend.PrepareRange(ctx, ledgerRange)
		if err != nil {
			if strings.Contains(err.Error(), "is greater than max available in history archives") {
				panic(err)
			}
			panic(err.Error())
		}

		for seq := fromSeq; seq <= fromSeq+step; seq++ {
			fmt.Printf("==========Processed ledger %d========\n", seq)

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
				panicIf(err)
				if tx.Result.Successful() {
					fmt.Println("txHash:", tx.Result.TransactionHash.HexString())
					fmt.Println(tx.FeeChanges[0].State.Data.Account.Balance)
					// fmt.Printf("%v\n", tx)
				}
			}
		}
		fromSeq += step + 1
	}

}
