package main

import (
	// "encoding/hex"
	// "fmt"
	// "regexp"
	ma "github.com/DongCoNY/store-go/ex"
	// "time"
)

// type mm struct {
// 	t1 string
// }
// type aa struct {
// 	tt chan mm
// }

// func (a *aa) loopp() {
// 	// time.Sleep(2 * time.Second)
// 	aaa := <-a.tt
// 	fmt.Println("---", aaa)
// }

func main() {
	ma.Main()
	// err := "error starting prepare range: opening subprocess: to sequence: 557528 is greater than max available in history archives: 557503"
	// re := regexp.MustCompile(`(\d+)`)

	// matches := re.FindAllString(err, -1)

	// if len(matches) >= 2 {
	// 	// Lấy giá trị của a và b từ các phần khớp tìm được
	// 	a := matches[len(matches)-2]
	// 	b := matches[len(matches)-1]
	// 	fmt.Println("a:", a)
	// 	fmt.Println("b:", b)
	// } else {
	// 	fmt.Println("Không tìm thấy đủ số nguyên trong chuỗi.")
	// }
	// fmt.Println("========")
	// // start := time.Now()
	// a := []mm{
	// 	{t1: "11111"},
	// 	{t1: "22222"},
	// 	{t1: "33333"},
	// 	{t1: "44444"},
	// 	{t1: "55555"},
	// 	{t1: "66666"},
	// }
	// x := aa{
	// 	tt: make(chan mm), // Khởi tạo kênh tt
	// }
	// for i := 0; i < len(a); i++ {
	// 	fmt.Println(a[i])
	// 	// x.tt <- a[i]
	// 	go func(i int) {
	// 		x.tt <- a[i] // Gửi dữ liệu vào kênh tt
	// 	}(i)
	// }
	// fmt.Println("========")
	// for i := 0; i < len(a); i++ {
	// 	x.loopp()
	// }
	// time.Sleep(2 * time.Second)
	// for i := 0; i < len(a); i++ {
	// 	fmt.Println(a[i])
	// 	go func(i int) {
	// 		x.tt <- a[i] // Gửi dữ liệu vào kênh tt
	// 	}(i)
	// }
	// for i := 0; i < len(a); i++ {
	// 	x.loopp()
	// }

	// txHash, err := hex.DecodeString("mmmm")
	// if err != nil {
	// 	fmt.Println("eerrr", err)
	// }
	// fmt.Println(txHash)

	// time.Sleep(2 * time.Second)
	// start := time.Now()
	// a, _ := ma.GetLatestLedger()
	// b := uint32(554766)

	// numLedger := int64(a - b)
	// ledgerClosingTime := 4 * time.Second

	// timeWait := time.Duration(numLedger * ledgerClosingTime.Nanoseconds())
	// fmt.Println(timeWait)
	// fmt.Println(a)
	// time.Sleep(timeWait)

	// end := time.Now()
	// duration := end.Sub(start)
	// fmt.Println("Thời gian thực hiện:", duration)

}
