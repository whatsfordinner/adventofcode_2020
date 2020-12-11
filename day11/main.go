package main

import (
	"bufio"
	"log"
	"os"
)

type floor struct {
	spots [][]*spot
}

func (f *floor) evaluate() bool {
	changes := false
	for i, r := range f.spots {
		for j := range r {
			if f.evaluateLineOfSight(i, j) {
				changes = true
			}
		}
	}

	return changes
}

func (f *floor) tick() {
	for _, r := range f.spots {
		for _, c := range r {
			c.tick()
		}
	}
}

func (f *floor) getAdjacent(r int, c int) []*spot {
	result := new([]*spot)
	var startRow int
	var finRow int
	var startColumn int
	var finColumn int

	if r == 0 {
		startRow = 0
	} else {
		startRow = r - 1
	}

	if r == len(f.spots)-1 {
		finRow = r
	} else {
		finRow = r + 1
	}

	if c == 0 {
		startColumn = 0
	} else {
		startColumn = c - 1
	}

	if c == len(f.spots[r])-1 {
		finColumn = c
	} else {
		finColumn = c + 1
	}

	for i := startRow; i <= finRow; i++ {
		for j := startColumn; j <= finColumn; j++ {
			if !(i == r && j == c) {
				*result = append(*result, f.spots[i][j])
			}
		}
	}

	return *result
}

func (f *floor) evaluateLineOfSight(r int, c int) bool {
	occupiedCount := 0
	// Look Left
	for i := c - 1; i >= 0; i-- {
		if f.spots[r][i].currentState == "#" {
			occupiedCount++
			break
		}

		if f.spots[r][i].currentState == "L" {
			break
		}
	}

	// Look Right
	for i := c + 1; i < len(f.spots[r]); i++ {
		if f.spots[r][i].currentState == "#" {
			occupiedCount++
			break
		}

		if f.spots[r][i].currentState == "L" {
			break
		}
	}

	// Look Up
	for i := r - 1; i >= 0; i-- {
		if f.spots[i][c].currentState == "#" {
			occupiedCount++
			break
		}

		if f.spots[i][c].currentState == "L" {
			break
		}
	}

	// Look Down
	for i := r + 1; i < len(f.spots); i++ {
		if f.spots[i][c].currentState == "#" {
			occupiedCount++
			break
		}

		if f.spots[i][c].currentState == "L" {
			break
		}
	}
	// Look Up-Left
	for i, j := r-1, c-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if f.spots[i][j].currentState == "#" {
			occupiedCount++
			break
		}

		if f.spots[i][j].currentState == "L" {
			break
		}
	}

	// Look Up-Right
	for i, j := r-1, c+1; i >= 0 && j < len(f.spots[r]); i, j = i-1, j+1 {
		if f.spots[i][j].currentState == "#" {
			occupiedCount++
			break
		}

		if f.spots[i][j].currentState == "L" {
			break
		}
	}

	// Look Down-Left
	for i, j := r+1, c-1; i < len(f.spots) && j >= 0; i, j = i+1, j-1 {
		if f.spots[i][j].currentState == "#" {
			occupiedCount++
			break
		}

		if f.spots[i][j].currentState == "L" {
			break
		}
	}

	// Look Down-Right
	for i, j := r+1, c+1; i < len(f.spots) && j < len(f.spots[r]); i, j = i+1, j+1 {
		if f.spots[i][j].currentState == "#" {
			occupiedCount++
			break
		}

		if f.spots[i][j].currentState == "L" {
			break
		}
	}

	if f.spots[r][c].currentState == "#" && occupiedCount >= 5 {
		f.spots[r][c].nextState = "L"
		return true
	}

	if f.spots[r][c].currentState == "L" && occupiedCount == 0 {
		f.spots[r][c].nextState = "#"
		return true
	}

	if f.spots[r][c].currentState == "." {
		f.spots[r][c].nextState = "."
	}

	return false

}

func (f *floor) countOccupied() int {
	count := 0

	for _, r := range f.spots {
		for _, c := range r {
			if c.currentState == "#" {
				count++
			}
		}
	}

	return count
}

func (f *floor) toString() string {
	output := "\n"
	for _, r := range f.spots {
		for _, c := range r {
			output += c.currentState
		}
		output += "\n"
	}

	return output
}

type spot struct {
	currentState string
	nextState    string
}

func (s *spot) evaluate(adjacent []*spot) bool {
	if s.currentState == "L" {
		for _, n := range adjacent {
			if n.currentState == "#" {
				s.nextState = "L"
				return false
			}
		}
		s.nextState = "#"
		return true
	}

	if s.currentState == "#" {
		count := 0
		for _, n := range adjacent {
			if n.currentState == "#" {
				count++
			}
		}

		if count >= 4 {
			s.nextState = "L"
			return true
		}
	}

	if s.currentState == "." {
		s.nextState = "."
	}

	return false
}

func (s *spot) tick() {
	s.currentState = s.nextState
}

func getInput(filename string) *floor {
	result := new(floor)
	spots := new([][]*spot)

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if scanner.Text() != "" {
			newRow := new([]*spot)
			for _, s := range scanner.Text() {
				newSpot := new(spot)
				newSpot.currentState = string(s)
				*newRow = append(*newRow, newSpot)
			}
			*spots = append(*spots, *newRow)
		}
	}

	result.spots = *spots
	return result
}

func main() {
	floor := getInput(os.Args[1])
	iters := 0
	for {
		if !floor.evaluate() {
			break
		}
		iters++
		floor.tick()
	}
	log.Printf("Stable after %d iterations with %d seats occupied", iters, floor.countOccupied())
}
