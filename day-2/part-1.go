package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	raw, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	content := strings.Split(string(raw), "\n")

	for _, line := range content {
		if line == "" {
			continue
		}

		fmt.Println(line)
	}
}

func getDiff(num1 int, num2 int) int {
	diff := num1 - num2
	if diff < 0 {
		diff = -diff
	}
	return diff
}
