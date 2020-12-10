package main

import (
	"bufio"
	"log"
	"os"
	"sort"
)

type boardingPass struct {
	code   string
	row    int
	column int
	uid    int
}

func (b *boardingPass) decode() {
	b.row = findPos(0, 127, b.code[:7])
	b.column = findPos(0, 7, b.code[7:])
	b.uid = b.row*8 + b.column
}

func findPos(min int, max int, code string) int {
	if len(code) == 1 {
		if code[0] == 'F' || code[0] == 'L' {
			return min
		}
		return max
	}

	diff := (max - min + 1) / 2

	if code[0] == 'F' || code[0] == 'L' {
		return findPos(min, max-diff, code[1:])
	}
	return findPos(min+diff, max, code[1:])
}

func sortByUID(bps []boardingPass) []boardingPass {
	sort.Slice(bps, func(i int, j int) bool {
		return bps[i].uid < bps[j].uid
	})

	return bps
}

func getInput(filename string) *[]boardingPass {
	result := new([]boardingPass)

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		if line != "" {
			newBoardingPass := boardingPass{code: line}
			newBoardingPass.decode()
			*result = append(*result, newBoardingPass)
		}
	}

	return result
}

func main() {
	inputFile := os.Args[1]
	input := getInput(inputFile)

	sorted := sortByUID(*input)

	last := 0
	current := 0

	for _, i := range sorted {
		current = i.uid

		if last != 0 {
			if current-last > 1 {
				log.Printf("Gap found between %d and %d", current, last)
			}
		}
		last = current
	}
}
