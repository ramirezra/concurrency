package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	rand.Seed(time.Now().UnixNano())
}

func main() {
	fmt.Println(runtime.NumCPU())
	tourist := make(chan int)
	for i := 1; i <= 25; i++ {
		tourist <- i
	}
	close(tourist)

	// go terminal(tourist)
	for chat := range tourist {
		fmt.Println(chat)
	}
}

// func terminal(tourist chan int) {
// 	for chat := range tourist {
// 		fmt.Println(chat)
// 	}
// }
