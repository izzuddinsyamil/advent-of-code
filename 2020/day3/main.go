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
		log.Fatal(err)
	}
	defer file.Close()

	var grid [][]string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		grid = append(grid, strings.Split(scanner.Text(), ""))

	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	var r, c, treeCount int
	for r < len(grid) {
		for c > len(grid[r]) {
			//expand the grid
			expandRow(r, grid)
		}

		if grid[r][c] == "#" {
			treeCount += 1
		}
		c += 3
		r += 1
	}

	fmt.Println(treeCount)
}

func expandRow(r int, grid [][]string) {
	grid[r] = append(grid[r], grid[r]...)
}
