package main

import (
	"fmt"
	"strings"
)

func main() {
	// reader
	str := strings.NewReader("Hello, World!")

	// read
	buffer := make([]byte, 2)
	for {
		// read bytes
		n, err := str.Read(buffer)
		if err != nil {
			switch err.Error() {
				case "EOF":
					println("EOF")
				default:
					println(err.Error())
			}
			return
		}

		fmt.Printf("Message: %s - byte: %d\n", buffer[:n], n)
	}
}