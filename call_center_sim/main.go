package main

import (
	"container/heap"
	"fmt"
	"math/rand"
	"time"
)

const simSecond = 1_000 * time.Microsecond

func main() {

	request := make(chan int)
	response := make(chan int)
	prospect := make(chan int)

	go queueManager(request, response, prospect)
	for i := 0; i < 1_000; i++ {
		insertLead(prospect)
	}

	go leadSource(prospect)
	for agentId := 0; agentId < 300; agentId++ {
		go agent(request, response, agentId)
		time.Sleep(simSecond)
	}
	time.Sleep(8.5 * 60 * 60 * simSecond)
	// time.Sleep(100 * simSecond)
}

func insertLead(prospect chan int) {
	prospect <- 1 + rand.Intn(1_000)
}

func leadSource(prospect chan int) {

	for i := 0; i < 8*60*5; i++ {
		insertLead(prospect)
		time.Sleep(20 * simSecond)
	}

	close(prospect)
}

func agent(request chan int, sorted chan int, agentId int) {

	fmt.Println("Agent", agentId, "ready for call")
	request <- agentId

	for sVal, sOpen := <-sorted; sOpen; sVal, sOpen = <-sorted {
		fmt.Println("Agent", agentId, "calling lead with score of", sVal)
		minutes := time.Duration(rand.Intn(5) + 5)
		time.Sleep(minutes * 60 * simSecond)
		request <- agentId
	}
}

func queueManager(request chan int, sorted, source chan int) {

	queue := &IntHeap{}
	heap.Init(queue)
	var cVal int

	sourceOpen := true

	for {

		select {

		case agentId := <-request:
			fmt.Println("Lead requested by agent", agentId)
			if queue.Len() > 0 {
				cVal = heap.Pop(queue).(int)
			} else if sourceOpen {
				cVal = <-source
			}
			fmt.Println("Sending lead with score of", cVal, "to agent", agentId)
			if sourceOpen {
				sorted <- cVal
			}

		default:
			// pass
		}

		sVal, sOpen := <-source
		if sOpen {
			fmt.Println("Received lead with score of", sVal)
			heap.Push(queue, sVal)
		} else if queue.Len() == 0 && sourceOpen {
			close(sorted)
			sourceOpen = false
		}
	}

}

type IntHeap []int

func (h IntHeap) Len() int { return len(h) }

/*
The heap package implements a MinHeap, which prioritizes lower values.
Because we want to prioritize higher values, we'll say that a higher value
is "Less" than a lower value
*/
func (h IntHeap) Less(i, j int) bool { return h[i] > h[j] }

func (h IntHeap) Max() int {
	temp := h[0]
	for _, val := range h {
		if val > temp {
			temp = val
		}
	}
	return temp
}

func (h IntHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's
	// length, not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
