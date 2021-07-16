package main

import (
	"fmt"
)

func main() {
	myChan := make(chan int)
	go publish(myChan)
	for i, ok := <-myChan; ok; i, ok = <-myChan {
		fmt.Println(i)
	}
}

func publish(myChan chan int) {
	for i := 0; i < 10; i++ {
		myChan <- i
	}
	close(myChan)
}
