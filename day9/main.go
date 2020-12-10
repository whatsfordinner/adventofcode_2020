package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func getInput(filename string) *[]int {
	result := new([]int)

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if scanner.Text() != "" {
			newInt, err := strconv.Atoi(scanner.Text())
			if err != nil {
				log.Fatal(err)
			}

			*result = append(*result, newInt)
		}
	}

	return result
}

func findVulnerability(input *[]int, preambleLength int) int {
	for i := preambleLength; i < len(*input); i++ {
		testNumber := (*input)[i]
		testPreamble := (*input)[i-preambleLength : i]

		if !isValidNumber(testNumber, testPreamble) {
			return testNumber
		}
	}

	return 0
}

func isValidNumber(input int, preamble []int) bool {
	for i := 0; i < len(preamble)-1; i++ {
		for j := i + 1; j < len(preamble); j++ {
			if preamble[i]+preamble[j] == input {
				return true
			}
		}
	}
	return false
}

func findContiguousList(input *[]int, vulnerableNumber int) *[]int {
	for i := 0; i < len(*input); i++ {
		testList := new([]int)
		testSum := 0
		for j := i; j < len(*input); j++ {
			*testList = append(*testList, (*input)[j])
			testSum += (*input)[j]

			if testSum == vulnerableNumber {
				return testList
			}
		}
	}

	return nil
}

func findMin(input *[]int) int {
	min := (*input)[0]
	for _, i := range *input {
		if i < min {
			min = i
		}
	}

	return min
}

func findMax(input *[]int) int {
	max := (*input)[0]
	for _, i := range *input {
		if i > max {
			max = i
		}
	}

	return max
}

func main() {
	input := getInput(os.Args[1])
	preambleLength, err := strconv.Atoi(os.Args[2])
	if err != nil {
		log.Fatal(err)
	}

	vulnerableNumber := findVulnerability(input, preambleLength)
	contiguousList := findContiguousList(input, vulnerableNumber)
	log.Printf("List: %v", (*contiguousList))
	log.Printf("Sum: %d", findMin(contiguousList)+findMax(contiguousList))
}
