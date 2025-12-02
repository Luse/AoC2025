package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"koltrast.net/AOC2025/utils"
)

func part1(lines []string) uint64 {
	var currentDuplicatePattern uint64 = 0
	for _, line := range lines {
		items := strings.Split(line, ",")
		for _, item := range items {
			value := strings.Split(item, "-")
			lower, err := strconv.ParseUint(value[0], 10, 64)
			if err != nil {
				continue
			}
			upper, err := strconv.ParseUint(value[1], 10, 64)
			if err != nil {
				fmt.Printf("Error converting upper boundary: %v\n", err)
				continue
			}
			list := getListFromBoundaries(lower, upper)
			for _, value := range list {
				pattern := findDuplicateStringPattern(value)
				if pattern != 0 {
					currentDuplicatePattern += pattern
				}
			}
		}

	}
	return currentDuplicatePattern
}

func part2(lines []string) uint64 {
	var currentDuplicatePattern uint64 = 0
	for _, line := range lines {
		items := strings.Split(line, ",")
		for _, item := range items {
			value := strings.Split(item, "-")
			lower, err := strconv.ParseUint(value[0], 10, 64)
			if err != nil {
				continue
			}
			upper, err := strconv.ParseUint(value[1], 10, 64)
			if err != nil {
				fmt.Printf("Error converting upper boundary: %v\n", err)
				continue
			}
			list := getListFromBoundaries(lower, upper)
			for _, value := range list {
				pattern := findDuplicateStringPatternElectric(value)
				if pattern != 0 {
					currentDuplicatePattern += pattern
				}
			}
		}

	}
	return currentDuplicatePattern
}

func main() {
	lines, err := utils.ReadLines("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	start := time.Now()
	result1 := part1(lines)
	duration1 := time.Since(start)
	fmt.Printf("Part 1: %d (took %v)\n", result1, duration1)

	start = time.Now()
	result2 := part2(lines)
	duration2 := time.Since(start)
	fmt.Printf("Part 2: %d (took %v)\n", result2, duration2)
}

func getListFromBoundaries(start, end uint64) []uint64 {
	list := make([]uint64, 0, end-start+1)
	for i := start; i <= end; i++ {
		list = append(list, i)
	}
	return list
}

func findDuplicateStringPattern(s uint64) uint64 {
	sStr := strconv.FormatUint(s, 10)
	n := len(sStr)

	// Must be even length to split in half
	if n%2 != 0 {
		return 0
	}

	firstHalf := sStr[0 : n/2]
	secondHalf := sStr[n/2 : n]

	if firstHalf == secondHalf {
		return s
	}
	return 0
}

func findDuplicateStringPatternElectric(s uint64) uint64 {
	sStr := strconv.FormatUint(s, 10)
	n := len(sStr)

	for patternLen := 1; patternLen <= n/2; patternLen++ {
		if n%patternLen != 0 {
			continue
		}

		pattern := sStr[0:patternLen]
		repeats := n / patternLen

		if repeats < 2 {
			continue
		}

		valid := true
		for i := 0; i < n; i += patternLen {
			if sStr[i:i+patternLen] != pattern {
				valid = false
				break
			}
		}

		if valid {
			return s
		}
	}

	return 0
}
