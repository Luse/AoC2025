#!/bin/bash
# Usage: ./create_day.sh 02

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

# Create the day folder
mkdir -p "$foldername"

# Create main.go in the folder
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

echo "Created $foldername/"
echo "Don't forget to add your input.txt to $foldername/"
echo "Run with: cd $foldername && go run main.go"
