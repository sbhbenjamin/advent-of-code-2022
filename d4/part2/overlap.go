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

		hasOverlap := firstEnd >= secondStart && firstStart <= secondEnd

		if hasOverlap {
			result += 1
		}
	}

	fmt.Println(result)
}