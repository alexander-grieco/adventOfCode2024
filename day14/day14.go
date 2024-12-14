package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"image"
	"strings"
	"sync"
)

// switch to input.txt to use actual input
//
//go:embed input.txt
var input string

func main() {
	f := strings.NewReader(strings.Trim(input, "\n"))
	s := bufio.NewScanner(f)
	// quads := map[int]int{}
	// part1(s, quads)
	part2(s)
}

// /////////////////////////////////////// Part 2 ////////////////////////////////////////
func part2(s *bufio.Scanner) {
	// Used to determine if using test input
	gridSizeX, gridSizeY := 101, 103
	testInput := len(strings.Split(strings.Trim(input, "\n"), "\n")) < 15
	if testInput {
		gridSizeX, gridSizeY = 11, 7
	}

	var x, y, velx, vely []int
	var mu sync.Mutex
	for s.Scan() {
		var sX, sY, vX, vY int
		fmt.Sscanf(s.Text(), "p=%d,%d v=%d,%d", &sX, &sY, &vX, &vY)
		mu.Lock()
		x = append(x, sX)
		mu.Unlock()
		mu.Lock()
		y = append(y, sY)
		mu.Unlock()
		mu.Lock()
		velx = append(velx, vX)
		mu.Unlock()
		mu.Lock()
		vely = append(vely, vY)
		mu.Unlock()
	}

	count := 1
	for count < 12000 {
		for j := 0; j < len(x); j++ {
			x[j] = (x[j] + velx[j]) % gridSizeX
			if x[j] < 0 {
				x[j] = gridSizeX + x[j]
			}
			y[j] = (y[j] + vely[j]) % gridSizeY
			if y[j] < 0 {
				y[j] = gridSizeY + y[j]
			}
		}
		if isChristmasTree2(x, y) {
			fmt.Println(count)
		}
		count++
	}
}

// This is janky, but one of the answers is the correct one...so I guess it works
func isChristmasTree2(x, y []int) bool {
	grid := map[image.Point]int{}
	for i := 0; i < len(x); i++ {
		grid[image.Point{x[i], y[i]}]++
	}

	if len(grid) == len(x) {
		return true
	}
	return false
}

// /////////////////////////////////////// Part 1 ////////////////////////////////////////
func part1(s *bufio.Scanner, quads map[int]int) {
	// Used to determine if using test input
	gridSizeX, gridSizeY := 101, 103
	testInput := len(strings.Split(strings.Trim(input, "\n"), "\n")) < 15
	if testInput {
		gridSizeX, gridSizeY = 11, 7
	}

	midX := gridSizeX / 2
	midY := gridSizeY / 2
	for s.Scan() {
		var sX, sY, vX, vY int
		fmt.Sscanf(s.Text(), "p=%d,%d v=%d,%d", &sX, &sY, &vX, &vY)

		for i := 1; i <= 100; i++ {
			sX = (sX + vX) % gridSizeX
			if sX < 0 {
				sX = gridSizeX + sX
			}
			sY = (sY + vY) % gridSizeY
			if sY < 0 {
				sY = gridSizeY + sY
			}
		}

		if sX < midX && sY < midY {
			quads[1]++
		} else if sX < midX && sY > midY {
			quads[2]++
		} else if sX > midX && sY < midY {
			quads[3]++
		} else if sX > midX && sY > midY {
			quads[4]++
		}

	}
	count := 1
	for _, quad := range quads {
		count *= quad
	}
	fmt.Println(count)
}
