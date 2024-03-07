package tutorial

import (
	"context"
	"fmt"
	"io"

	"github.com/stellar/go/ingest"
	backends "github.com/stellar/go/ingest/ledgerbackend"
)

func check2() {
	seq := uint32(454016)
	// seq := uint32(453952)
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

	a, err := backend.GetLatestLedgerSequence(ctx)
	panicIf(err)
	fmt.Println("xxxx", a)

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
			txHash := tx.Result.TransactionHash.HexString()
			inSuccessfulContractCalls, err := tx.GetDiagnosticEvents()

			panicIf(err)
			// txReader.GetHeader().Header.PreviousLedgerHash

			// GetDiagnosticEvents()
			fmt.Println("=====")
			fmt.Println(txReader.GetHeader().Header.IdPool)
			fmt.Println(tx.LedgerVersion)
			fmt.Println(tx.FeeChanges[0].Type)
			fmt.Println(tx.FeeChanges[0].State.Data)
			fmt.Println(tx.Envelope.Type)
			fmt.Println(inSuccessfulContractCalls)
			// fmt.Println(inSuccessfulContractCalls[0].Event.ContractId.HexString())
			// fmt.Println(inSuccessfulContractCalls[0].InSuccessfulContractCall)
			// fmt.Println(tx.)
			fmt.Println(txHash)
			// fmt.Println(tx)

			fmt.Println("=====")
		}
	}
	// Now `ledger` is a raw `xdr.LedgerCloseMeta` object containing the
	// transactions contained within this ledger.
	fmt.Printf("\nSequence: %d.\n", ledger.LedgerSequence())
}
