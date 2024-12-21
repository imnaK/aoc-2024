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

const inFileName = "input2.txt"

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
		levels, err := getLevels(reportString)
		if err != nil {
			log.Fatal(err)
		}

		// if only one level, it's safe anyway
		if len(levels) <= 1 {
			safeReports++
			continue
		}

		// checking if report is safe, also when a specific index is ignored
		problemDamper := 0
		safeDampers := 1
		ignoreLevel := -1

		for problemDamper <= safeDampers {
			modLevels := make([]int, len(levels))
			copy(modLevels, levels)

			// create a new slice with a missing level to check
			if ignoreLevel >= 0 {
				modLevels = getLevelsWithIgnoredLevel(modLevels, ignoreLevel)
			}

			isSafe, newIgnoreLevel := isReportSafe(modLevels)
			fmt.Printf("[%v] || %v, %v, | %v, %v | %v\n", reportIdx, problemDamper, ignoreLevel, newIgnoreLevel, isSafe, modLevels)

			if ignoreLevel == 0 && !isSafe {
				problemDamper--
			}

			if isSafe {
				safeReports++
				break
			}

			if ignoreLevel == -1 && newIgnoreLevel == 2 {
				ignoreLevel = 0
			} else {
				ignoreLevel = newIgnoreLevel
			}

			problemDamper++
		}
	}

	fmt.Printf("safeReports:%d\n", safeReports)
}

// returns if safe and if not, on which level it failed
func isReportSafe(levels []int) (bool, int) {
	lastAscending := false

	for i := 0; i < len(levels); i++ {
		safe, newAscending := isLevelSafe(levels, i, lastAscending)
		lastAscending = newAscending

		if !safe {
			return false, i
		}
	}

	return true, -1
}

func isLevelSafe(levels []int, level int, ascending bool) (bool, bool) {
	if level == 0 {
		return true, ascending
	}

	prevLevel := levels[level-1]
	currLevel := levels[level]
	levelDiff := getDiff(prevLevel, currLevel)
	stillAscending := prevLevel < currLevel

	// check for out of range
	if levelDiff < 1 || levelDiff > 3 {
		return false, stillAscending
	}

	// check if ascending/decending order is still the same
	if level > 1 && (stillAscending != ascending) {
		return false, stillAscending
	}

	return true, stillAscending
}

func getLevels(levelsString string) ([]int, error) {
	if levelsString == "" {
		return []int{}, errors.New("Input string must not be empty")
	}

	levelsChars := strings.Split(levelsString, " ")
	levels := make([]int, len(levelsChars))

	for idx, levelChar := range levelsChars {
		levelAsInt, err := strconv.Atoi(levelChar)
		if err != nil {
			return []int{}, err
		}
		levels[idx] = levelAsInt
	}

	return levels, nil
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
