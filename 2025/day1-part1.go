package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Howdy")
	inPath := "./day1-input.txt"
	inFile, err := os.ReadFile(inPath)
	if err != nil {
		fmt.Printf("Failed to open file %s: %s", inPath, err)
	}
	commandList := strings.Split(string(inFile), "\n")
	// fmt.Printf("%v", commandList)

	dialPosition := 50
	zeroCount := 0
	fmt.Printf("- The dial starts by pointing at %v.\n", dialPosition)
	for _, command := range commandList {
		var direction int
		if len(command) == 0 {
			break
		}

		switch command[0] {
		case 'R':
			direction = 1
		case 'L':
			direction = -1
		}

		steps, err := strconv.Atoi(command[1:])

		if err != nil {
			fmt.Printf("Failed to parse int portion %s!", command[1:])
		}
		steps *= direction

		dialPosition = (dialPosition + steps) % 100
		for dialPosition < 0 {
			dialPosition += 100
		}
		if dialPosition == 0 {
			zeroCount += 1
		}

		fmt.Printf("- The dial is rotated %s to point at %v.\n", command, dialPosition)
	}
	fmt.Printf("zeroCount: %v\n", zeroCount)
}
