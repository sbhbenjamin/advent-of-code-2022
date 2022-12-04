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

func addToMap(mapping []rune, line []rune) {
	for i := 0; i < len(line); i++ {
		c := normalise(line[i])
		mapping[c] += 1
	}
}

func main() {
	sum := 0
	fileName := os.Args[2]
	file, openErr := os.Open(fileName)

	if openErr != nil {
		log.Fatal(openErr)
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		var duplicates []int

		mapping1 := make([]rune, 53)
		mapping2 := make([]rune, 53)
		mapping3 := make([]rune, 53)

		addToMap(mapping1, []rune(scanner.Text()))
		scanner.Scan()
		addToMap(mapping2, []rune(scanner.Text()))
		scanner.Scan()
		addToMap(mapping3, []rune(scanner.Text()))

		for i := 1; i < 53; i++ {
			if mapping1[i] > 0 && mapping2[i] > 0 && mapping3[i] > 0 {
				duplicates = append(duplicates, i)
			}
		}

		for i := 0; i < len(duplicates); i++ {
			sum += duplicates[i]
		}
	}

	fmt.Println(sum)
}