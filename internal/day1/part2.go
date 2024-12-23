package day1

import (
	"log"
	"strconv"
	"strings"
)

func Day1Part2(data string) any {
	content := strings.Split(data, "\n")
	arrSize := len(content)
	left := make([]int, arrSize)
	counts := make(map[int]int)
	var sum int

	for idx, line := range content {
		if line == "" {
			continue
		}

		// assign numbers to arrays
		split := strings.Split(line, "   ")

		leftVal, err := strconv.Atoi(split[0])
		if err != nil {
			log.Fatal(err)
		}
		left[idx] = leftVal

		rightVal, err := strconv.Atoi(split[1])
		if err != nil {
			log.Fatal(err)
		}

		_, exists := counts[rightVal]
		if exists {
			counts[rightVal] += 1
		} else {
			counts[rightVal] = 1
		}
	}

	for idx := 0; idx < arrSize; idx++ {
		leftVal := left[idx]
		rightCount, exists := counts[leftVal]
		if !exists {
			continue
		}

		sum += leftVal * rightCount
	}

	return sum
}
