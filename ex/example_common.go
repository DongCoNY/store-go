package tutorial

import (
	"fmt"

	"github.com/stellar/go/ingest/ledgerbackend"
)

var (
	config = captiveCoreConfig()
)

func captiveCoreConfig() ledgerbackend.CaptiveCoreConfig {
	archiveURLs := []string{
		// "https://soroban-testnet.stellar.org:443",
		// "https://history-testnet.stellar.org/prd/core-testnet/core_testnet_001",
		// "http://history.stellar.org/prd/core-live/core_live_002/",
		// "https://history-testnet.stellar.org",
		// "https://horizon-testnet.stellar.org/",
		// "https://futurenet.sorobandev.com/soroban/rpc",
		// "https://friendbot.stellar.org",

		"https://history.stellar.org/prd/core-testnet/core_testnet_001",
		"https://history.stellar.org/prd/core-testnet/core_testnet_002",
		"https://history.stellar.org/prd/core-testnet/core_testnet_003",
	}
	// networkPassphrase := "Test SDF Network ; September 2015"
	// networkPassphrase := "Public Global Stellar Network ; September 2015"
	networkPassphrase := "Test SDF Network ; September 2015"
	// networkPassphrase := "Test SDF Future Network ; October 2022"

	captiveCoreToml, err := ledgerbackend.NewCaptiveCoreToml(ledgerbackend.CaptiveCoreTomlParams{
		NetworkPassphrase:  networkPassphrase,
		HistoryArchiveURLs: archiveURLs,
	})
	panicIf(err)

	captiveCoreToml, err = captiveCoreToml.CatchupToml()
	panicIf(err)

	return ledgerbackend.CaptiveCoreConfig{
		// Change these based on your environment:
		BinaryPath:         "/usr/local/bin/stellar-core",
		NetworkPassphrase:  networkPassphrase,
		HistoryArchiveURLs: archiveURLs,
		Toml:               captiveCoreToml,
	}
}

func panicIf(err error) {
	if err != nil {
		panic(fmt.Errorf("an error occurred, panicking: %s", err))
	}
}
