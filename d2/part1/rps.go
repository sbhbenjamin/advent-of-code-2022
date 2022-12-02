package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func calculateScore(o, y string) int {
	score := 0

	switch y {
	case "X":
		score = score + 1
	case "Y":
		score = score + 2
	case "Z":
		score = score + 3
	}

	if o == "A" && y == "Y" || o == "B" && y == "Z" || o == "C" && y == "X" {
		score = score + 6
	} else if o == "A" && y == "X" || o == "B" && y == "Y" || o == "C" && y == "Z" {
		score = score + 3
	}

	return score
}

func main() {
	result := 0
	fileName := os.Args[2]
	file, openErr := os.Open(fileName)

	if openErr != nil {
		log.Fatal(openErr)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		readLine := scanner.Text()
		s := strings.Split(readLine, " ")
		result = result + calculateScore(s[0], s[1])
	}

	fmt.Println(result)
}