// Problem: https://adventofcode.com/2022/day/1
// Time Complexity: O(nlogn)

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {
	currentTotal, res := 0, 0
	elves := []int{}

	fileName := os.Args[2]
	file, openErr := os.Open(fileName)
	
	if openErr != nil {
		log.Fatal(openErr)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		readLine := scanner.Text()

		// new line
		if readLine == "" {
			elves = append(elves, currentTotal)
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

	// for final elf unidentified by new line
	elves = append(elves, currentTotal)

	// scan error
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// sort in descending
	sort.Slice(elves, func(i, j int) bool {
		return elves[i] > elves[j] 
	})

	// sum up top 3
	for i := 0; i < 3; i++ {
		res = res + elves[i]
	}

	fmt.Println(res)
}
