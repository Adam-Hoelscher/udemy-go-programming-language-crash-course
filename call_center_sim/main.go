package main

import (
	"callCenterSim/priority"
	"container/heap"
	"fmt"
	"math/rand"
	"time"
)

const simSecond = time.Millisecond
const simMinute = 60 * simSecond
const simHour = 60 * simMinute

type lead struct {
	value int
	new   bool
}

func main() {

	request := make(chan int)
	response := make(chan int)
	prospect := make(chan lead)

	go queueManager(request, response, prospect)
	for i := 0; i < 1_000; i++ {
		insertLead(prospect, false)
	}

	go leadSource(prospect)

	quittingTime := make([]chan bool, 0)
	for agentId := 0; agentId < 50; agentId++ {
		clockOut := make(chan bool)
		quittingTime = append(quittingTime, clockOut)
		go agent(request, response, agentId, clockOut)
		time.Sleep(simSecond)
	}

	for _, clockOut := range quittingTime {
		<-clockOut
	}
}

func insertLead(prospect chan lead, new bool) {
	newLead := lead{1 + rand.Intn(1_000), new}
	prospect <- newLead
}

func leadSource(prospect chan lead) {
	for {
		time.Sleep(4 * simSecond)
		insertLead(prospect, true)
	}
}

func agent(request chan int, sorted chan int, agentId int, clockOut chan bool) {

	var callDuration time.Duration

	fmt.Println("Agent", agentId, "logged in and requested lead")
	request <- agentId

	for sVal, sOpen := <-sorted; sOpen; sVal, sOpen = <-sorted {

		fmt.Println("Agent", agentId, "calling lead with score of", sVal)

		if rand.Intn(1000) < sVal {
			callDuration = time.Duration(5+rand.NormFloat64()*2) * 60 * simSecond
		} else {
			callDuration = time.Duration(0)
		}
		// callDuration = time.Duration(5) * 30 * simSecond
		time.Sleep(callDuration)

		request <- agentId
	}

	fmt.Println("Agent", agentId, "logged out")
	clockOut <- true
}

func queueManager(request, response chan int, prospect chan lead) {

	startTime := time.Now()
	stopTime := startTime.Add(8 * simHour)

	queue := &priority.IntMaxHeap{}

	for time.Now().Before(stopTime) {
		serveRequests(request, response, prospect, queue)
		readNewLead(prospect, queue)
	}

	// Stop responding and discard the remaining requests
	close(response)
	for {
		<-request
	}

}

func serveRequests(
	request, response chan int,
	prospect chan lead,
	queue *priority.IntMaxHeap) {

	var cVal int
	requestsEmpty := false

	for !requestsEmpty {

		select {

		case agentId := <-request:

			fmt.Println("Lead requested by agent", agentId)

			if queue.Len() > 0 {
				fmt.Println("Popping a lead from the queue")
				cVal = heap.Pop(queue).(int)
			} else {
				fmt.Println("Queue is empty. Waiting for a lead")
				cVal = (<-prospect).value
			}
			fmt.Println("Sending lead with value", cVal, "to agent", agentId, queue.Len(), "remaining")
			response <- cVal

		default:
			requestsEmpty = true
		}
	}

}

func readNewLead(
	prospect chan lead,
	queue *priority.IntMaxHeap) {

	select {
	case sLead := <-prospect:
		if sLead.new {
			fmt.Println("Received lead with value of", sLead.value)
		}
		heap.Push(queue, sLead.value)
	default:
		// pass
	}

}
