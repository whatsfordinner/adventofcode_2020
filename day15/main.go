package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func getInput(filename string) (map[int]*[]int, int) {
	result := make(map[int]*[]int)
	lastNumber := 0

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		for i, n := range strings.Split(scanner.Text(), ",") {
			in, err := strconv.Atoi(n)
			if err != nil {
				log.Fatal(err)
			}

			if result[in] == nil {
				result[in] = new([]int)
			}

			*result[in] = append(*result[in], i+1)
			lastNumber = in
		}
	}

	return result, lastNumber
}

func main() {
	game, lastNumber := getInput(os.Args[1])

	for i := len(game) + 1; i <= 30000000; i++ {
		if len(*game[lastNumber]) <= 1 {
			if game[0] == nil {
				game[0] = new([]int)
			}
			*game[0] = append(*game[0], i)
			lastNumber = 0
		} else {
			newNumber := (*game[lastNumber])[len(*game[lastNumber])-1] - (*game[lastNumber])[len(*game[lastNumber])-2]
			if game[newNumber] == nil {
				game[newNumber] = new([]int)
			}

			*game[newNumber] = append(*game[newNumber], i)
			lastNumber = newNumber
		}

	}

	log.Printf("%d", lastNumber)
}
