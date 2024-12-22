package day3

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
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

func parseStatement(stmt string) (int, error) {
	stmtSplit := strings.Split(stmt, ",")
	firstNum, err := strconv.Atoi(stmtSplit[0][4:])
	if err != nil {
		return 0, errors.New("First parameter is not a whole number")
	}
	secondNum, err := strconv.Atoi(stmtSplit[1][:len(stmtSplit[1])-1])
	if err != nil {
		return 0, errors.New("Second parameter is not a whole number")
	}

	return firstNum * secondNum, nil
}
