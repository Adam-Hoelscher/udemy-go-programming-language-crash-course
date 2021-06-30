package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/jossef/format"
)

var reader = bufio.NewReader(os.Stdin)

type User struct {
	userName string
	userAge  int
	userNum  float64
}

func main() {

	var user User

	user.userName = readString("What is your name?")
	user.userAge = readInt("How old are you?")
	user.userNum = readFloat("What's your favorite number?")

	formattedString := format.String(
		"Your name is {userName}. You are {userAge} years old. Your favorite number is {userNum}",
		format.Items{"userName": user.userName, "userAge": user.userAge, "userNum": user.userNum})

	fmt.Println(formattedString)
}

func prompt() {
	fmt.Print("-> ")
}

func readString(s string) string {

	fmt.Println(s)
	prompt()

	userInput, _ := reader.ReadString('\n')
	userInput = strings.Replace(userInput, "\r", "", -1)
	userInput = strings.Replace(userInput, "\n", "", -1)
	if userInput == "" {
		fmt.Println("Please enter a string")
		userInput = readString(s)
	}
	return userInput

}

func readInt(s string) int {

	userInput := readString(s)
	num, err := strconv.Atoi(userInput)

	if err != nil {
		fmt.Println("Please enter a whole number")
		num = readInt(s)
	}

	return num
}

func readFloat(s string) float64 {

	userInput := readString(s)
	num, err := strconv.ParseFloat(userInput, 64)

	if err != nil {
		fmt.Println("Please enter a number")
		num = readFloat(s)
	}

	return num
}
