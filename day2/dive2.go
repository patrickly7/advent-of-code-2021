package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func getDepthPositionResult(horizontalPositionChanges []int, depthChanges []int) int {
	var currDepth = 0
	var currHorizontalPosition = 0

	for _, depthChange := range depthChanges {
		currDepth += depthChange
	}

	for _, horizontalPositionChange := range horizontalPositionChanges {
		currHorizontalPosition += horizontalPositionChange
	}

	return currDepth * currHorizontalPosition
}

func main() {
	var aim = 0
	var depthMovements []int
	var horizontalMovements []int

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}

	defer file.Close()

	var scanner = bufio.NewScanner(file)

	for scanner.Scan() {
		var splitString = strings.Fields(scanner.Text())
		movement, _ := strconv.Atoi(splitString[1])

		if splitString[0] == "forward" {
			horizontalMovements = append(horizontalMovements, movement)
			depthMovements = append(depthMovements, aim*movement)
		} else if splitString[0] == "up" {
			aim -= movement
		} else if splitString[0] == "down" {
			aim += movement
		}
	}

	fmt.Printf("The multiplication result of final depth and height is: %d\n", getDepthPositionResult(depthMovements, horizontalMovements))
}
