package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func open_and_scan(fileName string) *bufio.Scanner {
	file, openErr := os.Open(fileName)

	if openErr != nil {
		log.Fatal(openErr)
	}

	scanner := bufio.NewScanner(file)
	return scanner
}

func get_num_stacks(line string) int {
	crates := []rune(line)
	line_length := len(crates)

	num_spaces := (line_length - 3) / (3 + 1)
	num_stacks := (line_length - num_spaces) / 3
	
	return num_stacks;
}

func parse_and_add_to_stacks(stacks [][]int, scanner *bufio.Scanner) {
	for scanner.Scan() {
		line := scanner.Text()
		crates := []rune(line)
		line_length := len(line)

		current_stack := 0

		for i := 1; i < line_length; i = i + 4 {
			current_crate := int(crates[i])

			is_space := current_crate == 32
			is_number := current_crate >= 49 && current_crate <= 57

			if !is_space && !is_number {
				stacks[current_stack] = append(stacks[current_stack], current_crate)
			}

			current_stack += 1
		}

		if line == "" {
			break
		}
	}
}

func parse_and_rearrange(stacks [][]int, scanner *bufio.Scanner) {
	for scanner.Scan() {
		line := scanner.Text()
		splitLine := strings.Split(line, " ")

		num_to_move, _ := strconv.Atoi(splitLine[1])
		move_from, _ := strconv.Atoi(splitLine[3])
		move_to, _ := strconv.Atoi(splitLine[5])

		move_from -= 1
		move_to -= 1
		
		for i := 0; i < num_to_move; i++ {
			value_to_move := stacks[move_from][0]
			stacks[move_from] = stacks[move_from][1:]
			stacks[move_to] = append([]int{value_to_move}, stacks[move_to]...)
		}
	}
}

func retrieve_answer(stacks [][]int) string {
	answer := ""
	for i := 0; i < len(stacks); i++ {
		answer += string(stacks[i][0])
	}
	return answer
}

func main() {
	fileName := os.Args[2]
	scanner := open_and_scan(fileName)

	// calculate number of stacks
	scanner.Scan()
	first_line := scanner.Text()
	num_stacks := get_num_stacks(first_line)

	// create stacks
	stacks := make([][]int, num_stacks)
	for i := 0; i < num_stacks; i++ {
		stacks[i] = make([]int, 0)
	}

	// reset scanner
	scanner = open_and_scan(fileName)

	parse_and_add_to_stacks(stacks, scanner)
	
	parse_and_rearrange(stacks, scanner)

	// print answer
	answer := retrieve_answer(stacks)
	fmt.Println(answer)
}