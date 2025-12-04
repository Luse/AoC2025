package main

import (
	"fmt"
	"log"

	"koltrast.net/AOC2025/utils"
)

const (
	EMPTY_FLOOR = '.'
	PAPER_ROLL  = '@'
)

func part1(lines []string) int {
	accessiblePaperRolls := 0
	for verticalIndex, line := range lines {
		atVertialEdgeTop, atVerticalEdgeBottom := findVerticalEdge(lines, verticalIndex)

		for horizontalIndex, char := range line {
			paperRollCount := 0
			atHorizontalEdgeLeft, atHorizontalEdgeRight := findHorizontalEdge(line, horizontalIndex)
			neighbors := []rune{}

			if !atVertialEdgeTop {
				if !atHorizontalEdgeLeft {
					neighbors = append(neighbors, rune(lines[verticalIndex-1][horizontalIndex-1])) // top-left
				}
				neighbors = append(neighbors, rune(lines[verticalIndex-1][horizontalIndex])) // top
				if !atHorizontalEdgeRight {
					neighbors = append(neighbors, rune(lines[verticalIndex-1][horizontalIndex+1])) // top-right
				}
			}

			if !atHorizontalEdgeLeft {
				neighbors = append(neighbors, rune(lines[verticalIndex][horizontalIndex-1])) // left
			}

			if !atHorizontalEdgeRight {
				neighbors = append(neighbors, rune(lines[verticalIndex][horizontalIndex+1])) // right
			}

			if !atVerticalEdgeBottom {
				if !atHorizontalEdgeLeft {
					neighbors = append(neighbors, rune(lines[verticalIndex+1][horizontalIndex-1])) // bottom-left
				}
				neighbors = append(neighbors, rune(lines[verticalIndex+1][horizontalIndex])) // bottom
				if !atHorizontalEdgeRight {
					neighbors = append(neighbors, rune(lines[verticalIndex+1][horizontalIndex+1])) // bottom-right
				}
			}

			for _, neighbor := range neighbors {
				if neighbor == PAPER_ROLL {
					paperRollCount++
				}
			}

			if char == PAPER_ROLL && paperRollCount < 4 {
				accessiblePaperRolls++
			}
		}
	}
	return accessiblePaperRolls
}

func part2(lines []string) int {
	totalRemoved := 0

	for {
		removed := 0
		newLines := make([]string, len(lines))
		copy(newLines, lines)

		for verticalIndex, line := range lines {
			atVertialEdgeTop, atVerticalEdgeBottom := findVerticalEdge(lines, verticalIndex)

			for horizontalIndex, char := range line {
				if char != PAPER_ROLL {
					continue
				}

				paperRollCount := 0
				atHorizontalEdgeLeft, atHorizontalEdgeRight := findHorizontalEdge(line, horizontalIndex)
				neighbors := []rune{}

				if !atVertialEdgeTop {
					if !atHorizontalEdgeLeft {
						neighbors = append(neighbors, rune(lines[verticalIndex-1][horizontalIndex-1])) // top-left
					}
					neighbors = append(neighbors, rune(lines[verticalIndex-1][horizontalIndex])) // top
					if !atHorizontalEdgeRight {
						neighbors = append(neighbors, rune(lines[verticalIndex-1][horizontalIndex+1])) // top-right
					}
				}

				if !atHorizontalEdgeLeft {
					neighbors = append(neighbors, rune(lines[verticalIndex][horizontalIndex-1])) // left
				}

				if !atHorizontalEdgeRight {
					neighbors = append(neighbors, rune(lines[verticalIndex][horizontalIndex+1])) // right
				}

				if !atVerticalEdgeBottom {
					if !atHorizontalEdgeLeft {
						neighbors = append(neighbors, rune(lines[verticalIndex+1][horizontalIndex-1])) // bottom-left
					}
					neighbors = append(neighbors, rune(lines[verticalIndex+1][horizontalIndex])) // bottom
					if !atHorizontalEdgeRight {
						neighbors = append(neighbors, rune(lines[verticalIndex+1][horizontalIndex+1])) // bottom-right
					}
				}

				for _, neighbor := range neighbors {
					if neighbor == PAPER_ROLL {
						paperRollCount++
					}
				}

				if paperRollCount < 4 {
					removed++
					newLines[verticalIndex] = newLines[verticalIndex][:horizontalIndex] + string(EMPTY_FLOOR) + newLines[verticalIndex][horizontalIndex+1:]
				}
			}
		}

		if removed == 0 {
			break
		}

		totalRemoved += removed
		lines = newLines
	}

	return totalRemoved
}

func main() {
	lines, err := utils.ReadLines("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Part 1:", part1(lines))
	fmt.Println("Part 2:", part2(lines))
}

func findHorizontalEdge(line string, i int) (bool, bool) {
	atHorizontalEdgeLeft := i == 0
	atHorizontalEdgeRight := i == len(line)-1
	return atHorizontalEdgeLeft, atHorizontalEdgeRight
}

func findVerticalEdge(lines []string, i int) (bool, bool) {
	atVerticalEdgeTop := i == 0
	atVerticalEdgeBottom := i == len(lines)-1

	return atVerticalEdgeTop, atVerticalEdgeBottom
}
