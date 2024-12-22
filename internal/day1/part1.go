package day1

import (
	"aoc-2024/pkg/utils"
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"
)

func Day1Part1(data string) {
	content := strings.Split(data, "\n")
	arrSize := len(content)
	left := make([]int, arrSize)
	right := make([]int, arrSize)

	for idx, line := range content {
		if line == "" {
			continue
		}

		// assign numbers to arrays
		split := strings.Split(line, "   ")

		leftVar, err := strconv.Atoi(split[0])
		if err != nil {
			log.Fatal(err)
		}
		left[idx] = leftVar

		rightVar, err := strconv.Atoi(split[1])
		if err != nil {
			log.Fatal(err)
		}
		right[idx] = rightVar
	}

	// sort
	sort.Ints(left)
	sort.Ints(right)

	// get and sum the diffs
	var sums int
	for i := 0; i < arrSize; i++ {
		sums += utils.GetDiff(left[i], right[i])
	}

	fmt.Println(sums)
}
