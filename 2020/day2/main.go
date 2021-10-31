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
	var inp []inputLine
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		i := processLine(scanner.Text())
		inp = append(inp, i)

	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	for _, i := range inp {
		processRes := processInput(i)

		if processRes {
			count += 1
		} else {
			fmt.Printf("inp : %+v, res : %t\n", i, processRes)
		}
	}
	fmt.Println(count)
}

type inputLine struct {
	pos1          int
	pos2          int
	charToInclude string
	password      string
}

func processLine(line string) inputLine {
	lineColonSplit := strings.Split(line, ":")
	pwd := strings.TrimSpace(lineColonSplit[1])

	ruleSplit := strings.Split(lineColonSplit[0], " ")
	charToInclude := ruleSplit[1]

	position := strings.Split(ruleSplit[0], "-")
	pos1, _ := strconv.Atoi(position[0])
	pos2, _ := strconv.Atoi(position[1])

	return inputLine{
		pos1:          pos1 - 1,
		pos2:          pos2 - 1,
		charToInclude: charToInclude,
		password:      pwd,
	}
}

func processInput(inp inputLine) bool {
	pos1Char := inp.password[inp.pos1 : inp.pos1+1]
	pos2Char := inp.password[inp.pos2 : inp.pos2+1]

	if pos1Char != inp.charToInclude && pos2Char != inp.charToInclude {
		return false
	}

	if pos1Char == inp.charToInclude && pos2Char == inp.charToInclude {
		return false
	}

	return true
}
