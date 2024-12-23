package day3

import (
	"errors"
	"strconv"
	"strings"
)

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
