package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	inPath := "./example"
	if len(os.Args) > 1 {
		inPath = os.Args[1]
	}
	inFile, err := os.ReadFile(inPath)
	if err != nil {
		fmt.Printf("Failed to open file %s: %s", inPath, err)
	}

	// Split input file into two lists of fresh id ranges and ingredient ids
	superList := strings.Split(string(inFile), "\n\n")

	// Parse the fresh id range strings into type FreshRange
	freshRanges := []FreshRange{}
	for _, rangeStr := range strings.Split(superList[0], "\n") {
		newRange := FreshRange{}
		err := newRange.Parse(rangeStr)
		if err == nil {
			freshRanges = append(freshRanges, newRange)
		}
	}

	ingredientList := strings.Split(superList[1], "\n")

	freshCount := 0
	for _, ingredient := range ingredientList {
		if len(ingredient) > 0 {
			id, err := strconv.Atoi(string(ingredient))
			if err != nil {
				fmt.Println(err)
			}
			if isFresh(freshRanges, id) {
				freshCount += 1
			}
		}
	}

	fmt.Printf("Available fresh ingredients: %v\n", freshCount)
}

func isFresh(freshRanges []FreshRange, id int) bool {
	for _, fr := range freshRanges {
		if fr.Includes(id) {
			return true
		}
	}
	return false
}

type FreshRange struct {
	Min int
	Max int
}

func (fr *FreshRange) Parse(rangeStr string) (err error) {
	rangeList := strings.Split(rangeStr, "-")
	if len(rangeList) < 2 {
		return fmt.Errorf("Input range list insufficiently long")
	}
	fr.Min, err = strconv.Atoi(rangeList[0])
	fr.Max, err = strconv.Atoi(rangeList[1])
	return nil
}

func (fr *FreshRange) Includes(id int) (includes bool) {
	if (id >= fr.Min) && (id <= fr.Max) {
		return true
	} else {
		return false
	}
}
