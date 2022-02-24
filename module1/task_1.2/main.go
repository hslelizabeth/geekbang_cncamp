package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	msg := make(chan int, 10)
	done := make(chan bool)

	defer close(msg)

	//produce
	go func() {
		ticker := time.NewTicker(1 * time.Second)
		for _ = range ticker.C {
			select {
			case <-msg:
				fmt.Println("produce process interrupt...")
				return
			default:
				msg <- rand.Intn(10)
				fmt.Println("produce process")
			}
		}
	}()

	//consume
	go func() {
		ticker := time.NewTicker(1 * time.Second)
		for _ = range ticker.C {
			select {
			case <-done:
				fmt.Println("consume process interrupt...")
				return
			default:
				fmt.Printf("receive message: %d\n", <-msg)
			}
		}
	}()

	time.Sleep(15 * time.Second)
	close(done)
	time.Sleep(1 * time.Second)
	fmt.Println("main process exit!")
}
