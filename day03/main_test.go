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
	expected := 357 // TODO: Add expected result

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
	expected := 0 // TODO: Add expected result

	if result != expected {
		t.Errorf("part2() = %d, want %d", result, expected)
	}
}
