package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func main() {
	inputFile := os.Args[1]
	input := getInput(inputFile)
	log.Printf("Result: %d", findPair(input))
}

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
			entry, err := strconv.Atoi(scanner.Text())
			if err != nil {
				log.Fatal(err)
			}
			*result = append(*result, entry)
		}
	}

	return result
}

func findPair(input *[]int) int {
	for i := 0; i < len(*input); i++ {
		for j := i + 1; j < len(*input); j++ {
			for k := j + 1; k < len(*input); k++ {
				if (*input)[i]+(*input)[j]+(*input)[k] == 2020 {
					return (*input)[i] * (*input)[j] * (*input)[k]
				}
			}
		}
	}

	return 0
}
