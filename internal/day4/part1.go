package day4

import (
	"aoc-2024/pkg/datastructures"
	"aoc-2024/pkg/utils"
	"strings"
)

type Direction int

const (
	horizontal Direction = iota
	vertical
	diagonalTopLeft
	diagonalTopRight
)

const searchWord string = "XMAS"
const searchWordReversed string = "SAMX"

func Day4Part1(data string) any {
	//startOwn := time.Now()
	//for i := 0; i < 100; i++ {
	//	solve(data)
	//}
	//durationOwn := time.Since(startOwn)

	//startOther := time.Now()
	//for i := 0; i < 100; i++ {
	//	countWordOccurrences(data, searchWord)
	//}
	//durationOther := time.Since(startOther)

	//resOwn := solve(data)
	//return fmt.Sprintf("Own: %v | Other: %v\nResOwn: %v", durationOwn, durationOther, resOwn)

	return solve(data)
}

func solve(data string) any {
	puzzleWidth := strings.Index(data, "\n")
	puzzleHeight := strings.Count(data, "\n")

	dataAsBytes := []byte(strings.ReplaceAll(data, "\n", ""))
	founds := 0

	founds += countForwardAndBackward(getPuzzleRotated(dataAsBytes, horizontal, puzzleWidth, puzzleHeight, searchWord))
	founds += countForwardAndBackward(getPuzzleRotated(dataAsBytes, vertical, puzzleWidth, puzzleHeight, searchWord))
	founds += countForwardAndBackward(getPuzzleRotated(dataAsBytes, diagonalTopLeft, puzzleWidth, puzzleHeight, searchWord))
	founds += countForwardAndBackward(getPuzzleRotated(dataAsBytes, diagonalTopRight, puzzleWidth, puzzleHeight, searchWord))

	return founds
}

func countForwardAndBackward(puzzle []string) int {
	joinedPuzzle := strings.Join(puzzle, "\n")
	return strings.Count(joinedPuzzle, searchWord) + strings.Count(joinedPuzzle, searchWordReversed)
}

func getPuzzleRotated(data []byte, direction Direction, puzzleWidth int, puzzleHeight int, searchWord string) []string {
	var rotatedData []string

	switch direction {
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

func countWordOccurrences(data string, word string) int {
	board := utils.RemoveEmpty(strings.Split(data, "\n"))

	rows, cols := len(board), len(board[0])
	count := 0
	directions := [][]int{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1}, {0, 1},
		{1, -1}, {1, 0}, {1, 1},
	}

	var search func(i, j, dir int, index int) bool
	search = func(i, j, dir int, index int) bool {
		if index == len(word) {
			return true
		}
		if i < 0 || i >= rows || j < 0 || j >= cols || board[i][j] != word[index] {
			return false
		}
		return search(i+directions[dir][0], j+directions[dir][1], dir, index+1)
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if board[i][j] == word[0] {
				for dir := 0; dir < 8; dir++ {
					if search(i, j, dir, 0) {
						count++
					}
				}
			}
		}
	}

	return count
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
