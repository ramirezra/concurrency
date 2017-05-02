package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	var Ball int
	table := make(chan int)
	go player(table, "Ping")
	go player(table, "Pong")
	go player(table, "Peng")

	table <- Ball
	time.Sleep(1 * time.Second)
	<-table
}

func player(table chan int, state string) {
	for {
		ball := <-table
		ball++
		time.Sleep(100 * time.Millisecond)
		fmt.Println(state + ": " + strconv.Itoa(ball))
		table <- ball
	}
}
