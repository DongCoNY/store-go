package tutorial

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type InfoTx struct {
	ApplicationOrder      string
	CreatedAt             string
	EnvelopeXdr           string
	LatestLedger          string
	LatestLedgerCloseTime string
	Ledger                string
	OldestLedger          string
	OldestLedgerCloseTime string
	ResultMetaXdr         string
	ResultXdr             string
	Status                string
}

// save 1439 ledger
// hex.EncodeToString
func GetInfoTxFromTxHash(txHash string) (*InfoTx, error) {

	url := "https://soroban-testnet.stellar.org"
	requestBody, err := json.Marshal(map[string]interface{}{
		"jsonrpc": "2.0",
		"id":      8675309,
		"method":  "getTransaction",
		"params": map[string]interface{}{
			"hash": txHash,
		},
	})
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return nil, err
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		fmt.Println("Error making HTTP request:", err)
		return nil, err
	}
	defer resp.Body.Close()

	var responseData map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&responseData); err != nil {
		fmt.Println("Error decoding JSON response:", err)
		return nil, nil
	}
	var infoTx InfoTx
	result := responseData["result"].(map[string]interface{})

	if result["applicationOrder"] == nil {
		// fmt.Println(responseData)
		infoTx.OldestLedgerCloseTime = result["oldestLedgerCloseTime"].(string)
		// infoTx.OldestLedger = result["oldestLedger"].(string)
		infoTx.Status = result["status"].(string)
		infoTx.LatestLedgerCloseTime = result["latestLedgerCloseTime"].(string)
		// infoTx.LatestLedger = result["latestLedger"].(string)
		return &infoTx, nil
	}

	infoTx.ApplicationOrder = result["applicationOrder"].(string)
	infoTx.CreatedAt = result["createdAt"].(string)
	infoTx.EnvelopeXdr = result["envelopeXdr"].(string)
	infoTx.LatestLedger = result["latestLedger"].(string)
	infoTx.LatestLedgerCloseTime = result["latestLedgerCloseTime"].(string)
	infoTx.Ledger = result["ledger"].(string)
	infoTx.OldestLedger = result["oldestLedger"].(string)
	infoTx.OldestLedgerCloseTime = result["oldestLedgerCloseTime"].(string)
	infoTx.ResultMetaXdr = result["resultMetaXdr"].(string)
	infoTx.ResultXdr = result["resultXdr"].(string)
	infoTx.Status = result["status"].(string)

	return &infoTx, nil
}
