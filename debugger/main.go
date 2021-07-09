package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		fmt.Print("i: ", i)
		for j := 0; j < 3; j++ {
			fmt.Print("\tj: ", j)
		}
		fmt.Println()
	}
}
