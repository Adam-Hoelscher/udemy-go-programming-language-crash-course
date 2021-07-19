package main

import (
	"fmt"
	"time"
)

// const second time.Millisecond = 1

// func main() {
// 	request := make(chan int)
// 	response := make(chan int)
// 	prospect := make(chan int)
// 	go leadSource(prospect)
// 	go queueManager(request, response, prospect)
// 	for agentId := 0; agentId < 300; agentId++ {
// 		go agent(request, response, agentId)
// 		time.Sleep(time.Millisecond)
// 	}
// 	// time.Sleep(8.5 * 60 * 60 * time.Millisecond)
// 	time.Sleep(60 * time.Millisecond)
// }

// func leadSource(source chan int) {
// 	for i := 0; i < 8*60*5; i++ {
// 		source <- 1 + rand.Intn(1000)
// 		time.Sleep(20 * time.Millisecond)
// 	}
// 	fmt.Println("closing source")
// 	close(source)
// }

// func agent(request chan int, sorted chan int, agentId int) {
// 	// fmt.Println("Agent", agentId, "ready for call")
// 	request <- agentId
// 	for sVal, sOpen := <-sorted; sOpen; sVal, sOpen = <-sorted {
// 		fmt.Println("Agent", agentId, "calling lead with score of", sVal)
// 		minutes := time.Duration(rand.Intn(5) + 5)
// 		time.Sleep(minutes * 60 * time.Millisecond)
// 		request <- agentId
// 	}
// 	fmt.Println("done calling")
// }

// func queueManager(request chan int, sorted, source chan int) {

// 	queue := &IntHeap{}
// 	heap.Init(queue)
// 	var cVal int

// 	sourceClosed := false

// 	for {

// 		select {

// 		case agentId := <-request:
// 			fmt.Println("Lead requested by agent", agentId)
// 			if queue.Len() > 0 {
// 				cVal = heap.Pop(queue).(int)
// 			} else if !sourceClosed {
// 				cVal = (<-source)
// 			}
// 			fmt.Println("Sending lead with score of", cVal, "to agent", agentId)
// 			sorted <- cVal

// 		default:
// 			// pass
// 		}

// 		sVal, sOpen := <-source
// 		if sOpen {
// 			fmt.Println("Received lead with score of", sVal)
// 			heap.Push(queue, sVal)
// 		} else if queue.Len() == 0 && !sourceClosed {
// 			close(sorted)
// 			sourceClosed = true
// 		}
// 	}

// }

// type IntHeap []int

// func (h IntHeap) Len() int { return len(h) }

// /*
// The heap package implements a MinHeap, which prioritizes lower values.
// Because we want to prioritize higher values, we'll say that a higher value
// is "Less" than a lower value
// */
// func (h IntHeap) Less(i, j int) bool { return h[i] > h[j] }
// func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

// func (h *IntHeap) Push(x interface{}) {
// 	// Push and Pop use pointer receivers because they modify the slice's
// 	// length, not just its contents.
// 	*h = append(*h, x.(int))
// }

// func (h *IntHeap) Pop() interface{} {
// 	old := *h
// 	n := len(old)
// 	x := old[n-1]
// 	*h = old[0 : n-1]
// 	return x
// }

func main() {
	var chan1 chan int
	var chan2 chan int
	go pub(chan1, chan2)
	go sub(chan1, chan2)
	time.Sleep(5 * time.Second)
}

func pub(chan1, chan2 chan int) {
	chan1 <- 1
	// time.Sleep(time.Second)
	chan2 <- 1
	// time.Sleep(time.Second)
	chan1 <- 1
	chan2 <- 2
}

func sub(chan1, chan2 chan int) {
	for i := 0; i < 3; i++ {
		fmt.Println("pass", i)
		select {
		case q := <-chan1:
			fmt.Println("got", q, "from 1")
		case q := <-chan2:
			fmt.Println("got", q, "from 2")
		default:
			fmt.Println("got nothing")
		}
		time.Sleep(time.Second)
	}
}
