package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

var reader = bufio.NewReader(os.Stdin)

func main() {
	k := setup_game()
	run_game(k)
}

func setup_game() [4]int {
	rand.Seed(time.Now().UnixNano())
	int1 := rand.Intn(10) + 2
	int2 := rand.Intn(10) + 2
	int3 := rand.Intn(10) + 2
	ans := int1*int2 - int3
	numbers := [4]int{int1, int2, int3, ans}
	return numbers
}

func run_game(values [4]int) {
	// get them started
	fmt.Println("Think of a number")
	reader.ReadString('\n')

	// do some multiplication
	fmt.Println("Multiply by", values[0])
	reader.ReadString('\n')

	// a little more
	fmt.Println("Multiply by", values[1])
	reader.ReadString('\n')

	// apply the trickery
	fmt.Println("Divide by your original number")
	reader.ReadString('\n')

	// the dreaded subtraction
	fmt.Println("Subtract", values[2])
	reader.ReadString('\n')

	// give them the answer
	fmt.Println("Now you have", fmt.Sprint(values[3])+"!")
}
