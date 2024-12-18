package main

import (
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
	left := make([]int, len(content))
	right := make([]int, len(content))

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

		// sort
	}
}
