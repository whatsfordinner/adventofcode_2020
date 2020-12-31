package main

import (
	"bufio"
	"log"
	"os"
)

type cell struct {
	x int
	y int
	z int
	w int
}

func (c *cell) isAdjacent(other *cell) bool {
	if c.x == other.x && c.y == other.y && c.z == other.z && c.w == other.w {
		return false
	}

	if absDiff(c.x, other.x) > 1 {
		return false
	}

	if absDiff(c.y, other.y) > 1 {
		return false
	}

	if absDiff(c.z, other.z) > 1 {
		return false
	}

	if absDiff(c.w, other.w) > 1 {
		return false
	}

	return true
}

func absDiff(a int, b int) int {
	result := a - b
	if result < 0 {
		return result * -1
	}

	return result
}

func addIfNotPresent(cells *[]*cell, cell *cell) {
	for _, c := range *cells {
		if c.x == cell.x && c.y == cell.y && c.z == cell.z && c.w == cell.w {
			return
		}
	}

	*cells = append(*cells, cell)
}

func getSurroundingCells(cells *[]*cell, input cell) *[]*cell {
	result := new([]*cell)

	for _, c := range *cells {
		if input.isAdjacent(c) {
			*result = append(*result, c)
		}
	}

	return result
}

func tick(input *[]*cell) *[]*cell {
	result := new([]*cell)

	for _, c := range *input {
		for x := c.x - 1; x <= c.x+1; x++ {
			for y := c.y - 1; y <= c.y+1; y++ {
				for z := c.z - 1; z <= c.z+1; z++ {
					for w := c.w - 1; w <= c.w+1; w++ {
						adjacentCells := getSurroundingCells(input, cell{x, y, z, w})
						if x == c.x && y == c.y && z == c.z && w == c.w {
							if len(*adjacentCells) == 2 {
								addIfNotPresent(result, c)
							}
						}
						if len(*adjacentCells) == 3 {
							addIfNotPresent(result, &cell{x, y, z, w})
						}
					}
				}
			}
		}
	}

	return result
}

func getInput(filename string) *[]*cell {
	result := new([]*cell)

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	y, z, w := 0, 0, 0
	for scanner.Scan() {
		for x, v := range scanner.Text() {
			if v == '#' {
				*result = append(*result, &cell{x, y, z, w})
			}
		}
		y++
	}

	return result
}

func main() {
	input := getInput(os.Args[1])
	for _, v := range *input {
		log.Printf("%+v", *v)
	}

	for i := 0; i < 6; i++ {
		input = tick(input)
	}

	log.Printf("%d", len(*input))
}
