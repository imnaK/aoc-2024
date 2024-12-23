package day4

import (
	"aoc-2024/pkg/utils"
	"fmt"
	"strings"
)

type Direction int

const (
	horizontal Direction = iota
	horizontalReversed
	vertical
	verticalReversed
	diagonalTopLeft
	diagonalTopRight
	diagonalTopLeftReversed
	diagonalTopRightReversed
)

const xmas string = "XMAS"

func Day4Part1(data string) any {
	puzzleWidth := strings.Index(data, "\n")
	puzzleHeight := strings.Count(data, "\n")

	fmt.Printf("Puzzle width: %d\nPuzzle height: %d\n", puzzleWidth, puzzleHeight)

	rotatedPuzzle, _ := getPuzzleRotated([]byte(strings.ReplaceAll(data, "\n", "")), diagonalTopRight, puzzleWidth, puzzleHeight)

	return strings.Join(rotatedPuzzle, "\n")
}

func getPuzzleRotated(data []byte, direction Direction, puzzleWidth int, puzzleHeight int) ([]string, string) {
	searchString := []byte(xmas)
	var rotatedData []string

	switch direction {
	case horizontalReversed, verticalReversed, diagonalTopLeftReversed, diagonalTopRightReversed:
		utils.ReverseArray(searchString)
	case horizontal:
		rotatedData = make([]string, puzzleHeight)
		for y := range puzzleHeight {
			offset := y * puzzleWidth
			rotatedData[y] = string(data[offset : offset+puzzleWidth])
		}
		break
	case vertical:
		rotatedData = make([]string, puzzleHeight)
		for y := range puzzleHeight {
			newLine := make([]byte, puzzleHeight)
			for x := range puzzleWidth {
				z := x*puzzleWidth + y
				newLine[x] = data[z]
			}
			rotatedData[y] = string(newLine)
		}
		break
	case diagonalTopLeft:
		rotatedData = make([]string, puzzleHeight+puzzleWidth-1)
		for y := range puzzleHeight {
			newLineLen := y + 1
			if newLineLen > puzzleWidth-1 {
				newLineLen = puzzleWidth
			}

			newLine := make([]byte, newLineLen)

			for z := range newLineLen {
				someY := puzzleHeight - y - 1 + z
				dataIdx := someY*puzzleWidth + z
				newLine[z] = data[dataIdx]
			}

			rotatedData[y] = string(newLine)
		}

		// minus 1 for the mid diagonal row which already was added by above
		for x := range puzzleWidth - 1 {
			newLineLen := x + 1
			if newLineLen > puzzleHeight-1 {
				newLineLen = puzzleHeight
			}

			newLine := make([]byte, newLineLen)

			for z := range newLineLen {
				someX := puzzleWidth - x - 1 + z
				dataIdx := z*puzzleWidth + someX
				newLine[z] = data[dataIdx]
			}

			rotatedData[x+puzzleHeight] = string(newLine)
		}
		break
	case diagonalTopRight:
		rotatedData = make([]string, puzzleHeight+puzzleWidth-1)
		for y := range puzzleHeight {
			newLineLen := y + 1
			if newLineLen > puzzleWidth-1 {
				newLineLen = puzzleWidth
			}

			newLine := make([]byte, newLineLen)

			for z := range newLineLen {
				someY := puzzleHeight - y - 1 + z
				dataIdx := someY*puzzleWidth + (puzzleWidth - z - 1)
				newLine[z] = data[dataIdx]
			}

			rotatedData[y] = string(newLine)
		}

		// minus 1 for the mid diagonal row which already was added by above
		for x := range puzzleWidth - 1 {
			newLineLen := x + 1
			if newLineLen > puzzleHeight-1 {
				newLineLen = puzzleHeight
			}

			newLine := make([]byte, newLineLen)

			for z := range newLineLen {
				someX := x - z
				dataIdx := z*puzzleWidth + someX
				newLine[z] = data[dataIdx]
			}

			rotatedData[x+puzzleHeight] = string(newLine)
		}
		break
	}

	return rotatedData, string(searchString)
}
