package tutorial

import (
	"context"
	"fmt"
	"io"

	"github.com/stellar/go/ingest"
	backends "github.com/stellar/go/ingest/ledgerbackend"
)

func check2() {
	seq := uint32(435801)
	ctx := context.Background()
	backend, err := backends.NewCaptive(config)
	panicIf(err)
	defer backend.Close()

	// Prepare a single ledger to be ingested,
	err = backend.PrepareRange(ctx, backends.BoundedRange(seq, seq))
	panicIf(err)

	// then retrieve it:434400
	ledger, err := backend.GetLedger(ctx, seq)
	panicIf(err)

	txReader, err := ingest.NewLedgerTransactionReader(
		ctx, backend, config.NetworkPassphrase, seq,
	)
	panicIf(err)
	defer txReader.Close()

	for {
		tx, err := txReader.Read()
		if err == io.EOF {
			break
		}
		panicIf(err)

		fmt.Println(tx.Result.Successful())
		if tx.Result.Successful() {
			fmt.Println("=====")
			// fmt.Println(tx.)
			fmt.Println(tx.Result.TransactionHash.HexString())
			// fmt.Println(tx)

			fmt.Println("=====")
		}
	}
	// Now `ledger` is a raw `xdr.LedgerCloseMeta` object containing the
	// transactions contained within this ledger.
	fmt.Printf("\nSequence: %d.\n", ledger.LedgerSequence())
}
