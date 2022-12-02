// Problem: https://adventofcode.com/2022/day/1
// Time Complexity: O(n)

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	currentTotal, most := 0, 0

	fileName := os.Args[2]
	file, openErr := os.Open(fileName)
	
	if openErr != nil {
		log.Fatal(openErr)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for {
		hasNextLine := scanner.Scan()
		readLine := scanner.Text()

		// eof
		if hasNextLine == false {
			most = max(currentTotal, most)
			break
		}

		// new line
		if readLine == "" {
			most = max(currentTotal, most)
			currentTotal = 0
		} else {
			// number line
			i, atoiErr := strconv.Atoi(readLine)
			if atoiErr != nil {
				log.Fatal(atoiErr)
			}
			currentTotal = currentTotal + i
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(most)
}
