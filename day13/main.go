package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func getInput(filename string) (int, *[]int) {
	result := new([]int)
	input := new([]string)

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

	timestamp, err := strconv.Atoi((*input)[0])
	if err != nil {
		log.Fatal(err)
	}

	for _, s := range strings.Split((*input)[1], ",") {
		if s != "x" {
			busID, err := strconv.Atoi(s)
			if err != nil {
				log.Fatal(err)
			}
			*result = append(*result, busID)
		}
	}

	return timestamp, result
}

func timeUntil(timestamp int, busID int) int {
	// timestamp % busID gives the time since the last bus arrived
	return busID - (timestamp % busID)
}

func main() {
	timestamp, busIDs := getInput(os.Args[1])
	log.Printf("Timestamp: %d", timestamp)
	log.Printf("Bus IDs: %+v", *busIDs)
	for _, b := range *busIDs {
		busID := b
		time := timeUntil(timestamp, b)
		log.Printf("Bus ID: %d, Time until it arrives: %d. Result: %d", busID, time, busID*time)
	}
}
