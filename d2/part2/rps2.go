package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type score struct {
	win int
	draw int
	lose int
}
// Rock: 1, Paper: 2, Scissors: 3
// A: Rock, B: Paper, C: Scissors
// X – lose, Y – draw, Z: win
func calculateScore(o, y string) int {
	// map of opponent's move to your move's associated score
	m := make(map[string]score)
	m["A"] = score{ win: 2, draw: 1, lose: 3 } // Rock
	m["B"] = score{ win: 3, draw: 2, lose: 1 } // Paper
	m["C"] = score{ win: 1, draw: 3, lose: 2 } // Scissors
	
	score := 0

	switch y {
	case "X":
		score = score + m[o].lose + 0
	case "Y":
		score = score + m[o].draw + 3
	case "Z":
		score = score + m[o].win + 6
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