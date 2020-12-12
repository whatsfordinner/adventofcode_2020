package main

import (
	"bufio"
	"log"
	"math"
	"os"
	"strconv"
)

type boat struct {
	posX      int
	posY      int
	waypointX int
	waypointY int
}

type instruction struct {
	direction string
	distance  int
}

func (b *boat) move(i instruction) {
	if i.direction == "N" {
		b.waypointY += i.distance
	}

	if i.direction == "S" {
		b.waypointY -= i.distance
	}

	if i.direction == "E" {
		b.waypointX += i.distance
	}

	if i.direction == "W" {
		b.waypointX -= i.distance
	}

	if i.direction == "F" {
		b.posX += b.waypointX * i.distance
		b.posY += b.waypointY * i.distance
	}

	if i.direction == "R" {
		bearing := i.distance % 360
		b.waypointX, b.waypointY = rotate(b.waypointX, b.waypointY, bearing)
	}

	if i.direction == "L" {
		bearing := (360 - i.distance) % 360
		b.waypointX, b.waypointY = rotate(b.waypointX, b.waypointY, bearing)
	}
}

func (b *boat) getManhattanDistance() int {
	return int(math.Abs(float64(b.posX)) + math.Abs(float64(b.posY)))
}

func rotate(x int, y int, theta int) (int, int) {
	if theta == 90 {
		return y, -x
	}

	if theta == 180 {
		return -x, -y
	}

	if theta == 270 {
		return -y, x
	}

	return x, y
}

func getInput(filename string) *[]instruction {
	result := new([]instruction)

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if scanner.Text() != "" {
			direction := string(scanner.Text()[0])
			distance, err := strconv.Atoi(scanner.Text()[1:])
			if err != nil {
				log.Fatal(err)
			}
			*result = append(*result, instruction{direction, distance})
		}
	}

	return result
}

func main() {
	input := getInput(os.Args[1])
	b := new(boat)
	b.posX = 0
	b.posY = 0
	b.waypointX = 10
	b.waypointY = 1

	for _, d := range *input {
		b.move(d)
		log.Printf("%+v -> %+v", d, *b)
	}

	log.Printf("Current location: (%d, %d)", b.posX, b.posY)
	log.Printf("Manhattan distance travelled: %d", b.getManhattanDistance())
}
