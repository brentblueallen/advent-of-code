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
	accessibilityList := []string{}
	fmt.Printf("Room dimensions: %v x %v\n", roomWidth, roomHeight)

	accessibleRolls := 0
	for posY := 0; posY < roomHeight; posY += 1 {
		accessibilityLine := ""
		for posX := 0; posX < roomWidth; posX += 1 {
			if string(roomList[posY][posX]) == "@" {
				fmt.Printf("Roll @ (%v,%v)!\n", posX, posY)
				adjacentRolls := 0
				for offsetX := -1; offsetX <= 1; offsetX += 1 {
					for offsetY := -1; offsetY <= 1; offsetY += 1 {
						if (posX+offsetX >= 0) &&
							(posY+offsetY >= 0) &&
							(posX+offsetX < roomWidth) &&
							(posY+offsetY < roomHeight) {
							if string(roomList[posY+offsetY][posX+offsetX]) == "@" {
								adjacentRolls += 1
							}
						}
					}
				}
				if adjacentRolls <= 4 {
					accessibleRolls += 1
					accessibilityLine += "x"
				} else {
					accessibilityLine += "@"
				}
			} else {
				accessibilityLine += "."
			}
		}
		fmt.Println(accessibilityLine)
		accessibilityList = append(accessibilityList, accessibilityLine)
	}
	for posY := 0; posY < roomHeight; posY += 1 {
		fmt.Println(accessibilityList[posY])
	}
	fmt.Printf("Accessible rolls: %v\n", accessibleRolls)

}
