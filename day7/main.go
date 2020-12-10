package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

type bagValue struct {
	holds []struct {
		quantity int
		key      string
	}
}

func (b *bagValue) contains(k string) bool {
	for _, c := range b.holds {
		if c.key == k {
			return true
		}
	}

	return false
}

func (b *bagValue) bagCount(rules map[string]*bagValue) int {
	if len(b.holds) == 0 {
		return 0
	}

	count := 0

	for _, k := range b.holds {
		count += k.quantity
		count += k.quantity * rules[k.key].bagCount(rules)
	}

	return count
}

func bagValueFromString(bagString string) (string, *bagValue) {
	newValue := new(bagValue)
	bagTokens := strings.Split(bagString, " bags contain ")
	newKey := bagTokens[0]

	if bagTokens[1] != "no other bags." {
		containsTokens := strings.Split(bagTokens[1], ", ")
		for _, t := range containsTokens {
			containTokens := strings.Split(t, " ")
			quantity, err := strconv.Atoi(containTokens[0])
			if err != nil {
				log.Fatal(err)
			}
			description := strings.Join(containTokens[1:3], " ")
			newValue.holds = append(newValue.holds, struct {
				quantity int
				key      string
			}{quantity, description})
		}
	}

	return newKey, newValue
}

func getInputMap(filename string) map[string]*bagValue {
	result := make(map[string]*bagValue)

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if scanner.Text() != "" {
			newKey, newValue := bagValueFromString(scanner.Text())
			result[newKey] = newValue
		}
	}

	return result
}

func contains(slice []string, s string) bool {
	for _, e := range slice {
		if e == s {
			return true
		}
	}

	return false
}

func main() {
	bags := getInputMap(os.Args[1])
	count := bags["shiny gold"].bagCount(bags)
	log.Printf("Bags inside a shiny gold bag: %d", count)
}
