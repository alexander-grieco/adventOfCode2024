package main

// import (
// 	_ "embed"
// 	"fmt"
// 	"strings"
// )
//
// //go:embed testInput.txt
// var gridFile string
//
// var directions = [][]int{
// 	{0, 1},   // Right
// 	{1, 0},   // Down
// 	{0, -1},  // Left
// 	{-1, 0},  // Up
// 	{-1, -1}, // Up-Left Diagonal
// 	{-1, 1},  // Up-Right Diagonal
// 	{1, -1},  // Down-Left Diagonal
// 	{1, 1},   // Down-Right Diagonal
// }
//
// func isValid(x, y, rows, cols int) bool {
// 	return x >= 0 && x < rows && y >= 0 && y < cols
// }
//
// func searchWord(grid [][]rune, word string, x, y, dx, dy, index int) bool {
// 	if index == len(word) {
// 		return true
// 	}
//
// 	rows, cols := len(grid), len(grid[0])
// 	if !isValid(x, y, rows, cols) || grid[x][y] != rune(word[index]) {
// 		return false
// 	}
//
// 	return searchWord(grid, word, x+dx, y+dy, dx, dy, index+1)
// }
//
// func countOccurrences(grid [][]rune, word string) int {
// 	count := 0
// 	rows, cols := len(grid), len(grid[0])
// 	for i := 0; i < rows; i++ {
// 		for j := 0; j < cols; j++ {
// 			if grid[i][j] == rune(word[0]) {
// 				for _, d := range directions {
// 					if searchWord(grid, word, i, j, d[0], d[1], 0) {
// 						count++
// 					}
// 				}
// 			}
// 		}
// 	}
// 	return count
// }
//
// func main() {
//
// 	// Parse the grid
// 	lines := strings.Split(strings.TrimSpace(string(gridFile)), "\n")
// 	runeGrid := make([][]rune, len(lines))
// 	for i := range lines {
// 		runeGrid[i] = []rune(lines[i])
// 	}
// 	word := "XMAS"
// 	count := countOccurrences(runeGrid, word)
// 	fmt.Printf("The word %s appears %d times in the grid.\n", word, count)
// }
