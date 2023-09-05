package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// file
	f, err := os.Open("./files/files.txt")
	if err != nil {
		println(err.Error())
		return
	}
	defer f.Close()

	// read
	buf := make([]byte, 20)
	for {
		// read bytes
		n, err := f.Read(buf)
		if err != nil {
			switch err.Error() {
				case "EOF":
					println("EOF")
				default:
					println(err.Error())
			}
			break
		}

		// print
		fmt.Println("Message:", string(buf[:n]), "- byte:", n)
	}

	// reset
	f.Seek(0, io.SeekStart)

	// read all
	data, err := io.ReadAll(f)
	if err != nil {
		println(err.Error())
		return
	}

	// print
	fmt.Println("Message:", string(data))

	io.Copy(os.Stdout, f)
}