package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/eiannone/keyboard"
)

var reader = bufio.NewReader(os.Stdin)

func main() {

	err := keyboard.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer keyboard.Close()

	pizzas := map[int]string{
		1: "Cheese",
		2: "Pepperoni",
		3: "Sausage",
		4: "Mushrooms",
		5: "Pineapple",
	}

	fmt.Println("MENU")
	fmt.Println("====")
	fmt.Println(pizzas)
	// for key, val := range pizzas {
	// 	fmt.Println(key, "-", val)
	// }
	// fmt.Println("Q - QUIT")

	for {

		char, _, err := keyboard.GetSingleKey()
		if err != nil {
			log.Fatal(err)
		}

		if char == 'Q' || char == 'q' {
			break
		}

		i, _ := strconv.Atoi(string(char))
		val, found := pizzas[i]
		if found {
			fmt.Println(fmt.Sprintf("You chose %s", val))
		} else {
			fmt.Println("That's not valid!")
		}
	}

	fmt.Println("Exiting")

}
