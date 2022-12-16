// assumption: input is a data-stream, hence we need to read byte-by-byte
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func pop(slice []byte) (byte, []byte) {
	el, remainder := slice[0], slice[1:]
	slice = remainder

	return el, remainder
}

func main() {
	fileName := os.Args[2]
	file, openErr := os.Open(fileName)
	memory := make([]bool, 27)
	window := make([]byte, 0)
	n := 0

	if openErr != nil {
		log.Fatal(openErr)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanBytes)

	for scanner.Scan() {
		n++
		currentChar := scanner.Bytes()[0] - 97 + 1

		if memory[currentChar] {
			for memory[currentChar] {
				el, remainder := pop(window)
				window = remainder
				memory[el] = false
			}
		}
		window = append(window, currentChar)
		memory[currentChar] = true

		if len(window) == 14 {
			fmt.Println(n)
			break
		}
	}
}