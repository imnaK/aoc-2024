package day4

import (
	"aoc-2024/pkg/datastructures"
	"aoc-2024/pkg/utils"
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
	return solve(data)
}

func solve(data string) any {
	puzzleWidth := strings.Index(data, "\n")
	puzzleHeight := strings.Count(data, "\n")

	dataAsBytes := []byte(strings.ReplaceAll(data, "\n", ""))
	founds := 0

	for _, direction := range []Direction{horizontal, horizontalReversed, vertical, verticalReversed, diagonalTopLeft, diagonalTopLeftReversed, diagonalTopRight, diagonalTopRightReversed} {
		rotatedPuzzle, searchString := getPuzzleRotated(dataAsBytes, direction, puzzleWidth, puzzleHeight)

		for _, line := range rotatedPuzzle {
			founds += strings.Count(line, searchString)
		}
	}

	return founds
}

func getPuzzleRotated(data []byte, direction Direction, puzzleWidth int, puzzleHeight int) ([]string, string) {
	searchString := []byte(xmas)
	var rotatedData []string

	switch direction {
	case horizontalReversed, verticalReversed, diagonalTopLeftReversed, diagonalTopRightReversed:
		utils.ReverseArray(searchString)
	}

	switch direction {
	case horizontal, horizontalReversed:
		rotatedData = make([]string, puzzleHeight)
		for y := range puzzleHeight {
			offset := y * puzzleWidth
			rotatedData[y] = string(data[offset : offset+puzzleWidth])
		}
		break
	case vertical, verticalReversed:
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
	case diagonalTopLeft, diagonalTopLeftReversed:
		offset := len(searchString) - 1
		rotatedData = make([]string, puzzleHeight+puzzleWidth-1-offset*2)
		for y := offset; y < puzzleHeight; y++ {
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
				someX := puzzleWidth - x - 1 + z
				dataIdx := z*puzzleWidth + someX
				newLine[z] = data[dataIdx]
			}

			rotatedData[x+puzzleHeight-offset*2] = string(newLine)
		}
		break
	case diagonalTopRight, diagonalTopRightReversed:
		offset := len(searchString) - 1
		rotatedData = make([]string, puzzleHeight+puzzleWidth-1-offset*2)
		for y := offset; y < puzzleHeight; y++ {
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
				someX := x - z
				dataIdx := z*puzzleWidth + someX
				newLine[z] = data[dataIdx]
			}

			rotatedData[x+puzzleHeight-offset*2] = string(newLine)
		}
		break
	}

	return rotatedData, string(searchString)
}

func getFoundWords(data string, words []string) any {
	board := utils.RemoveEmpty(strings.Split(data, "\n"))

	trie := datastructures.NewTrie()
	for _, word := range words {
		trie.Insert(word)
	}

	result := make(map[string]bool)
	rows, cols := len(board), len(board[0])

	var dfs func(i, j int, node *datastructures.TrieNode, word string)
	dfs = func(i, j int, node *datastructures.TrieNode, word string) {
		if i < 0 || i >= rows || j < 0 || j >= cols {
			return
		}

		ch := rune(board[i][j])
		if next, exists := node.Children[ch]; exists {
			word += string(ch)
			if next.IsEnd {
				result[word] = true
			}

			directions := [][]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}
			for _, dir := range directions {
				dfs(i+dir[0], j+dir[1], next, word)
			}
		}
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			dfs(i, j, trie.Root, "")
		}
	}

	foundWords := make([]string, 0, len(result))
	for word := range result {
		foundWords = append(foundWords, word)
	}
	return foundWords
}
