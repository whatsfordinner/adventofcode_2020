package main

import (
	"bufio"
	"log"
	"os"
	"sort"
	"strconv"
)

func getInput(filename string) *[]int {
	result := new([]int)

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	*result = append(*result, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if scanner.Text() != "" {
			newNumber, err := strconv.Atoi(scanner.Text())
			if err != nil {
				log.Fatal(err)
			}
			*result = append(*result, newNumber)
		}
	}

	sort.Slice(*result, func(i int, j int) bool {
		return (*result)[i] < (*result)[j]
	})
	return result
}

func makeChain(input *[]int) int {
	oneJumps := 0
	threeJumps := 0
	currentJolts := 0

	for _, i := range *input {
		if i-currentJolts == 3 {
			threeJumps++
		}

		if i-currentJolts == 1 {
			oneJumps++
		}

		currentJolts = i
	}

	threeJumps++

	return oneJumps * threeJumps
}

func breakIntoThrees(input *[]int) *[]*[]int {
	result := new([]*[]int)
	workingSet := new([]int)
	for _, i := range *input {
		if len(*workingSet) == 0 {
			*workingSet = append(*workingSet, i)
		} else if i-(*workingSet)[len(*workingSet)-1] == 3 {
			*result = append(*result, workingSet)
			workingSet = new([]int)
			*workingSet = append(*workingSet, i)
		} else {
			*workingSet = append(*workingSet, i)
		}
	}

	*result = append(*result, workingSet)

	return result
}

func howManyCombos(input []int) int {
	if len(input) <= 2 {
		log.Printf("%+v has 1 combo", input)
		return 1
	}
	combos := 0
	for i := 1; i < len(input); i++ {
		if input[i]-input[0] <= 3 {
			if input[i] == input[len(input)-1] {
				combos++
			} else {
				combos += howManyCombos(input[i:])
			}
		}
	}
	log.Printf("%+v has %d combos", input, combos)
	return combos
}

func main() {
	input := getInput(os.Args[1])
	combos := 1
	for _, c := range *breakIntoThrees(input) {
		combos *= howManyCombos(*c)
	}

	log.Printf("Can be put together %d different ways", combos)
}
