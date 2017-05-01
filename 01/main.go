package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	// Create an buffered channel

	channel := make(chan int)
	message := make(chan string)
	done := make(chan bool)

	fmt.Println("Let's go for a walk.")

	name := []string{"Bob", "Alice"}

	for i := 0; i < 2; i++ {
		go func() {
			channel <- i
			fmt.Printf(name[i] + " started getting ready.\n")

			seed := rand.NewSource(time.Now().Unix())
			count := rand.New(seed)
			duration := count.Intn(60) + 30
			message <- (name[i] + " spent " + strconv.Itoa(duration) + " seconds getting ready.")
			done <- true
		}()
	}
	go func() {
		for i := 0; i < 2; i++ {
			<-done
		}
		close(channel)
		close(message)
	}()
	time.Sleep(4 * time.Second)

	for n := range message {
		fmt.Println(n)
	}
	fmt.Println("Arming alarm.")

}
