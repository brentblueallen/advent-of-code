package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	inPath := "./day2-example-input.txt"
	if len(os.Args) > 1 {
		inPath = os.Args[1]
	}
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
		for id := startId; id <= endId; id += 1 {
			idString := fmt.Sprintf("%v", id)
			idStringLen := len(idString)
			for chunkCount := 2; chunkCount <= idStringLen; chunkCount += 1 {
				// If evenly divisible into n chunks
				if idStringLen%chunkCount == 0 {
					chunkLen := idStringLen / chunkCount
					chunks := []string{}
					// Chunk it up
					for segmentIdx := 0; segmentIdx < idStringLen; segmentIdx += chunkLen {
						chunks = append(chunks, idString[segmentIdx:segmentIdx+chunkLen])
					}
					// Check if the chunks are all matching
					if allChunksMatch(chunks) {
						invalidIdSum += id
						// fmt.Printf("%s chunked by %v: %v -> %v\n", idString, chunkLen, chunks, allChunksMatch(chunks))
						// Break on first match to avoid double-counting
						break
					}
				}
			}
		}
	}

	fmt.Printf("Sum of invalid IDs: %v\n", invalidIdSum)
}

func allChunksMatch(chunks []string) bool {
	allMatch := true
	for i := 1; i < len(chunks); i += 1 {
		// If any chunk doesn't match the first, all don't match
		if chunks[i] != chunks[0] {
			allMatch = false
		}
	}
	return allMatch
}
