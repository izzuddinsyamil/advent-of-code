package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal("Error opening file:", err)
	}
	defer file.Close()

	moveToScore := map[string]int{
		"X": 1,
		"Y": 2,
		"Z": 3,
	}
	_ = moveToScore

	desiredOutcomeToScore := map[string]int{
		"X": 0,
		"Y": 3,
		"Z": 6,
	}

	totalScore := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lineSeparated := strings.Split(line, " ")
		opMove := lineSeparated[0]
		desiredOutcome := lineSeparated[1]

		totalScore += desiredOutcomeToScore[desiredOutcome]
		totalScore += getScoreFromDesiredOutcome(desiredOutcome, opMove)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal("Error reading lines:", err)
	}
	fmt.Println("res:", totalScore)
}

// Moves: rock, paper, scissors
// Rules:
// rock beats scrissors, lose to paper
// paper beats rock, lose to scissors
// scissors beats paper, lose to rock
//
// OP moves:
// A = Rock
// B = Paper
// C = Scissors
func getScoreFromMoves(yourMove, opMove string) int {
	switch yourMove {
	case "X": // rock
		if opMove == "A" { // rock
			return 3
		} else if opMove == "B" { // paper
			return 0
		} else { // scissors
			return 6
		}
	case "Y": // paper
		if opMove == "A" { // rock
			return 6
		} else if opMove == "B" { // paper
			return 3
		} else { // scissors
			return 0
		}
	case "Z": // scissors
		if opMove == "A" { // rock
			return 0
		} else if opMove == "B" { // paper
			return 6
		} else { // scissors
			return 3
		}
	}

	return 0
}

// Desired outcome:
// X = You must lose
// Y = You must draw
// Z = You must win
//
// OP moves:
// A = Rock
// B = Paper
// C = Scissors
//
// Shape score:
// Rock = 1
// Paper = 2
// Scissors = 3
func getScoreFromDesiredOutcome(desiredOutcome, opMove string) int {
	switch desiredOutcome {
	case "X":
		if opMove == "A" {
			return 3
		} else if opMove == "B" {
			return 1
		} else {
			return 2
		}
	case "Y":
		if opMove == "A" {
			return 1
		} else if opMove == "B" {
			return 2
		} else {
			return 3
		}
	case "Z":
		if opMove == "A" {
			return 2
		} else if opMove == "B" {
			return 3
		} else {
			return 1
		}
	}
	return 0
}
