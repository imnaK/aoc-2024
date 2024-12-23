package day3

import (
	"fmt"
	"regexp"
)

func Day3Part2(data string) {
	impossibleMuls := regexp.MustCompile(`don't\(\)(.|\r|\n)*?(do\(\)|$)`)
	validMuls := regexp.MustCompile(`mul\(\d+,\d+\)`)
	sum := 0

	possibleMatches := impossibleMuls.ReplaceAllString(data, "")
	matches := validMuls.FindAllString(possibleMatches, -1)
	for _, match := range matches {
		parsedStmt, err := parseStatement(match)
		if err != nil {
			fmt.Println(fmt.Errorf("Failed to parse statement: %w", err))
		}

		sum += parsedStmt
	}

	fmt.Println(sum)
}
