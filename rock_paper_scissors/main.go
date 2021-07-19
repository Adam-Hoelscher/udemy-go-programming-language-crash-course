package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"
)

const (
	ROCK     = 0
	PAPER    = 1
	SCISSORS = 2
)

func main() {
	clearScreen()
	for i := 0; i < 3; i++ {
		game()
	}
}

func game() {
	rand.Seed(time.Now().UnixNano())
	playerChoice := ""
	playerValue := -1

	computerValue := rand.Intn(3)

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Please enter rock, paper, or scissors -> ")
	playerChoice, _ = reader.ReadString('\n')
	playerChoice = strings.Replace(playerChoice, "\n", "", -1)

	if playerChoice == "rock" {
		playerValue = ROCK
	} else if playerChoice == "paper" {
		playerValue = PAPER
	} else if playerChoice == "scissors" {
		playerValue = SCISSORS
	}

	switch computerValue {
	case ROCK:
		fmt.Println("Computer chose ROCK")
	case PAPER:
		fmt.Println("Computer chose PAPER")
	case SCISSORS:
		fmt.Println("Computer chose SCISSORS")
	default:
		fmt.Println("How the heck did this happen!?")
	}

	switch (3 + playerValue - computerValue) % 3 {
	case 0:
		fmt.Println("draw!")
	case 1:
		fmt.Println("win!")
	case 2:
		fmt.Println("lose!")
	default:
		fmt.Println("What!? Again?")
	}

}

// clearScreen clears the screen
func clearScreen() {
	var cmd *exec.Cmd

	if strings.Contains(runtime.GOOS, "windows") {
		// windows
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		// linux or mac
		cmd = exec.Command("clear")
	}

	cmd.Stdout = os.Stdout
	cmd.Run()
}
