package main

import (
	"testing"

	"koltrast.net/AOC2025/utils"
)

func TestPart1(t *testing.T) {
	input, err := utils.ReadLines("input_test.txt")
	if err != nil {
		t.Fatalf("Failed to read test input: %v", err)
	}

	result := part1(input)
	var expected uint64 = 1227775554

	if result != expected {
		t.Errorf("part1() = %d, want %d", result, expected)
	}
}

func TestPart2(t *testing.T) {
	input, err := utils.ReadLines("input_test.txt")
	if err != nil {
		t.Fatalf("Failed to read test input: %v", err)
	}

	result := part2(input)
	expected := uint64(4174379265)

	if result != expected {
		t.Errorf("part2() = %d, want %d", result, expected)
	}
}
