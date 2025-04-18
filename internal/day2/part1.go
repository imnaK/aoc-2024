package day2

import (
	"aoc-2024/pkg/utils"
	"log"
	"strconv"
	"strings"
)

func Day2Part1(data string) any {
	content := strings.Split(data, "\n")
	passed := 0

	for _, line := range content {
		if line == "" {
			continue
		}

		// get levels as array
		levels := strings.Split(line, " ")

		// check for ascending
		firstNum, err := strconv.Atoi(levels[0])
		if err != nil {
			log.Fatal(err)
		}
		secondNum, err := strconv.Atoi(levels[1])
		if err != nil {
			log.Fatal(err)
		}
		ascending := firstNum < secondNum
		passing := true

		// check all levels
		for i := 1; i < len(levels); i++ {
			lastNum, err := strconv.Atoi(levels[i-1])
			if err != nil {
				log.Fatal(err)
			}
			currNum, err := strconv.Atoi(levels[i])
			if err != nil {
				log.Fatal(err)
			}

			// check if still ascending/decending
			if (lastNum < currNum) != ascending {
				passing = false
				break
			}

			// check if number range is too big
			diff := utils.GetDiff(lastNum, currNum)
			if diff < 1 || diff > 3 {
				passing = false
				break
			}
		}

		if passing {
			passed++
		}
	}

	return passed
}
