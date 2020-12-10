package main

import (
	"bufio"
	"log"
	"os"
)

func getInput(filename string) *[]string {
	result := new([]string)

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if scanner.Text() != "" {
			*result = append(*result, scanner.Text())
		}
	}

	return result
}

func countTrees(right int, down int, slope *[]string) int {
	count := 0
	position := 0

	for i := 0; i < len(*slope); i += down {
		if (*slope)[i][position] == '#' {
			count++
		}

		position = (position + right) % len((*slope)[i])
	}

	log.Printf("Angle of (%d,%d): %d", right, down, count)
	return count
}

func main() {
	inputFile := os.Args[1]
	input := getInput(inputFile)
	count := countTrees(1, 1, input)
	count *= countTrees(3, 1, input)
	count *= countTrees(5, 1, input)
	count *= countTrees(7, 1, input)
	count *= countTrees(1, 2, input)
	log.Printf("Result: %d", count)
}
