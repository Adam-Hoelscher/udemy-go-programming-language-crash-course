// package main

// import (
// 	"fmt"
// 	"log"
// 	"math/rand"
// 	"time"
// )

// func main() {
// 	x := 2
// 	y := 0

// 	// delay := 30 * time.Millisecond
// 	stopAt := time.Now().Add(10 * time.Second)

// 	go square(&x)
// 	go increment(&x)

// 	for time.Now().Unix() < stopAt.Unix() {
// 		if y != x {
// 			y = x
// 			log.Println("x is now", x)
// 		}
// 	}
// }

// func square(x *int) {
// 	delay := time.Duration(500_000*rand.Float32()) * time.Microsecond
// 	fmt.Println("Square delay is", delay)
// 	for i := 0; i < 5; i++ {
// 		time.Sleep(delay)
// 		log.Println("^2")
// 		*x *= *x
// 	}
// }

// func increment(x *int) {
// 	delay := time.Duration(500_000*rand.Float32()) * time.Microsecond
// 	fmt.Println("Increment delay is", delay)
// 	for i := 0; i < 5; i++ {
// 		time.Sleep(delay)
// 		log.Println("+1")
// 		*x += 1
// 	}
// }

package main

import (
	"fmt"

	"github.com/eiannone/keyboard"
)

var keyPressChan chan keyboard.KeyEvent

func main() {

	keyPressChan = make(chan keyboard.KeyEvent)

	go listenForKeyPress(keyPressChan)

	fmt.Println("Press any key or ESC to quit")

	keyboard.Open()
	defer keyboard.Close()

	for {

		ch, key, err := keyboard.GetSingleKey()
		event := keyboard.KeyEvent{Key: key, Rune: ch, Err: err}

		if event.Key == keyboard.KeyEsc {
			break
		}
		keyPressChan <- event
	}

}

func listenForKeyPress(keyPressChan chan keyboard.KeyEvent) {
	for {
		event := <-keyPressChan
		fmt.Println("You pressed", string(event.Rune))
	}
}
