package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	max := []int{}
	currVal := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			fmt.Println("currVal", currVal)
			// check if the current value can be included in the top 3
			max = checkTopThreeMaxes(max, currVal)

			// reset curret value for another set
			currVal = 0
		} else {
			val, _ := strconv.Atoi(line)
			currVal += val
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading lines:", err)
	}

	var res int
	for _, m := range max {
		res += m
	}
	fmt.Println("res: ", res)
}

func checkTopThreeMaxes(currMax []int, currVal int) []int {
	if len(currMax) < 3 {
		currMax = append(currMax, currVal)
		return currMax
	}
	sort.Ints(currMax)
	for i, m := range currMax {
		if currVal > m {
			currMax[i] = currVal
			break
		}
	}
	return currMax
}
