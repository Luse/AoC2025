package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"koltrast.net/AOC2025/utils"
)

func part1(lines []string) int {
	sum := []int{}
	total := 0

	for _, line := range lines {
		lline := strings.Split(line, "")
		firstBattery := 0
		secondBattery := 0
		firstBattery, firstBatteryIndex := largestNumberInSlice(lline)
		if firstBatteryIndex == len(lline)-1 {
			lline = lline[:len(lline)-1]
			secondBattery = firstBattery
		} else {
			lline = lline[firstBatteryIndex+1:]
		}
		if secondBattery == 0 {
			secondBattery, _ = largestNumberInSlice(lline)
		} else {
			firstBattery, _ = largestNumberInSlice(lline)
		}
		composite := ""

		composite = strconv.FormatInt(int64(firstBattery), 10) + strconv.FormatInt(int64(secondBattery), 10)
		compositeInt, err := strconv.Atoi(composite)
		if err != nil {
			log.Printf("Error converting composite string to int: %v\n", err)
			continue
		}
		sum = append(sum, compositeInt)
	}
	for _, val := range sum {
		total += val
	}
	return total
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

func largestNumberInSlice(numbers []string) (int, int) {
	largest := -1
	index := 0

	for i, numStr := range numbers {

		num, err := strconv.Atoi(numStr)
		if err != nil {
			log.Printf("Error converting string to int: %v\n", err)
			continue
		}
		if num > largest {
			largest = num
			index = i

		}
	}
	return largest, index
}
