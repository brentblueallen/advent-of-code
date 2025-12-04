package main

import (
	"fmt"
	"os"
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

	// Determine dimensions of room
	roomList := strings.Split(string(inFile), "\n")
	roomWidth := len(roomList[0])
	roomHeight := len(roomList)

	// Drop last line if empty (protect for newline at end of input)
	if roomList[roomHeight-1] == "" {
		roomHeight -= 1
		roomList = roomList[:roomHeight]
	}
	fmt.Printf("Room dimensions: %v x %v\n", roomWidth, roomHeight)

	accessibleRolls := 0
	accessedRolls := 0
	for {
		accessibleRolls, roomList = accessRolls(roomList)
		if accessibleRolls == 0 {
			break
		}
		accessedRolls += accessibleRolls
	}

	fmt.Printf("Accessed rolls: %v\n", accessedRolls)
}

func accessRolls(rollMap []string) (accessibleRolls int, accessedMap []string) {
	mapWidth := len(rollMap[0])
	mapHeight := len(rollMap)

	accessedMap = []string{}
	accessibleRolls = 0
	for posY := 0; posY < mapHeight; posY += 1 {
		accessedLine := ""
		for posX := 0; posX < mapWidth; posX += 1 {
			if string(rollMap[posY][posX]) == "@" {
				// fmt.Printf("Roll @ (%v,%v)!\n", posX, posY)
				adjacentRolls := 0
				for offsetX := -1; offsetX <= 1; offsetX += 1 {
					for offsetY := -1; offsetY <= 1; offsetY += 1 {
						if (posX+offsetX >= 0) &&
							(posY+offsetY >= 0) &&
							(posX+offsetX < mapWidth) &&
							(posY+offsetY < mapHeight) {
							if string(rollMap[posY+offsetY][posX+offsetX]) == "@" {
								adjacentRolls += 1
							}
						}
					}
				}
				if adjacentRolls <= 4 {
					accessibleRolls += 1
					accessedLine += "."
				} else {
					accessedLine += "@"
				}
			} else {
				accessedLine += "."
			}
		}
		accessedMap = append(accessedMap, accessedLine)
	}
	return
}
