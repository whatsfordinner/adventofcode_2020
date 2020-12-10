package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type passport struct {
	fields map[string]string
}

func (p *passport) isValid() bool {
	if p.fields["byr"] != "" {
		birthYear, err := strconv.Atoi(p.fields["byr"])
		if err != nil {
			return false
		}

		if birthYear < 1920 || birthYear > 2002 {
			return false
		}

	} else {
		return false
	}

	if p.fields["iyr"] != "" {
		issueYear, err := strconv.Atoi(p.fields["iyr"])
		if err != nil {
			return false
		}

		if issueYear < 2010 || issueYear > 2020 {
			return false
		}

	} else {
		return false
	}

	if p.fields["eyr"] != "" {
		expiryYear, err := strconv.Atoi(p.fields["eyr"])
		if err != nil {
			return false
		}

		if expiryYear < 2020 || expiryYear > 2030 {
			return false
		}

	} else {
		return false
	}

	if p.fields["hgt"] != "" {
		units := p.fields["hgt"][len(p.fields["hgt"])-2:]
		valueStr := p.fields["hgt"][:len(p.fields["hgt"])-2]

		value, err := strconv.Atoi(valueStr)
		if err != nil {
			return false
		}

		if units == "cm" {
			if value < 150 || value > 193 {
				return false
			}
		} else if units == "in" {
			if value < 59 || value > 76 {
				return false
			}
		} else {
			return false
		}
	} else {
		return false
	}

	if p.fields["hcl"] != "" {
		re := regexp.MustCompile(`^#[0-9a-f]{6}$`)
		if !re.MatchString(p.fields["hcl"]) {
			return false
		}
	} else {
		return false
	}

	if p.fields["ecl"] != "" {
		eyeColour := p.fields["ecl"]
		if eyeColour != "amb" && eyeColour != "blu" && eyeColour != "brn" && eyeColour != "gry" && eyeColour != "grn" && eyeColour != "hzl" && eyeColour != "oth" {
			return false
		}
	} else {
		return false
	}

	if p.fields["pid"] != "" {
		re := regexp.MustCompile(`^\d{9}$`)
		if !re.MatchString(p.fields["pid"]) {
			return false
		}
	} else {
		return false
	}

	return true
}

func countValid(passports *[]passport) int {
	count := 0

	for _, p := range *passports {
		if p.isValid() {
			count++
		}
	}

	return count
}

func getInput(filename string) *[]passport {
	result := new([]passport)

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	newPassport := new(passport)
	newPassport.fields = make(map[string]string)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			*result = append(*result, *newPassport)
			newPassport = new(passport)
			newPassport.fields = make(map[string]string)
		} else {
			entries := strings.Split(line, " ")
			for _, entry := range entries {
				tokens := strings.Split(entry, ":")
				newPassport.fields[tokens[0]] = tokens[1]
			}
		}

	}

	*result = append(*result, *newPassport)

	return result
}

func main() {
	inputFile := os.Args[1]
	input := getInput(inputFile)
	log.Printf("Result: %d", countValid(input))
}
