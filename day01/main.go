package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"koltrast.net/AOC2025/utils"
)

const (
	Left  = "L"
	Right = "R"
)

func part1(lines []string) int {
	timesDialTurnedToZero := 0
	currentDialPosition := 50

	for _, line := range lines {
		currentDialPosition, _ = turnDial(line, currentDialPosition)
		if currentDialPosition == 0 {
			timesDialTurnedToZero++
		}
	}
	return timesDialTurnedToZero
}

func part2(lines []string) int {
	timesDialTurnedToZero := 0
	currentDialPosition := 50

	for _, line := range lines {
		crossed := 0
		currentDialPosition, crossed = turnDial(line, currentDialPosition)
		timesDialTurnedToZero += crossed
	}
	return timesDialTurnedToZero
}

func main() {
	lines, err := utils.ReadLines("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Part 1:", part1(lines))
	fmt.Println("Part 2:", part2(lines))
}

func turnDial(instruction string, current int) (int, int) {
	direction := strings.TrimRight(instruction, "0123456789")
	valueStr := strings.TrimLeft(instruction, "LR")
	value, err := strconv.Atoi(valueStr)
	if err != nil {
		return current, 0
	}

	timesCrossedZero := 0

	switch direction {
	case Left:
		for i := 0; i < value; i++ {
			if current == 0 {
				current = 99
			} else {
				current--
			}
			if current == 0 {
				timesCrossedZero++
			}
		}
		return current, timesCrossedZero
	case Right:
		for i := 0; i < value; i++ {
			if current == 99 {
				current = 0
			} else {
				current++
			}
			if current == 0 {
				timesCrossedZero++
			}
		}
		return current, timesCrossedZero
	}
	return current, 0
}
