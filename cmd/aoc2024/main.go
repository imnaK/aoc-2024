package main

import (
	"aoc-2024/internal/day1"
	"aoc-2024/internal/day2"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

type dayPartRunFunction func(string)

func main() {
	args := os.Args
	if len(args) <= 2 {
		log.Fatal("Please provide a day and a part to process with numbers")
	}

	dayPartMap := map[string]dayPartRunFunction{
		"11": day1.Day1Part1,
		"12": day1.Day1Part2,
		"21": day2.Day2Part1,
		"22": day2.Day2Part2,
	}

	day := args[1]
	part := args[2]

	dayPartMap[fmt.Sprintf("%s%s", day, part)](getDataFromFile(day))
}

func getDataFromFile(day string) string {
	// get project root
	_, filename, _, _ := runtime.Caller(0)
	porjectRoot := filepath.Dir(filepath.Dir(filename))

	// construct filePath
	fileName := fmt.Sprintf("day-%s.txt", day)
	filePath := filepath.Join(porjectRoot, "../inputs", fileName)

	data, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}

	return string(data)
}
