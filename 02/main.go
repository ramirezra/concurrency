package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

type dish struct {
	name string
	qty  int
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

// Order five dishes, each between 5 - 10 morsels.
func orderChorizo(dish dish, chorizo chan dish, wg *sync.WaitGroup) {
	fmt.Printf("%d dishes of %s were served.\n", dish.qty, dish.name)
	for i := dish.qty; i > 0; i-- {
		defer wg.Done()
		dish.qty--
		fmt.Printf("%d piece(s) of %s are left.\n", dish.qty, dish.name)
		chorizo <- dish
	}
	fmt.Printf("All pieces of %s have been eaten.\n", dish.name)
	close(chorizo)
}

func orderBunuelo(dish dish, bunuelo chan dish, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("%d dishes of %s were served.\n", dish.qty, dish.name)
	for i := dish.qty; i > 0; i-- {
		bunuelo <- dish
		dish.qty--
		fmt.Printf("%d piece(s) of %s are left.\n", dish.qty, dish.name)
	}
	fmt.Printf("All pieces of %s have been eaten.\n", dish.name)
	close(bunuelo)
}

func orderChicharron(dish dish, chicharron chan dish, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("%d dishes of %s were served.\n", dish.qty, dish.name)
	for i := dish.qty; i > 0; i-- {
		chicharron <- dish
		dish.qty--
		fmt.Printf("%d piece(s) of %s are left.\n", dish.qty, dish.name)
	}
	fmt.Printf("All pieces of %s have been eaten.\n", dish.name)
	close(chicharron)
}

func orderEmpanadas(dish dish, empanadas chan dish, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("%d dishes of %s were served.\n", dish.qty, dish.name)
	for i := dish.qty; i > 0; i-- {
		empanadas <- dish
		dish.qty--
		fmt.Printf("%d piece(s) of %s are left.\n", dish.qty, dish.name)
	}
	fmt.Printf("All pieces of %s have been eaten.\n", dish.name)
	close(empanadas)
}

func orderAlbondigas(dish dish, albondigas chan dish, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("%d dishes of %s were served.\n", dish.qty, dish.name)
	for i := dish.qty; i > 0; i-- {
		albondigas <- dish
		dish.qty--
		fmt.Printf("%d piece(s) of %s are left.\n", dish.qty, dish.name)
	}
	fmt.Printf("All pieces of %s have been eaten.\n", dish.name)
	close(albondigas)
}

// Spend between 30 seconds and 3 minutes eating each morsel.
func eatDish(friendName string, chorizo chan dish, bunuelos chan dish, chicharron chan dish, empanadas chan dish, albondigas chan dish) {
	chorizo_closed := false
	for {
		if chorizo_closed {
			return
		}
		select {
		case cho, ok := <-chorizo:
			if !ok {
				chorizo_closed = true
			} else {
				fmt.Printf("%v is enjoying some %v\n", friendName, cho.name)
				time.Sleep(time.Second * time.Duration((rand.Intn(3) + 1)))
			}
		case bun := <-bunuelos:
			fmt.Printf("%v is enjoying some %v\n", friendName, bun.name)
			time.Sleep(time.Second * time.Duration((rand.Intn(3) + 1)))
		case chi := <-chicharron:
			fmt.Printf("%v is enjoying some %v\n", friendName, chi.name)
			time.Sleep(time.Second * time.Duration((rand.Intn(3) + 1)))
		case emp := <-empanadas:
			fmt.Printf("%v is enjoying some %v\n", friendName, emp.name)
			time.Sleep(time.Second * time.Duration((rand.Intn(3) + 1)))
		case alb := <-albondigas:
			fmt.Printf("%v is enjoying some %v\n", friendName, alb.name)
			time.Sleep(time.Second * time.Duration((rand.Intn(3) + 1)))
		}
	}
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	// set the number of OS threads to use the number of CPUs.
	fmt.Println("Bon appétit!")
	var wg sync.WaitGroup
	wg.Add(9)
	// Declare five dishes.
	chorizo := make(chan dish)
	bunuelos := make(chan dish)
	chicharron := make(chan dish)
	empanadas := make(chan dish)
	albondigas := make(chan dish)
	go orderChorizo(dish{name: "chorizo", qty: (rand.Intn(5) + 5)}, chorizo, &wg)
	// go orderBunuelo(dish{name: "bunuelos", qty: (rand.Intn(5) + 5)}, bunuelos, &wg)
	// go orderChicharron(dish{name: "chicharron", qty: (rand.Intn(5) + 5)}, chicharron, &wg)
	// go orderEmpanadas(dish{name: "empanadas", qty: (rand.Intn(5) + 5)}, empanadas, &wg)
	// go orderAlbondigas(dish{name: "albondigas", qty: (rand.Intn(5) + 5)}, albondigas, &wg)

	go eatDish("Alice", chorizo, bunuelos, chicharron, empanadas, albondigas)
	go eatDish("Bob", chorizo, bunuelos, chicharron, empanadas, albondigas)
	go eatDish("Charlie", chorizo, bunuelos, chicharron, empanadas, albondigas)
	go eatDish("Dave", chorizo, bunuelos, chicharron, empanadas, albondigas)

	wg.Wait()
}
