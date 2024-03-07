package tutorial

import (
	"context"
	"fmt"
	"io"

	"github.com/sirupsen/logrus"
	"github.com/stellar/go/ingest"
	backends "github.com/stellar/go/ingest/ledgerbackend"
	"github.com/stellar/go/support/log"
)

func check4() {
	var fromSeq uint32
	fromSeq = 439540

	ctx := context.Background()
	// Only log errors from the backend to keep output cleaner.
	lg := log.New()
	lg.SetLevel(logrus.ErrorLevel)
	config.Log = lg

	backend, err := backends.NewCaptive(config)

	panicIf(err)
	defer backend.Close()

	for {
		ledgerRange := backends.BoundedRange(fromSeq, fromSeq+5)
		err = backend.PrepareRange(ctx, ledgerRange)
		if err == io.EOF {
			break
		}
		panicIf(err)
		for seq := fromSeq; seq <= fromSeq+5; seq++ {
			fmt.Printf("Processed ledger %d...\r", seq)

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
					// fmt.Println("=====")
					// fmt.Println(tx.)

					fmt.Println(tx.Result.TransactionHash.HexString())
					fmt.Printf("%v\n", tx)

					// fmt.Println(tx)

					// fmt.Println("=====")
				}
			}
		}
		fromSeq += 5
		fmt.Println("=======")
	}

}
