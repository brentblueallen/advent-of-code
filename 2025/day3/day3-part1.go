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

	bankList := strings.Split(string(inFile), "\n")
	fmt.Printf("%v\n", bankList)

	sumJolts := 0
	for _, bank := range bankList {
		if len(bank) == 0 {
			break
		}
		fmt.Println(bank)
		maxJolts := 0
		for i := 0; i < len(bank)-1; i++ {
			for j := i + 1; j < len(bank); j++ {
				jolts, err := strconv.Atoi(string(bank[i]) + string(bank[j]))
				if err != nil {
					fmt.Printf("Failed to parse integer from string \"%s%s\"\n", bank[i], bank[j])
				}
				if jolts > maxJolts {
					fmt.Printf("maxJolts %v -> %v\n", maxJolts, jolts)
					maxJolts = jolts
				}
			}
		}
		sumJolts += maxJolts
	}

	fmt.Printf("Sum of bank jolt ratings: %v\n", sumJolts)

}
