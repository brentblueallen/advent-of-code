package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	cNumJoltDigits int = 12
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

	bankList := strings.Split(string(inFile), "\n")
	// fmt.Printf("%v\n", bankList)

	sumJolts := 0
	for _, bank := range bankList {
		if len(bank) == 0 {
			break
		}
		// fmt.Println(bank)
		lastFoundIndex := -1
		jolts := 0
		for digitsLeft := cNumJoltDigits; digitsLeft > 0; {
			var foundDigit int
			foundDigit, lastFoundIndex = maxInRange(bank, lastFoundIndex+1, digitsLeft)
			digitsLeft -= 1
			// fmt.Printf("digit %v @ index %v, %v digits remaining\n", foundDigit, lastFoundIndex, digitsLeft)
			jolts = (jolts * 10) + foundDigit
		}
		// fmt.Println(jolts)
		sumJolts += jolts
	}

	fmt.Printf("Sum of bank jolt ratings: %v\n", sumJolts)

}

// Search through a string for the greatest digit, starting at a specified index and
// leaving a specified number of digits unexplored
func maxInRange(bank string, leftPad int, rightPad int) (maxDigit int, maxIndex int) {
	maxDigit = 0
	maxIndex = 0

	// Define chunk to search by last found digit index and digits to save for next search
	endIndex := len(bank) - rightPad
	searchChunk := bank[leftPad : endIndex+1]
	// fmt.Printf("%s\n", searchChunk)

	// Search through chunk for highest value and return its index within input string
	for idx := 0; idx < len(searchChunk); idx += 1 {
		digit, err := strconv.Atoi(string(searchChunk[idx]))
		if err != nil {
			fmt.Printf("Failed to parse integer from character %s\n", string(searchChunk[idx]))
		}
		if digit > maxDigit {
			maxDigit = digit
			maxIndex = idx + leftPad
			// fmt.Printf("\tmaxDigit=%v @ %v\n", maxDigit, maxIndex)
		}
	}
	return
}
