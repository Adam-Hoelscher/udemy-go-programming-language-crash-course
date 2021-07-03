package main

import "fmt"

type Animal interface {
	Sound() string
	Legs() int
}

type Dog struct {
	Name  string
	sound string
	legs  int
}

func (dog *Dog) Sound() string {
	return dog.sound
}

func (dog *Dog) Legs() int {
	return dog.legs
}

func main() {
	myDog := Dog{Name: "Fido", sound: "woof", legs: 4}
	riddle(&myDog)

}

func riddle(animal Animal) {
	fmt.Printf("It stands on their %d legs and says '%s'\n", animal.Legs(), animal.Sound())
}

func fib(x int) int {
	if x < 2 {
		return x
	} else {
		return fib(x-1) + fib(x-2)
	}
}
