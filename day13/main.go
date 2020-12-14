package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func getInput(filename string) *[]int {
	input := new([]string)
	result := new([]int)

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if scanner.Text() != "" {
			*input = append(*input, scanner.Text())
		}
	}

	for _, s := range strings.Split((*input)[0], ",") {
		if s == "x" {
			*result = append(*result, 0)
		} else {
			busID, err := strconv.Atoi(s)
			if err != nil {
				log.Fatal(err)
			}
			*result = append(*result, busID)
		}
	}

	return result
}

func findLargestID(input *[]int) (int, int) {
	index := 0
	result := 0

	for i, n := range *input {
		if n > result {
			index = i
			result = n
		}
	}

	return index, result
}

func main() {
	busIDs := getInput(os.Args[1])
	log.Printf("Bus IDs: %+v", *busIDs)

	increment := 1
	syncedBuses := new([]int)
	for ts := 1; ; ts += increment {
		isCandidate := true
		for i, n := range *busIDs {
			if n != 0 {
				if (ts+i)%n == 0 {
					alreadySynced := false
					for _, b := range *syncedBuses {
						if b == n {
							alreadySynced = true
							break
						}
					}
					if !alreadySynced {
						log.Printf("Adding %d to list of synced buses", n)
						increment *= n
						*syncedBuses = append(*syncedBuses, n)
					}
				} else {
					isCandidate = false
				}
			}
		}

		if isCandidate {
			log.Printf("Found candidate: %d", ts)
			break
		}
	}
}
