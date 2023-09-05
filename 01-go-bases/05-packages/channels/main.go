package main

import (
	"fmt"
	"time"
)

func main() {
	// create channel to communicate main and goroutines
	ch := make(chan int)
	go workAddition(ch, 1, 2)
	go workAddition(ch, 3, 4)
	go workAddition(ch, 5, 6)

	// receive result from channel
	result1 := <-ch	// temporal deadlock

	fmt.Println("Waiting...")

	result2 := <-ch
	result3 := <-ch
	// result4 := <-ch ! panic: all goroutines are asleep - deadlock!

	// print result
	fmt.Println("Program finished:", result1, result2, result3)
}

func workAddition(ch chan int, a, b int) {
	// process addition
	result := a + b
	time.Sleep(1 * time.Second)

	// send result to channel
	ch <- result
}