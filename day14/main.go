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
	mem  map[int64]int64
}

func (b *bootloader) applyMask(input int64) []int64 {
	inputBinary := fmt.Sprintf("%036s", strconv.FormatInt(int64(input), 2))
	outputBinaries := []*string{}
	outputBinaries = append(outputBinaries, new(string))
	output := []int64{}

	for i, n := range b.mask {
		if n == 'X' {
			stop := len(outputBinaries)
			for i := 0; i < stop; i++ {
				outputBinaries = append(outputBinaries, new(string))
				*(outputBinaries[stop+i]) = *(outputBinaries[i])
			}
		}

		for j, o := range outputBinaries {
			if n == '0' {
				*o += string(inputBinary[i])
			}

			if n == '1' {
				*o += "1"
			}

			if n == 'X' {
				if float64(j+1)/float64(len(outputBinaries)) <= 0.5 {
					*o += "0"
				} else {
					*o += "1"
				}
			}
		}
	}

	for _, n := range outputBinaries {
		out, err := strconv.ParseInt(*n, 2, 64)
		if err != nil {
			log.Fatal(err)
		}

		output = append(output, out)
	}

	return output
}

func (b *bootloader) updateMemory(address int64, input int64) {
	addresses := b.applyMask(address)
	for _, a := range addresses {
		b.mem[a] = input
	}
}

func (b *bootloader) processInstruction(instruction string) {
	instructionTokens := strings.Split(instruction, " = ")
	if instructionTokens[0] == "mask" {
		b.mask = instructionTokens[1]
	} else {
		re := regexp.MustCompile(`\d+`)
		addressStr := string(re.Find([]byte(instructionTokens[0])))
		address, err := strconv.ParseInt(addressStr, 10, 64)
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
	boot.mem = make(map[int64]int64)
	for _, n := range *instructions {
		boot.processInstruction(n)
	}

	log.Printf("%+v", boot.sumMemory())
}
