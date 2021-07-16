package main

import (
	"container/heap"
	"fmt"
	"time"
)

// const (
// 	ROCK     = 0
// 	PAPER    = 1
// 	SCISSORS = 2
// )

// func main() {
// 	clearScreen()
// 	for i := 0; i < 3; i++ {
// 		game()
// 	}
// }

// func game() {
// 	rand.Seed(time.Now().UnixNano())
// 	playerChoice := ""
// 	playerValue := -1

// 	computerValue := rand.Intn(3)

// 	reader := bufio.NewReader(os.Stdin)

// 	fmt.Print("Please enter rock, paper, or scissors -> ")
// 	playerChoice, _ = reader.ReadString('\n')
// 	playerChoice = strings.Replace(playerChoice, "\n", "", -1)

// 	if playerChoice == "rock" {
// 		playerValue = ROCK
// 	} else if playerChoice == "paper" {
// 		playerValue = PAPER
// 	} else if playerChoice == "scissors" {
// 		playerValue = SCISSORS
// 	}

// 	switch computerValue {
// 	case ROCK:
// 		fmt.Println("Computer chose ROCK")
// 	case PAPER:
// 		fmt.Println("Computer chose PAPER")
// 	case SCISSORS:
// 		fmt.Println("Computer chose SCISSORS")
// 	default:
// 		fmt.Println("How the heck did this happen!?")
// 	}

// 	switch ((playerValue-computerValue)%3 + 3) % 3 {
// 	case 0:
// 		fmt.Println("draw!")
// 	case 1:
// 		fmt.Println("win!")
// 	case 2:
// 		fmt.Println("lose!")
// 	default:
// 		fmt.Println("What!? Again?")
// 	}

// }

// // clearScreen clears the screen
// func clearScreen() {
// 	if strings.Contains(runtime.GOOS, "windows") {
// 		// windows
// 		cmd := exec.Command("cmd", "/c", "cls")
// 		cmd.Stdout = os.Stdout
// 		cmd.Run()
// 	} else {
// 		// linux or mac
// 		cmd := exec.Command("clear")
// 		cmd.Stdout = os.Stdout
// 		cmd.Run()
// 	}
// }

func main() {
	control := make(chan int, 1)
	sorted := make(chan int)
	source := make(chan int)
	go writer(control, sorted, source)
	go reader(control, sorted)
	control <- 1
	for _, val := range [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0} {
		// fmt.Println(val)
		source <- val
	}
	close(source)
	time.Sleep(5 * time.Second)
}

func reader(control, sorted chan int) {
	var sVal int
	var sOpen bool
	fmt.Println("ready to read")
	control <- 0
	for sVal, sOpen = <-sorted; sOpen; sVal, sOpen = <-sorted {
		fmt.Println(sVal)
		control <- 1
		time.Sleep(time.Second)
		control <- 0
	}
	fmt.Println("done with read")
}

func writer(control, sorted, source chan int) {
	queue := &IntHeap{}
	heap.Init(queue)
	for cVal, cOpen := <-control; cOpen; cVal, cOpen = <-control {
		// fmt.Println("control", cVal)
		if cVal == 1 {
			uVal := <-source
			heap.Push(queue, uVal)
			// control <- 0
		} else {
			sVal := heap.Pop(queue)
			fmt.Println("writing to sorted")
			sorted <- sVal.(int)
			// sorted <- sVal
		}
		time.Sleep(500 * time.Millisecond)
		control <- 1
	}
}

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	// func (h *IntHeap) Push(x int) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	// func (h *IntHeap) Pop() int {
	var x int
	old := *h
	n := len(old)
	x = old[n-1]
	*h = old[0 : n-1]
	return x
	// return x.(int)
}

// This example inserts several ints into an IntHeap, checks the minimum,
// and removes them in order of priority.
func main2() {
	h := &IntHeap{2, 1, 5}
	heap.Init(h)
	var v int = 3
	heap.Push(h, v)
	fmt.Printf("minimum: %d\n", (*h)[0])
	for h.Len() > 0 {
		fmt.Printf("%d ", heap.Pop(h))
	}
}
