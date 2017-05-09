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

// func main() {
// 	fmt.Println(runtime.NumCPU())
// 	tourist := make(chan int)
// 	for i := 1; i <= 25; i++ {
// 		tourist <- i
// 	}
// 	close(tourist)
//
// 	// go terminal(tourist)
// 	for chat := range tourist {
// 		fmt.Println(chat)
// 	}
// }

// func terminal(tourist chan int) {
// 	for chat := range tourist {
// 		fmt.Println(chat)
// 	}
// }

type Request struct {
	args       []int
	f          func([]int) int
	resultChan chan int
}

func sum(a []int) (s int) {
	for _, v := range a {
		s += v
	}
	return
}

func main() {
	request := &Request{[]int{3, 4, 5}, sum, make(chan int)}

	clientRequests <- request

	fmt.Printf("answer: %d\n", <-request.resultChan)
}
