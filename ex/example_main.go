package tutorial

import (
	"fmt"
	"time"
	// "github.com/multiformats/go-multihash"
)

// import "fmt"

func Main() {
	start := time.Now()
	// seq, err := GetLatestLedger()
	// panicIf(err)
	// fmt.Println(seq)
	// GetEventsFromSeq(seq - uint32(2)*(64))

	// GetEventsFromSeq(uint32(435801))

	// x, _ := GetOldeastLedger()
	// fmt.Println("lll,", x)

	// GetInfoTxFromTxHash("fa83440afa28eb3c02e70da0a6befb38008ec2e0a7956ce74e74340d0a9eab4f")
	example()
	// helloworld()
	// statistics()
	// claimables()
	end := time.Now()
	duration := end.Sub(start)
	fmt.Println("Thời gian thực hiện:", duration)
}
