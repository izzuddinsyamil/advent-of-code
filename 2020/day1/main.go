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
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var inp []int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		i, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		inp = append(inp, i)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	sort.Ints(inp)

	target := 2020
	l := 0
	r := len(inp) - 1
	var num1, num2 int

	for l < r {
		if inp[l]+inp[r] == target {
			num1 = inp[l]
			num2 = inp[r]
			break
		}

		if inp[l]+inp[r] < target {
			l += 1
			continue
		}

		if inp[l]+inp[r] > target {
			r -= 1
			continue
		}
	}

	fmt.Println(num1, num2)
	fmt.Println(num1 * num2)
}
