package tutorial

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func GetEventsFromSeq(seqLedgerFrom uint32) {
	url := "https://soroban-testnet.stellar.org"
	requestBody, err := json.Marshal(map[string]interface{}{
		"jsonrpc": "2.0",
		"id":      8675309,
		"method":  "getEvents",
		"params": map[string]interface{}{
			"startLedger": seqLedgerFrom,
			"pagination": map[string]interface{}{
				"limit": 10000,
			},
		},
	})
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		fmt.Println("Error making HTTP request:", err)
		return
	}
	defer resp.Body.Close()

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		fmt.Println("Error decoding JSON response:", err)
		return
	}

	prettyJSON, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("Error formatting JSON:", err)
		return
	}

	fmt.Println(string(prettyJSON))
}
