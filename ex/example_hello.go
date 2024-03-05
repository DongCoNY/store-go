package tutorial

import (
	"context"
	"fmt"

	"github.com/stellar/go/ingest"
	backends "github.com/stellar/go/ingest/ledgerbackend"
)

func helloworld() {
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

	// fmt.Println(txReader.CountTransactions)

	tx, err := txReader.Read()
	panicIf(err)
	fmt.Println(tx.Result.Successful())
	// Now `ledger` is a raw `xdr.LedgerCloseMeta` object containing the
	// transactions contained within this ledger.
	fmt.Printf("\nSequence: %d.\n", ledger.LedgerSequence())
}
