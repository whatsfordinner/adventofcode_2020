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
		for j, c := range r {
			if c.evaluate(f.getAdjacent(i, j)) {
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
				*result = append(*result, f.spots[r][c])
			}
		}
	}

	return *result
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
	log.Print(floor.toString())
	floor.evaluate()
	floor.tick()
	log.Print(floor.toString())
	floor.evaluate()
	floor.tick()
	log.Print(floor.toString())

}
