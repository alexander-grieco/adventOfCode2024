package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

func isPatternMatch(grid [][]rune, pattern [][]rune, startX, startY int) bool {
	patternRows, patternCols := len(pattern), len(pattern[0])
	rows, cols := len(grid), len(grid[0])

	// Ensure pattern fits within the grid bounds
	if startX+patternRows > rows || startY+patternCols > cols {
		return false
	}

	// Check each cell in the pattern
	for i := 0; i < patternRows; i++ {
		for j := 0; j < patternCols; j++ {
			// Check if the current grid cell matches the pattern (consider '.' as a wildcard)
			if pattern[i][j] != '.' && pattern[i][j] != grid[startX+i][startY+j] {
				return false
			}
		}
	}

	return true
}

func countPatternOccurrences(grid [][]rune, patterns [][][]rune) int {
	count := 0
	rows, cols := len(grid), len(grid[0])

	// Check every position in the grid as a potential top-left corner for the patterns
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			for _, pattern := range patterns {
				if isPatternMatch(grid, pattern, i, j) {
					count++
				}
			}
		}
	}

	return count
}

func main() {
	// Parse the grid
	lines := strings.Split(strings.TrimSpace(string(input)), "\n")
	runeGrid := make([][]rune, len(lines))
	for i := range lines {
		runeGrid[i] = []rune(lines[i])
	}

	// Define the patterns
	pattern1 := [][]rune{
		{'M', '.', 'S'},
		{'.', 'A', '.'},
		{'M', '.', 'S'},
	}
	pattern2 := [][]rune{
		{'M', '.', 'M'},
		{'.', 'A', '.'},
		{'S', '.', 'S'},
	}
	pattern3 := [][]rune{
		{'S', '.', 'S'},
		{'.', 'A', '.'},
		{'M', '.', 'M'},
	}
	pattern4 := [][]rune{
		{'S', '.', 'M'},
		{'.', 'A', '.'},
		{'S', '.', 'M'},
	}

	// Count occurrences of both patterns
	patterns := [][][]rune{pattern1, pattern2, pattern3, pattern4}
	count := countPatternOccurrences(runeGrid, patterns)

	fmt.Printf("The patterns appear %d times in the grid.\n", count)
}


