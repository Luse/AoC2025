#!/bin/bash

if [ -z "$1" ]; then
    echo "Usage: ./create_day.sh <day_number>"
    echo "Example: ./create_day.sh 02"
    exit 1
fi

day=$1
foldername="day${day}"

if [ -d "$foldername" ]; then
    echo "Error: $foldername already exists!"
    exit 1
fi

mkdir -p "$foldername"

cat > "$foldername/main.go" << 'EOF'
package main

import (
	"fmt"
	"log"
	
	"koltrast.net/AOC2025/utils"
)

func part1(lines []string) int {
	// TODO: Implement Part 1
	return 0
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
EOF

cat > "$foldername/main_test.go" << 'EOF'
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
	expected := 0

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
	expected := 0

	if result != expected {
		t.Errorf("part2() = %d, want %d", result, expected)
	}
}
EOF

echo "Created $foldername/"
echo "Don't forget to add your input.txt to $foldername/"
echo "Run with: cd $foldername && go run main.go"
