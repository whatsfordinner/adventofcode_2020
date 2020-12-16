package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type field struct {
	name   string
	ranges []validRange
}

func (f field) isValid(input int) bool {
	for _, n := range f.ranges {
		if n.isValid(input) {
			return true
		}
	}

	return false
}

type validRange struct {
	min int
	max int
}

func (v validRange) isValid(input int) bool {
	return input >= v.min && input <= v.max
}

type ticket struct {
	fields     []int
	fieldNames []string
}

func (t ticket) isValid(fields []field) bool {
	for _, f := range t.fields {
		valid := false
		for _, n := range fields {
			if n.isValid(f) {
				valid = true
			}
		}
		if !valid {
			return false
		}
	}
	return true
}

func getInput(filename string) (*[]field, *[]ticket, *ticket) {
	fields := new([]field)
	tickets := new([]ticket)
	myTicket := new(ticket)

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() && scanner.Text() != "" {
		fieldTokens := strings.Split(scanner.Text(), ": ")
		rangeTokens := strings.Split(fieldTokens[1], " or ")
		ranges := new([]validRange)

		for _, n := range rangeTokens {
			newRange := new(validRange)
			rangeStrings := strings.Split(n, "-")
			newMin, err := strconv.Atoi(rangeStrings[0])
			if err != nil {
				log.Fatal(err)
			}
			newMax, err := strconv.Atoi(rangeStrings[1])
			if err != nil {
				log.Fatal(err)
			}
			newRange.min = newMin
			newRange.max = newMax
			*ranges = append(*ranges, *newRange)
		}
		*fields = append(*fields, field{fieldTokens[0], *ranges})
	}

	for scanner.Scan() && scanner.Text() != "" {
		if scanner.Text() != "your ticket:" {
			myFields := strings.Split(scanner.Text(), ",")
			for _, n := range myFields {
				newField, err := strconv.Atoi(n)
				if err != nil {
					log.Fatal(err)
				}
				myTicket.fields = append(myTicket.fields, newField)
			}
			myTicket.fieldNames = make([]string, len(myTicket.fields))
		}
	}

	for scanner.Scan() && scanner.Text() != "" {
		if scanner.Text() != "nearby tickets:" {
			ticketFields := strings.Split(scanner.Text(), ",")
			newTicketValues := make([]int, 0)
			for _, n := range ticketFields {
				newField, err := strconv.Atoi(n)
				if err != nil {
					log.Fatal(err)
				}
				newTicketValues = append(newTicketValues, newField)
			}
			newTicket := ticket{newTicketValues, make([]string, len(newTicketValues))}
			if newTicket.isValid(*fields) {
				*tickets = append(*tickets, newTicket)
			}
		}
	}

	return fields, tickets, myTicket
}

func matchTicketsToFields(tickets []ticket, fields []field) []string {
	result := make([]string, len(fields))
	candidates := make([][]string, len(fields))

	for _, f := range fields {
		for i := 0; i < len(fields); i++ {
			valid := true
			for _, t := range tickets {
				if !f.isValid(t.fields[i]) {
					valid = false
					break
				}
			}

			if valid {
				candidates[i] = append(candidates[i], f.name)
			}
		}
	}

	for {
		removedCandidate := false
		for i, n := range candidates {
			if len(n) == 1 {
				result[i] = candidates[i][0]
				candidates = removeField(candidates, candidates[i][0])
				removedCandidate = true
			}
		}

		if !removedCandidate {
			break
		}
	}

	return result
}

func removeField(c [][]string, f string) [][]string {
	for i, n := range c {
		for j, o := range n {
			if o == f {
				c[i] = append(n[:j], n[j+1:]...)
				break
			}
		}
	}

	return c
}

func main() {
	fields, validTickets, myTicket := getInput(os.Args[1])
	myTicket.fieldNames = matchTicketsToFields(*validTickets, *fields)
	re := regexp.MustCompile(`departure.*`)
	result := 1
	for i, n := range myTicket.fieldNames {
		if re.Match([]byte(n)) {
			result *= myTicket.fields[i]
		}
	}

	log.Printf("Result: %d", result)
}
