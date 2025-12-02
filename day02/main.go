package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"koltrast.net/AOC2025/utils"
)

func part1(lines []string) int {
	currentDuplicatePattern := 0
	for _, line := range lines {
		items := strings.Split(line, ",")
		for _, item := range items {
			value := strings.Split(item, "-")
			lower, err := strconv.Atoi(value[0])
			if err != nil {
				continue
			}
			upper, err := strconv.Atoi(value[1])
			if err != nil {
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

func part2(lines []string) int {
	// TODO: Implement Part 2
	return 0
}

func main() {
	lines, err := utils.ReadLines("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Part 1:", part1(lines))
	fmt.Println("Part 2:", part2(lines))
}

func getListFromBoundaries(start, end int) []int {
	list := make([]int, 0, end-start+1)
	for i := start; i <= end; i++ {
		list = append(list, i)
	}
	return list
}

func findDuplicateStringPattern(s int) int {
	sStr := strconv.Itoa(s)
	n := len(sStr)
	firstHalf := sStr[0 : n/2]
	secondHalf := sStr[n/2 : n]

	if firstHalf == secondHalf {
		return s
	}
	return 0
}
