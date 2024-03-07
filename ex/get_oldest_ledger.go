package tutorial

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// save 1439 ledger
func GetOldeastLedger() (uint32, error) {
	url := "https://soroban-testnet.stellar.org"
	requestBody, err := json.Marshal(map[string]interface{}{
		"jsonrpc": "2.0",
		"id":      8675309,
		"method":  "getTransaction",
		"params": map[string]interface{}{
			"hash": "54b6e9f1352c8b3014ee71f6acdeef8b8b78aab0a69eb89d512a25dec0713301", // tx of ledger old
		},
	})
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return 0, err
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		fmt.Println("Error making HTTP request:", err)
		return 0, err
	}
	defer resp.Body.Close()

	var responseData map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&responseData); err != nil {
		fmt.Println("Error decoding JSON response:", err)
		return 0, err
	}

	result := responseData["result"].(map[string]interface{})
	sequence := result["oldestLedger"].(float64)

	oldeastLedger := uint32(sequence)

	return oldeastLedger, nil
}
