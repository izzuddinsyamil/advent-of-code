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

	slopes := [][]int{
		{1, 1},
		{3, 1},
		{5, 1},
		{7, 1},
		{1, 2},
	}
	_ = slopes

	res := 1
	for _, slope := range slopes {
		var r, c, treeCount int

		for r < len(grid) {
			for c >= len(grid[r]) {
				//expand the grid
				expandRow(r, grid)
			}

			// fmt.Println(r, c, len(grid[r]), grid[r])
			if grid[r][c] == "#" {
				treeCount += 1
			}
			c += slope[0]
			r += slope[1]
		}

		res *= treeCount
	}

	fmt.Println(res)
}

func expandRow(r int, grid [][]string) {
	grid[r] = append(grid[r], grid[r]...)
}
