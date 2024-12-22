package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

const inFileName = "input.txt"

func main() {
	// read file content and split by line to []string
	fileContent, err := getFileContent()
	if err != nil {
		log.Fatal(err)
	}
	content := strings.Split(fileContent, "\n")
	safeReports := 0

	// loop to check over every report
	for reportIdx, reportString := range content {
		// skip if report has no content
		if reportString == "" {
			continue
		}

		// get levels as array
		levels, isAscending, err := getLevels(reportString)
		if err != nil {
			log.Fatal(err)
		}

		// if only one level, it's safe anyway
		if len(levels) <= 1 {
			safeReports++
			continue
		}

		for i := -1; i < len(levels); i++ {
			modLevels := make([]int, len(levels))
			copy(modLevels, levels)

			// create a new slice with a missing level to check
			if i >= 0 {
				modLevels = getLevelsWithIgnoredLevel(modLevels, i)
			}

			isSafe, newIgnoreLevel := isReportSafe(modLevels, isAscending)
			fmt.Printf("[%v] || %v, %v, | %v, %v | %v\n", reportIdx, isAscending, i, newIgnoreLevel, isSafe, modLevels)

			if isSafe {
				safeReports++
				break
			}
		}
	}

	fmt.Printf("safeReports:%d\n", safeReports)
}

// returns if safe and if not, on which level it failed
func isReportSafe(levels []int, isAscending bool) (bool, int) {
	for i := 0; i < len(levels); i++ {
		safe := isLevelSafe(levels, i, isAscending)

		if !safe {
			return false, i
		}
	}

	return true, -1
}

func isLevelSafe(levels []int, level int, isAscending bool) bool {
	if level == 0 {
		return true
	}

	prevLevel := levels[level-1]
	currLevel := levels[level]
	levelDiff := getDiff(prevLevel, currLevel)
	isCurrAscending := prevLevel < currLevel == isAscending

	// check for out of range
	if levelDiff < 1 || levelDiff > 3 {
		return false
	}

	// check if ascending/decending order is still the same
	if !isCurrAscending {
		return false
	}

	return true
}

func getLevels(levelsString string) ([]int, bool, error) {
	if levelsString == "" {
		return []int{}, false, errors.New("Input string must not be empty")
	}

	asc := 0
	levelsChars := strings.Split(levelsString, " ")
	levels := make([]int, len(levelsChars))

	for idx, levelChar := range levelsChars {
		levelAsInt, err := strconv.Atoi(levelChar)
		if err != nil {
			return []int{}, false, err
		}
		levels[idx] = levelAsInt

		if idx > 0 {
			if levels[idx-1] < levels[idx] {
				asc++
			} else {
				asc--
			}
		}
	}

	return levels, asc >= 0, nil
}

func getLevelsWithIgnoredLevel(levels []int, ignoreLevel int) []int {
	return append(levels[:ignoreLevel], levels[ignoreLevel+1:]...)
}

func getDiff(num1 int, num2 int) int {
	diff := num1 - num2
	if diff < 0 {
		diff = -diff
	}
	return diff
}

func getFileContent() (string, error) {
	// get the working dir
	wd, err := os.Getwd()
	if err != nil {
		return "", err
	}

	// get the parent dir
	pd := filepath.Dir(wd)

	// construct path to the input file
	inFilePath := filepath.Join(pd, inFileName)

	raw, err := os.ReadFile(inFilePath)
	if err != nil {
		return "", err
	}

	return string(raw), nil
}
