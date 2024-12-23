package day2

import (
	"aoc-2024/pkg/utils"
	"errors"
	"log"
	"strconv"
	"strings"
)

func Day2Part2(data string) any {
	content := strings.Split(data, "\n")
	safeReports := 0

	// loop to check over every report
	for _, reportString := range content {
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

			isSafe := isReportSafe(modLevels, isAscending)

			if isSafe {
				safeReports++
				break
			}
		}
	}

	return safeReports
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

// returns if safe and if not, on which level it failed
func isReportSafe(levels []int, isAscending bool) bool {
	for i := 0; i < len(levels); i++ {
		safe := isLevelSafe(levels, i, isAscending)

		if !safe {
			return false
		}
	}

	return true
}

func isLevelSafe(levels []int, level int, isAscending bool) bool {
	if level == 0 {
		return true
	}

	prevLevel := levels[level-1]
	currLevel := levels[level]
	levelDiff := utils.GetDiff(prevLevel, currLevel)
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
