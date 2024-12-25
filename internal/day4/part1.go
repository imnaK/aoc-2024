package day4

import (
	"strings"
)

type Direction int

const (
	vertical Direction = iota
	diagonalTopLeft
	diagonalTopRight
)

const searchWord string = "XMAS"
const searchWordReversed string = "SAMX"

func Day4Part1(data string) any {
	puzzleWidth := strings.Index(data, "\n")
	puzzleHeight := strings.Count(data, "\n")

	dataAsBytes := []byte(strings.ReplaceAll(data, "\n", ""))
	founds := 0

	founds += countForwardAndBackward(data)
	for _, direction := range []Direction{vertical, diagonalTopLeft, diagonalTopRight} {
		founds += countForwardAndBackward(strings.Join(getPuzzleRotated(dataAsBytes, direction, puzzleWidth, puzzleHeight), "\n"))
	}

	return founds
}

func countForwardAndBackward(puzzle string) int {
	return strings.Count(puzzle, searchWord) + strings.Count(puzzle, searchWordReversed)
}

func getPuzzleRotated(data []byte, direction Direction, puzzleWidth int, puzzleHeight int) []string {
	var rotatedData []string

	switch direction {
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
	case diagonalTopLeft, diagonalTopRight:
		offset := len(searchWord) - 1
		rotatedData = make([]string, puzzleHeight+puzzleWidth-1-offset*2)
		for y := offset; y < puzzleHeight; y++ {
			newLineLen := y + 1
			if newLineLen > puzzleWidth-1 {
				newLineLen = puzzleWidth
			}

			newLine := make([]byte, newLineLen)

			for z := range newLineLen {
				someY := puzzleHeight - y - 1 + z
				var dataIdx int
				if direction == diagonalTopLeft {
					dataIdx = someY*puzzleWidth + z
				} else {
					dataIdx = someY*puzzleWidth + (puzzleWidth - z - 1)
				}
				newLine[z] = data[dataIdx]
			}

			rotatedData[y-offset] = string(newLine)
		}

		// minus 1 for the mid diagonal row which already was added by above
		for x := offset; x < puzzleWidth-1; x++ {
			newLineLen := x + 1
			if newLineLen > puzzleHeight-1 {
				newLineLen = puzzleHeight
			}

			newLine := make([]byte, newLineLen)

			for z := range newLineLen {
				var someX int
				if direction == diagonalTopLeft {
					someX = puzzleWidth - x - 1 + z
				} else {
					someX = x - z
				}
				dataIdx := z*puzzleWidth + someX
				newLine[z] = data[dataIdx]
			}

			rotatedData[x+puzzleHeight-offset*2] = string(newLine)
		}
		break
	}

	return rotatedData
}
