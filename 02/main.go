package main

import "fmt"

func main() {
	fmt.Println("Bon appétit!")

	friend := make(chan string)
	food := make(chan string)
	go Order(food, "chorizo")
	go Eat(friend, "Alice")
	go Eat(friend, "Bob")
}

func Order(food chan string, foodName string) {
	food <- foodName
}
func Eat(friend chan string, name string) {
	fmt.Println(name + "is enjoying some" + foodname)
}
