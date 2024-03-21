package tutorial

import (
	"context"
	"fmt"
	"io"

	"github.com/stellar/go/historyarchive"
	"github.com/stellar/go/ingest"
	// "github.com/stellar/go/ingest/ledgerbackend"
	// "github.com/stellar/go/network"
	"github.com/stellar/go/support/storage"
	// "github.com/stellar/go/xdr"
)

func checkContractDataEntry() {
	// archiveURL := "http://history.stellar.org/prd/core-live/core_live_001"

	archiveURL := "https://history.stellar.org/prd/core-testnet/core_testnet_001"

	archive, err := historyarchive.Connect(
		archiveURL,
		historyarchive.ArchiveOptions{
			ConnectOptions: storage.ConnectOptions{
				Context: context.TODO(),
			},
		},
	)
	if err != nil {
		panic(err)
	}

	// Ledger must be a checkpoint ledger: (100031+1) mod 64 == 0.
	// 715839, 715903
	reader, err := ingest.NewCheckpointChangeReader(context.TODO(), archive, 718463)
	if err != nil {
		panic(err)
	}

	for {
		entry, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}

		if entry.Type == 6 {
			fmt.Println("=====")
			fmt.Printf("Ext: %v \n", entry.Post.Data.ContractData.Ext)
			fmt.Printf("Contract: %v \n", entry.Post.Data.ContractData.Contract)
			fmt.Printf("\t ContractID: %v \n", entry.Post.Data.ContractData.Contract.ContractId.HexString())
			fmt.Printf("Key: %v \n", entry.Post.Data.ContractData.Key)
			fmt.Printf("Durability: %v \n", entry.Post.Data.ContractData.Durability)
			fmt.Printf("Val: %v \n", entry.Post.Data.ContractData.Val)
		}
	}

}
