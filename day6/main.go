package main

import (
	"bufio"
	"log"
	"os"
)

func getInput(filename string) *[][]rune {
	result := new([][]rune)

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	isNew := true
	group := new([]rune)
	for scanner.Scan() {
		if scanner.Text() == "" {
			*result = append(*result, *group)
			group = new([]rune)
			isNew = true
		} else if isNew {
			for _, r := range scanner.Text() {
				*group = append(*group, r)
			}
			isNew = false
		} else {
			tempGroup := new([]rune)
			for _, r := range *group {
				for _, s := range scanner.Text() {
					if r == s {
						*tempGroup = append(*tempGroup, r)
					}
				}
			}
			*group = *tempGroup
		}
	}

	return result
}

func main() {
	input := getInput(os.Args[1])
	sum := 0
	for _, group := range *input {
		sum += len(group)
	}

	log.Printf("Sum of applicable questions: %d", sum)
}
