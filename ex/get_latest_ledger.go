package tutorial

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func GetLatestLedger() (uint32, error) {

	// Dữ liệu JSON để gửi trong yêu cầu POST
	requestData := map[string]interface{}{
		"jsonrpc": "2.0",
		"id":      8675309,
		"method":  "getLatestLedger",
	}

	// Chuyển đổi dữ liệu JSON sang dạng []byte
	requestBody, err := json.Marshal(requestData)
	if err != nil {
		fmt.Println("Lỗi khi chuyển đổi dữ liệu JSON:", err)
		return 0, err
	}

	// Gửi yêu cầu POST đến URL
	response, err := http.Post("https://soroban-testnet.stellar.org", "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		fmt.Println("Lỗi khi gửi yêu cầu POST:", err)
		return 0, err
	}
	defer response.Body.Close()

	// // Decode the JSON response into a custom struct
	// var responseData ResponseData
	// if err := json.NewDecoder(response.Body).Decode(&responseData); err != nil {
	// 	fmt.Println("Error decoding JSON response:", err)
	// 	return nil, err
	// }

	// Đọc và hiển thị phản hồi JSON
	var responseData map[string]interface{}
	if err := json.NewDecoder(response.Body).Decode(&responseData); err != nil {
		fmt.Println("Lỗi khi đọc phản hồi JSON:", err)
		return 0, err
	}

	// Lấy giá trị của "sequence" từ "result"
	result := responseData["result"].(map[string]interface{})
	sequence := result["sequence"].(float64) // Vì "sequence" là số nên chúng ta sử dụng float64

	latestLedger := uint32(sequence)
	return latestLedger, nil
}
