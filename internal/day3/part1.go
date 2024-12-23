package day3

import (
	"fmt"
	"regexp"
)

func Day3Part1(data string) {
	validMuls := regexp.MustCompile(`mul\(\d+,\d+\)`)
	sum := 0

	matches := validMuls.FindAllString(data, -1)
	for _, match := range matches {
		parsedStmt, err := parseStatement(match)
		if err != nil {
			fmt.Println(fmt.Errorf("Failed to parse statement: %w", err))
		}

		sum += parsedStmt
	}

	fmt.Println(sum)
}
