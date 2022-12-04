package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	// for a range (a1, a2) to fully contain the other (b1, b2)
	// a1 <= b1 and a2 >= b2

	result := 0

	fileName := os.Args[2]
	file, openError := os.Open(fileName)

	if openError != nil {
		log.Fatal(openError)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		splitLine := strings.Split(line, ",")

		first := strings.Split(splitLine[0], "-")
		firstStart, _ := strconv.Atoi(first[0])
		firstEnd, _ := strconv.Atoi(first[1])

		second := strings.Split(splitLine[1], "-")
		secondStart, _ := strconv.Atoi(second[0])
		secondEnd, _ := strconv.Atoi(second[1])

		firstContainsSecond := firstStart <= secondStart && firstEnd >= secondEnd
		secondContainsFirst := secondStart <= firstStart && secondEnd >= firstEnd

		if firstContainsSecond || secondContainsFirst {
			result += 1
		}
	}

	fmt.Println(result)
}