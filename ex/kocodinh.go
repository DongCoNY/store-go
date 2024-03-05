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

func check3() {
	var fromSeq uint32
	fromSeq = 439540

	ctx := context.Background()
	// Only log errors from the backend to keep output cleaner.
	lg := log.New()
	lg.SetLevel(logrus.ErrorLevel)
	config.Log = lg

	backend, err := backends.NewCaptive(config)

	fmt.Println("port:", config.Toml.HTTPPort)

	panicIf(err)
	defer backend.Close()

	fmt.Printf("From seq: %d\n", fromSeq)

	ledgerRange := backends.UnboundedRange(fromSeq)
	err = backend.PrepareRange(ctx, ledgerRange)
	panicIf(err)
	fmt.Printf("From seq: %d\n", fromSeq)
	seq := fromSeq - 1
	for {
		seq++
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
				fmt.Println("=====")
				// fmt.Println(tx.)
				fmt.Println(tx.Result.TransactionHash.HexString())
				// fmt.Println(tx)

				fmt.Println("=====")
			}

		}
	}
}
