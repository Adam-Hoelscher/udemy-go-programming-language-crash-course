package rps

import (
	"math/rand"
	"strings"
	"time"
)

var choiceValues = map[string]int{
	"ROCK":     0,
	"PAPER":    1,
	"SCISSORS": 2,
}

var valueChoices = [3]string{"ROCK", "PAPER", "SCISSORS"}

type Round struct {
	PlayerChoide   string `json:"player_choice"`
	ComputerChoice string `json:"computer_choice"`
	RoundResult    string `json:"round_result"`
}

func PlayRound(playerChoice string) Round {

	var computerChoice, roundResult string

	rand.Seed(time.Now().UnixNano())
	computerValue := rand.Intn(3)
	computerChoice = "Computer chose " + valueChoices[computerValue]

	playerChoice = strings.ToUpper(playerChoice)
	playerValue := choiceValues[playerChoice]
	playerChoice = "Player chose " + playerChoice

	result := (3 + playerValue - computerValue) % 3
	switch result {
	case 0:
		roundResult = "Draw!"
	case 1:
		roundResult = "Win!"
	case 2:
		roundResult = "Lose!"
	default:
		panic("What!? Again?")
	}

	return Round{playerChoice, computerChoice, roundResult}
}
