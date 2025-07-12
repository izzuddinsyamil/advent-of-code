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
		log.Fatal("Error opening file:", err)
	}
	defer file.Close()

	var res int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		pairs := strings.Split(line, ",")
		if len(pairs) != 2 {
			fmt.Println("wrong input")
			return
		}

		firstSplit := strings.Split(pairs[0], "-")
		var first []int
		for _, f := range firstSplit {
			i, _ := strconv.Atoi(f)
			first = append(first, i)
		}
		secondSplit := strings.Split(pairs[1], "-")
		var second []int
		for _, s := range secondSplit {
			i, _ := strconv.Atoi(s)
			second = append(second, i)
		}

		// check if first contains second
		if first[0] <= second[0] && first[1] >= second[1] {
			// fmt.Println("first contains second, pair:", pairs)
			res += 1
		} else if second[0] <= first[0] && second[1] >= first[1] { // check if second contains first
			// fmt.Println("second contains first, pair:", pairs)
			res += 1
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal("Error reading lines:", err)
	}

	fmt.Println("res:", res)
}
