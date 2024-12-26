package main

import (
	"aoc-2024/internal/day1"
	"aoc-2024/internal/day2"
	"aoc-2024/internal/day3"
	"aoc-2024/internal/day4"
	"aoc-2024/internal/day5"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

type dayPartRunFunction func(string) any

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
		"31": day3.Day3Part1,
		"32": day3.Day3Part2,
		"41": day4.Day4Part1,
		"42": day4.Day4Part2,
		"51": day5.Day5Part1,
		"52": day5.Day5Part2,
	}

	// parse arguments
	day := args[1]
	part := args[2]
	isExample := len(args) >= 4 && (args[3] == "1" || strings.ToLower(args[3]) == "true")
	funcCallKey := fmt.Sprintf("%s%s", day, part)

	// try to call function by day-part, if not print error
	if funcCall, ok := dayPartMap[funcCallKey]; ok {
		res := funcCall(getDataFromDayFile(day, isExample))
		fmt.Printf("Day %s | Part %s\n%v\n", day, part, res)
	} else {
		fmt.Println(fmt.Errorf("Day %v part %v does not exist.", day, part))
	}
}

func getDataFromDayFile(day string, isExample bool) string {
	// get project root
	_, filename, _, _ := runtime.Caller(0)
	porjectRoot := filepath.Dir(filepath.Dir(filename))

	// construct filePath
	var fileName string
	if isExample {
		fileName = fmt.Sprintf("day-%s-example.txt", day)
	} else {
		fileName = fmt.Sprintf("day-%s.txt", day)
	}
	filePath := filepath.Join(porjectRoot, "../inputs", fileName)

	data, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}

	return string(data)
}
