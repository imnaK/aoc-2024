package main

import (
	"fmt"
	"log"
	"os"
	"sort"
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
	right := make([]int, arrSize)

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
		right[idx], err = strconv.Atoi(split[1])
		if err != nil {
			log.Fatal(err)
		}
	}

	// sort
	sort.Ints(left)
	sort.Ints(right)

	// get and sum the diffs
	var sums int
	for i := 0; i < arrSize; i++ {
		sums += getDiff(left[i], right[i])
	}

	fmt.Println(sums)
}

func getDiff(num1 int, num2 int) int {
	diff := num1 - num2
	if diff < 0 {
		diff = -diff
	}
	return diff
}
