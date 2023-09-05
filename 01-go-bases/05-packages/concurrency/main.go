package main

import (
	"fmt"
	"time"
)

func main() {
	go workAddition(1, 2)
	go workAddition(3, 4)
	go workAddition(5, 6)

	time.Sleep(2 * time.Second)
	fmt.Println("Waiting...")
}

func workAddition(a, b int) {
	result := a + b
	time.Sleep(1 * time.Second)
	fmt.Println("Result:", result)
}