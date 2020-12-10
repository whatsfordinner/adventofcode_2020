package main

import (
	"bufio"
	"errors"
	"log"
	"os"
	"strconv"
	"strings"
)

type instruction struct {
	instruction   string
	argument      int
	timesExecuted int
}

type handheld struct {
	accumulator        int
	currentInstruction int
	instructions       []*instruction
}

func (h *handheld) process() error {
	i := h.instructions[h.currentInstruction]
	if i.timesExecuted > 0 {
		return errors.New("encountered repeat instruction")
	}

	i.timesExecuted++

	if i.instruction == "acc" {
		h.accumulator += i.argument
		h.currentInstruction++
	}

	if i.instruction == "jmp" {
		h.currentInstruction += i.argument
	}

	if i.instruction == "nop" {
		h.currentInstruction++
	}

	return nil
}

func getInput(filename string) *[]*instruction {
	result := new([]*instruction)

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if scanner.Text() != "" {
			*result = append(*result, getInstructionFromString(scanner.Text()))
		}
	}

	return result
}

func getInstructionFromString(s string) *instruction {
	tokens := strings.Split(s, " ")
	argument, err := strconv.Atoi(tokens[1])
	if err != nil {
		log.Fatal(err)
	}

	return &instruction{
		instruction:   tokens[0],
		argument:      argument,
		timesExecuted: 0,
	}
}

func updateInstructionSet(set []*instruction, toChange int) {
	count := 0
	for j, i := range set {
		if i.instruction == "nop" && count == toChange {
			log.Printf("Updating instruction %d from nop to jmp", j)
			i.instruction = "jmp"
			break
		} else if i.instruction == "jmp" && count == toChange {
			log.Printf("Updating instruction %d from jmp to nop", j)
			i.instruction = "nop"
			break
		} else {
			count++
		}
	}
}

func main() {
	instructionSet := getInput(os.Args[1])
	theHandheld := handheld{
		accumulator:        0,
		currentInstruction: 0,
		instructions:       *instructionSet,
	}
	toChange := 0

	for theHandheld.currentInstruction < len(theHandheld.instructions) {
		err := theHandheld.process()
		if err != nil {
			theHandheld.accumulator = 0
			theHandheld.currentInstruction = 0
			theHandheld.instructions = *getInput(os.Args[1])
			updateInstructionSet(theHandheld.instructions, toChange)
			toChange++
		}
	}

	log.Printf("Execution complete. Accumulator: %d", theHandheld.accumulator)
}
