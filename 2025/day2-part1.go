package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	inPath := "./day2-input.txt"
	inFile, err := os.ReadFile(inPath)
	if err != nil {
		fmt.Printf("Failed to open file %s: %s", inPath, err)
	}
	reString := regexp.MustCompile("\\d*-\\d*")
	idRangeList := reString.FindAllString(string(inFile), -1)
	// fmt.Printf("%v\n", idRangeList)

	invalidIdSum := 0

	for _, idRange := range idRangeList {
		idList := strings.Split(idRange, "-")
		// fmt.Printf("%v %v\n", idList, len(idList))
		if idList[0] == "" {
			break
		}
		startId, err := strconv.Atoi(idList[0])
		if err != nil {
			fmt.Errorf("Failed to parse int from string %v\n", idList[0])
		}
		endId, err := strconv.Atoi(idList[1])
		if err != nil {
			fmt.Errorf("Failed to parse int from string %v\n", idList[0])
		}
		// fmt.Printf("Start: %v, End: %v\n", startId, endId)
		for id := startId; id <= endId; id += 1 {
			idString := fmt.Sprintf("%v", id)
			midpoint := len(idString) / 2
			frontHalf := idString[:midpoint]
			backHalf := idString[midpoint:]
			if frontHalf == backHalf {
				// fmt.Printf("%v %s: %s %s (%v)\n", id, idString, frontHalf, backHalf, frontHalf == backHalf)
				invalidIdSum += id
			}
		}
	}

	fmt.Printf("Sum of invalid IDs: %v\n", invalidIdSum)
}
