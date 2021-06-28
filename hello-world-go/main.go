package main

import (
	"bufio"
	"fmt"
	"hello-world/doctor"
	"os"
	"strings"
)

// import "doctor"

func main() {
	reader := bufio.NewReader(os.Stdin)

	whatToSay := doctor.Intro()
	fmt.Println(whatToSay)

	for {
		fmt.Print("-> ")
		userInput, _ := reader.ReadString('\n')

		userInput = strings.Replace(userInput, "\n", "", -1)
		userInput = strings.Replace(userInput, "\r", "", -1)

		if userInput == "quit" {
			fmt.Println("about to quit")
			break
		} else {
			fmt.Println(doctor.Response(userInput))
		}

	}

}
