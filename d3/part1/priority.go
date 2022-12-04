package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func normalise(c rune) rune {
	if c > 96 && c < 123 {
		return c - 96
	} else if c > 64 && c < 91 {
		return c - 38
	}
	log.Fatal("ASCII character out of bounds")
	return -1
}

func main() {
	var duplicates []int
	sum := 0
	fileName := os.Args[2]
	file, openErr := os.Open(fileName)

	if openErr != nil {
		log.Fatal(openErr)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var mapping1[53]rune
		var mapping2[53]rune
		
		line := []rune(scanner.Text())
		mid := len(line) / 2
		left := line[:mid]
		right := line[mid:]

		for i := 0; i < mid; i++ {
			c := normalise(left[i])
			mapping1[c] += 1
		}

		for i := 0; i < mid; i++ {
			c := normalise(right[i])
			mapping2[c] += 1
		}

		for i := 1; i <= 52; i++ {
			if mapping1[i] > 0 && mapping2[i] > 0 {
				duplicates = append(duplicates, i)
			}
		}
	}

	for i := 0; i < len(duplicates); i++ {
		sum += duplicates[i]
	}

	fmt.Println(sum)
}