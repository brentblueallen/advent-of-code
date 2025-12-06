package main

import (
	"fmt"
	"os"
	"regexp"
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

	// Split input file into lists of arguments and operations (in the last non-empty line)
	superList := strings.Split(string(inFile), "\n")

	// Parse input file into list of argument lists and list of operators
	re := regexp.MustCompile("\\d+")
	argList := [][]int{}
	opList := []string{}
	for _, line := range superList {
		if line == "" {
			// If empty line (end-of-file)
			// fmt.Println("Empty line!")
			break
		} else if re.MatchString(line) {
			// If argument line
			// fmt.Println("Arg line!")
			argumentStrings := re.FindAllString(line, -1)
			argumentInts := []int{}
			for _, str := range argumentStrings {
				argInt, err := strconv.Atoi(str)
				if err != nil {
					fmt.Printf("Failed to parse \"%s\" to int", str)
				}
				argumentInts = append(argumentInts, argInt)
			}
			argList = append(argList, argumentInts)
		} else {
			// Else operator line
			// fmt.Println("Op line!")
			opList = regexp.MustCompile("[+\\*]").FindAllString(line, -1)
		}
	}
	// fmt.Printf("args: %v\nops: %v\n", argList, opList)

	// Perform calculations
	answerList := []int{}
	for col, op := range opList {
		prod := 1
		sum := 0
		for row := range len(argList) {
			switch op {
			case "+":
				sum += argList[row][col]
			case "*":
				prod *= argList[row][col]
			}
			/*if row == len(argList)-1 {
				fmt.Printf("%v = ", argList[row][col])
			} else {
				fmt.Printf("%v %s ", argList[row][col], opList[col])
			}*/
		}
		switch op {
		case "+":
			answerList = append(answerList, sum)
			// fmt.Println(sum)
		case "*":
			answerList = append(answerList, prod)
			// fmt.Println(prod)
		}
	}
	// fmt.Println(answerList)

	// Calculate sum of answers
	answerSum := 0
	for _, answer := range answerList {
		answerSum += answer
	}
	fmt.Println(answerSum)
}
