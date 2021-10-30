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
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var count int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		inp := processLine(line)

		if processInput(inp) {
			count += 1
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(count)

}

type inputLine struct {
	min           int
	max           int
	charToInclude string
	password      string
}

func processLine(line string) inputLine {
	lineColonSplit := strings.Split(line, ":")
	pwd := lineColonSplit[1]

	ruleSplit := strings.Split(lineColonSplit[0], " ")
	charToInclude := ruleSplit[1]

	minMax := strings.Split(ruleSplit[0], "-")
	min, _ := strconv.Atoi(minMax[0])
	max, _ := strconv.Atoi(minMax[1])

	return inputLine{
		min:           min,
		max:           max,
		charToInclude: charToInclude,
		password:      pwd,
	}

}

func processInput(inp inputLine) bool {

	if !strings.Contains(inp.password, inp.charToInclude) {
		return false
	}

	charToIncludeCount := strings.Count(inp.password, inp.charToInclude)
	if charToIncludeCount < inp.min {
		return false
	}

	if charToIncludeCount > inp.max {
		return false
	}

	return true
}
