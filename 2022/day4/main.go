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

		// part 1: check if first contains second
		// if first[0] <= second[0] && first[1] >= second[1] {
		// 	// fmt.Println("first contains second, pair:", pairs)
		// 	res += 1
		// } else if second[0] <= first[0] && second[1] >= first[1] { // check if second contains first
		// 	// fmt.Println("second contains first, pair:", pairs)
		// 	res += 1
		// }

		// part 2: check overlapping item
		// reorder the first and second.
		// any section with the first element < that the other, becomes the first
		temp := []int{}
		if second[0] < first[0] {
			temp = first
			first = second
			second = temp
		}

		// In the above example, the first two pairs (2-4,6-8 and 2-3,4-5) don't overlap, while the remaining four pairs (5-7,7-9, 2-8,3-7, 6-6,4-6, and 2-6,4-8) do overlap:
		// 5-7,7-9 overlaps in a single section, 7.
		// 2-8,3-7 overlaps all of the sections 3 through 7.
		// 6-6,4-6, reorderd into: 4-6,6-6 overlaps in a single section, 6.
		// 2-6,4-8 overlaps in sections 4, 5, and 6.

		// check if the first part of the second overlaps with first's area
		fmt.Println("checking", first, second)
		if second[0] >= first[0] && second[0] <= first[1] {
			fmt.Println("found overlapping", first, second)
			res += 1
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal("Error reading lines:", err)
	}

	fmt.Println("res:", res)
}
