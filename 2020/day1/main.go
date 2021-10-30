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
	var num1, num2, num3 int

	for i := 0; i < len(inp)-2; i++ {
		l := i + 1
		r := len(inp) - 1

		for l < r {
			su := inp[i] + inp[l] + inp[r]
			if su == target {
				num1 = inp[i]
				num2 = inp[l]
				num3 = inp[r]
				break
			}

			if su < target {
				l += 1
				continue
			}

			if su > target {
				r -= 1
				continue
			}
		}
	}

	fmt.Println(num1, num2, num3)
	fmt.Println(num1 * num2 * num3)
}
