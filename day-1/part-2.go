package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	raw, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	// craete vars and type out
	content := strings.Split(string(raw), "\n")
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

		left[idx], err = strconv.Atoi(split[0])
		if err != nil {
			log.Fatal(err)
		}
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

	fmt.Println(sum)
}
