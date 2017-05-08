package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

type Waiter struct {
	dishes []dish
	Ch     chan chan *string
}

type dish struct {
	name       string
	numMorsels int
}

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	rand.Seed(time.Now().UnixNano())
}

func main() {
	names := []string{"Alice", "Bob", "Charlie", "Dave"}
	fmt.Println("Bon appétit!")
	waiter := NewManager()
	var wg sync.WaitGroup

	for _, name := range names {
		wg.Add(1)
		go eatDish(name, &wg)
	}

	wg.Wait()
	close(waiter.Ch)
	fmt.Println("That was delicious")
}

func eatDish(name string, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		dishNameCh := make(chan *string)
		waiter.Ch <- dishNameCh
		dishName := <-dishNameCh
		if dishName == nil {
			return
		}
		fmt.Printf("%s is enjoying some %s\n", name, *dishName)
		time.Sleep(time.Second * time.Duration((rand.Intn(3) + 1)))
	}
}

func NewManager() *Waiter {
	ch := make(chan chan *string)
	waiter := &Waiter
	dishes := []dish{
		{name: "chorizos", numMorsels: (rand.Intn(5) + 5)},
		{name: "bunuelos", numMorsels: (rand.Intn(5) + 5)},
		{name: "chicharron", numMorsels: (rand.Intn(5) + 5)},
		{name: "empanadas", numMorsels: (rand.Intn(5) + 5)},
		{name: "albondigas", numMorsels: (rand.Intn(5) + 5)},
		Ch: ch,
	}
	return waiter
}
