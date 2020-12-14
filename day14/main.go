package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type bootloader struct {
	mask string
	mem  map[int]int64
}

func (b *bootloader) applyMask(input int64) int64 {
	inputBinary := fmt.Sprintf("%036s", strconv.FormatInt(int64(input), 2))
	outputBinary := ""

	for i, n := range b.mask {
		if n == 'X' {
			outputBinary += string(inputBinary[i])
		}

		if n == '1' {
			outputBinary += "1"
		}

		if n == '0' {
			outputBinary += "0"
		}
	}

	output, err := strconv.ParseInt(outputBinary, 2, 64)
	if err != nil {
		log.Fatal(err)
	}

	return output
}

func (b *bootloader) updateMemory(address int, input int64) {
	b.mem[address] = b.applyMask(input)
}

func (b *bootloader) processInstruction(instruction string) {
	instructionTokens := strings.Split(instruction, " = ")
	if instructionTokens[0] == "mask" {
		b.mask = instructionTokens[1]
	} else {
		re := regexp.MustCompile(`\d+`)
		addressStr := string(re.Find([]byte(instructionTokens[0])))
		address, err := strconv.Atoi(addressStr)
		if err != nil {
			log.Fatal(err)
		}

		value, err := strconv.ParseInt(instructionTokens[1], 10, 64)
		if err != nil {
			log.Fatal(err)
		}

		b.updateMemory(address, value)
	}
}

func (b *bootloader) sumMemory() int64 {
	sum := int64(0)
	for _, v := range b.mem {
		sum += v
	}

	return sum
}

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

func main() {
	instructions := getInput(os.Args[1])
	boot := new(bootloader)
	boot.mem = make(map[int]int64)
	for _, n := range *instructions {
		boot.processInstruction(n)
	}

	log.Printf("%+v", boot.sumMemory())
}
