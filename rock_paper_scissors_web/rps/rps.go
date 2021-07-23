package rps

import (
	"math/rand"
	"time"
)

const (
	ROCK     = 0
	PAPER    = 1
	SCISSORS = 2
)

type Round struct {
	Result         int
	ComputerChoice string
	RoundResult    string
}

func PlayRound(playerValue int) Round {

	var computerChoice, roundResult string

	rand.Seed(time.Now().UnixNano())
	computerValue := rand.Intn(3)

	switch computerValue {
	case ROCK:
		computerChoice = "Computer chose ROCK"
	case PAPER:
		computerChoice = "Computer chose PAPER"
	case SCISSORS:
		computerChoice = "Computer chose SCISSORS"
	default:
		panic("How the heck did this happen!?")
	}

	result := (playerValue - computerValue + 3) % 3
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

	return Round{result, computerChoice, roundResult}
}
