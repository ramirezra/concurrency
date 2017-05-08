package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

// GetReady function
func GetReady(done chan bool, name string) {
	fmt.Println(name + " started getting ready.")
	seed := rand.New(rand.NewSource(time.Now().UnixNano()))
	duration := seed.Intn(30) + 60

	fmt.Println(name + " spent " + strconv.Itoa(duration) + " seconds getting ready.")
	done <- true
}

func SetAlarm(done chan bool) {
	fmt.Println("Arming the alarm.")
	time.Sleep(time.Second * 2)
	fmt.Println("The alarm is set.")
	done <- true
}

func Shoes(done chan bool, name string) {
	fmt.Println(name + " started putting on shoes.")
	seed := rand.New(rand.NewSource(time.Now().UnixNano()))
	duration := seed.Intn(10) + 35
	fmt.Println(name + " spend " + strconv.Itoa(duration) + " seconds putting on shoes.")
	done <- true
}

func main() {
	fmt.Println("Let's go for a walk!")

	done := make(chan bool, 2)

	go GetReady(done, "Bob")
	go GetReady(done, "Alice")
	<-done
	<-done
	go SetAlarm(done)
	go Shoes(done, "Bob")
	go Shoes(done, "Alice")

	time.Sleep(time.Second * 1)
	<-done
	<-done
	fmt.Println("Exiting and locking the door.")
	<-done
	close(done)

	for text := range done {
		fmt.Println(text)
	}

	// Bob started putting on shoes
	// Alarm is counting down.
	// Alice started putting on shoes
	// Alice spent 37 seconds putting on shoes
	// Bob spent 39 seconds putting on shoes
	// Exiting and locking the door.
	// Alarm is armed.
}
