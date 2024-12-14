package main

import (
	_ "embed"
	"fmt"
	"image"
	"strconv"
	"strings"
)

const bigMult = 10000000000000

// switch to input.txt to use actual input
//
//go:embed input.txt
var input string

func main() {
	entries := strings.Split(strings.Trim(input, "\n"), "\n\n")
	part1(entries)
	part2(entries)
	// part2()
}

// /////////////////////////////////////// Part 2 ////////////////////////////////////////
/*
Strategy:
Look on Reddit and find out this is just Cramer's rule
Look up Cramer's rule to understand it
Use Cramer's rule
*/
func part2(entries []string) {
	count := 0
	for _, entry := range entries {
		tokens := processLargeEntry(entry)
		count += tokens
	}

	fmt.Println(count)
}

// Use Cramer's rule - Linear Algebra
// x = Det(Ax)/Det(A), y = Det(Ay)/Det(A)
// Where Ax is the X row of the matrix replaced by the solution (aka Prize)
// and Ay is the same for the Y coordinates
// The formula is A = detA/det and B = detB/det
func getMinTokens(dA, dB, prize image.Point) int {
	// The normal determinate of the matrix
	det := dA.X*dB.Y - dB.X*dA.Y

	// The determinate with the A values replaced by the prize
	detA := prize.X*dB.Y - dB.X*prize.Y
	// The determinate with the B values replaced by the prize
	detB := dA.X*prize.Y - prize.X*dA.Y

	// Check that the determinate isn't zero and that detA and detB are whole number multiples of the
	// determinate (otherwise there isn't a valid solution - you can't have partial tokens)
	if det != 0 && detA == (detA/det)*det && detB == (detB/det)*det {
		return detA/det*3 + detB/det
	}
	return 0
}

func processLargeEntry(ent string) int {
	inputs := strings.Split(ent, "\n")

	return getMinTokens(processLargeEntryInput(inputs))

}

func processLargeEntryInput(inp []string) (image.Point, image.Point, image.Point) {
	dA, dB, prize := image.Point{}, image.Point{}, image.Point{}
	for idx, row := range inp {
		x, y := getCoords(row)
		switch idx {
		case 0:
			dA.X = x
			dA.Y = y
		case 1:
			dB.X = x
			dB.Y = y
		case 2:
			prize.X = x + bigMult
			prize.Y = y + bigMult
		}
	}

	return dA, dB, prize
}

// /////////////////////////////////////// Part 1 ////////////////////////////////////////
// Test answer - 2 prizes; 480 tokens
// puzzle 1 and 3 are solvable, p1 280 tokens, p2 200 tokens
// 3 tokens for A, 1 for B-
// no more than 100 presses

// Strategy
// Search over 2D space (for _ in y for _ in x), each space in that map is the min tokens to get there. -1 if not possible

// This works for a small search space - part 2 is the "right" answer

type grid map[image.Point]int

func part1(entries []string) {
	count := 0
	for _, entry := range entries {
		tokens := processEntry(entry)
		// fmt.Println(tokens)
		count += tokens
	}

	fmt.Println(count)
}

func processEntry(ent string) int {
	inputs := strings.Split(ent, "\n")
	dA, dB, prize := processEntryInput(inputs)
	minTokens := 3*100 + 100 // Max # of tokens with A=3 and B=1
	for y := range 100 {
		for x := range 100 {
			if prize == dA.Mul(y).Add(dB.Mul(x)) {
				if y+x < minTokens {
					minTokens = 3*y + x
				}
			}
		}
	}
	if minTokens == 3*100+100 {
		return 0
	}
	return minTokens
}

func processEntryInput(inp []string) (image.Point, image.Point, image.Point) {
	dA, dB, prize := image.Point{}, image.Point{}, image.Point{}
	for idx, row := range inp {
		x, y := getCoords(row)
		switch idx {
		case 0:
			dA.X = x
			dA.Y = y
		case 1:
			dB.X = x
			dB.Y = y
		case 2:
			prize.X = x
			prize.Y = y
		}
	}

	return dA, dB, prize
}

// Extracts the coordinates from the input rows
func getCoords(row string) (int, int) {
	coordStr := strings.Split(strings.Split(row, ":")[1], ",")
	xCoord, _ := strconv.Atoi(strings.TrimFunc(coordStr[0], func(r rune) bool {
		if r == 'X' || r == 'Y' || r == '=' || r == '+' || r == ' ' {
			return true
		}
		return false
	}))
	yCoord, _ := strconv.Atoi(strings.TrimFunc(coordStr[1], func(r rune) bool {
		if r == 'X' || r == 'Y' || r == '=' || r == '+' || r == ' ' {
			return true
		}
		return false
	}))
	return xCoord, yCoord
}

// //////////////////////////// QUEUE /////////////////////////////////////////////////////
// queue implementation to help with BFS
type queue[T any] struct {
	items []T
}

func (q *queue[T]) add(item T) {
	q.items = append(q.items, item)
}

func (q *queue[T]) pop() (T, bool) {
	item := q.items[0]
	q.items = q.items[1:]
	return item, true
}

func (q *queue[T]) isEmpty() bool {
	return len(q.items) == 0
}
