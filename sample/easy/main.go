package main

import (
	"fmt"
	"time"
)

func main() {
	bridge := make(chan int) //create the channel
	go count(bridge)
	for i := 0; i < 10; i++ {
		bridge <- i
		fmt.Println(i, " in the main function ")
	}
	//
	time.Sleep(1 * time.Second)
}

func count(bridgereference chan int) {
	for i := 0; i < 10; i++ {
		time.Sleep(1 * time.Second)
		<-bridgereference
		fmt.Println(i, " In the Go routine")
	}

}
