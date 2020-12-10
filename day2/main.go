package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

type password struct {
	requiredChar rune
	firstChar    int
	secondChar   int
	password     string
}

func (p *password) IsValid() bool {
	count := 0

	for _, j := range p.password {
		if j == p.requiredChar {
			count++
		}
	}

	return []rune(p.password)[p.firstChar] == p.requiredChar && []rune(p.password)[p.secondChar] != p.requiredChar || []rune(p.password)[p.firstChar] != p.requiredChar && []rune(p.password)[p.secondChar] == p.requiredChar
}

func getInput(filename string) *[]password {
	result := new([]password)

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if scanner.Text() != "" {
			line := scanner.Text()
			newPassword := new(password)
			tokens := strings.Split(line, " ")

			newPassword.password = tokens[2]

			rangeTokens := strings.Split(tokens[0], "-")
			newFirst, err := strconv.Atoi(rangeTokens[0])
			if err != nil {
				log.Fatal(err)
			}
			newPassword.firstChar = newFirst - 1
			newSecond, err := strconv.Atoi(rangeTokens[1])
			if err != nil {
				log.Fatal(err)
			}
			newPassword.secondChar = newSecond - 1

			newPassword.requiredChar = []rune(tokens[1])[0]

			*result = append(*result, *newPassword)
		}
	}

	return result
}

func countValidPasswords(passwords *[]password) int {
	count := 0

	for _, password := range *passwords {
		if password.IsValid() {
			count++
		}
	}

	return count
}

func main() {
	inputFile := os.Args[1]
	input := getInput(inputFile)

	log.Printf("Valid passwords: %d", countValidPasswords(input))
}
