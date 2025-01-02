package day5

import (
	"aoc-2024/pkg/datastructures"
	"log"
	"strconv"
	"strings"
)

func Day5Part1(data string) any {
	myDag := datastructures.NewDag[int]()

	firstPart := strings.Split(data, "\n\n")[0]
	rulesStr := strings.Split(firstPart, "\n")
	rules := make([][]int, len(rulesStr))

	for idx, rule := range rulesStr {
		valsStr := strings.Split(rule, "|")
		vals := make([]int, 2)

		prev, err := strconv.Atoi(valsStr[0])
		if err != nil {
			log.Fatal(err)
		}
		next, err := strconv.Atoi(valsStr[1])
		if err != nil {
			log.Fatal(err)
		}

		vals[0] = prev
		vals[1] = next

		rules[idx] = vals
	}

	for _, rule := range rules {
		myDag.Insert(rule[0], rule[1])
	}

	return myDag.ToString()
}
