package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal("Error opening file:", err)
	}
	defer file.Close()

	var total int
	var iter int
	var groupCharMaps []map[rune]bool
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		iter += 1

		// construct map
		m := make(map[rune]bool)
		for _, r := range line {
			if _, isExist := m[r]; !isExist {
				m[r] = true
			}
		}
		groupCharMaps = append(groupCharMaps, m)

		// group is formed, find the common character and get the priority value
		if iter == 3 && len(groupCharMaps) == 3 {
			pivot := groupCharMaps[0]
			var commonChar rune
			for pivotChar := range pivot {
				_, isExistInSecond := groupCharMaps[1][pivotChar]
				_, isExistInThird := groupCharMaps[2][pivotChar]
				if isExistInSecond && isExistInThird {
					commonChar = pivotChar
				}
			}

			total += charToPrioValue(commonChar)

			// reset group
			iter = 0
			groupCharMaps = []map[rune]bool{}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal("Error reading lines:", err)
	}

	fmt.Println("res:", total)
}

func charToPrioValue(c rune) int {
	if c >= 'a' && c <= 'z' {
		return int(c-'a') + 1
	} else if c >= 'A' && c <= 'Z' {
		return int(c-'A') + 27
	}
	return 0
}
