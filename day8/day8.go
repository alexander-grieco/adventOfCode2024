package main

import (
	_ "embed"
	"fmt"
	"image"
	"strings"
)

// switch to input.txt to use actual input
//
//go:embed input.txt
var input string

func main() {
	grid := map[image.Point]string{}
	symbols := map[string][]image.Point{}
	for y, s := range strings.Split(strings.TrimSpace(string(input)), "\n") {
		for x, r := range s {
			grid[image.Point{y, x}] = string(r)
			if r != '.' {
				symbols[string(r)] = append(symbols[string(r)], image.Point{y, x})
			}
		}
	}
	fmt.Println(calcAntiNodes(grid, symbols, 1))
	fmt.Println(calcAntiNodes(grid, symbols, 2))
}

func getValidAntiNodes(one, two image.Point, g map[image.Point]string, a map[image.Point]int) map[image.Point]int {
	testPoints := []image.Point{one.Sub(two.Sub(one)), two.Add(two.Sub(one))}
	for _, point := range testPoints {
		if _, ok := g[point]; ok {
			a[point] += 1
		}
	}
	return a
}

func getValidAntiNodes2(one, two image.Point, g map[image.Point]string, a map[image.Point]int) map[image.Point]int {
	testPoints := []image.Point{}
	idx := 0

	for {
		p1 := one.Sub(two.Sub(one).Mul(idx))
		p2 := two.Add(two.Sub(one).Mul(idx))
		_, ok1 := g[p1]
		_, ok2 := g[p2]
		if !ok1 && !ok2 {
			break
		}
		testPoints = append(testPoints, p1)
		testPoints = append(testPoints, p2)
		idx++
	}

	for _, point := range testPoints {
		if _, ok := g[point]; ok {
			a[point] += 1
		}
	}

	return a
}

func calcAntiNodes(g map[image.Point]string, s map[string][]image.Point, problem int) int {
	antiNode := map[image.Point]int{}
	for k, v := range s {
		for i := 0; i < len(v)-1; i++ {
			for j := i + 1; j < len(v); j++ {
				if problem == 1 {
					getValidAntiNodes(s[k][i], s[k][j], g, antiNode)
				} else {
					getValidAntiNodes2(s[k][i], s[k][j], g, antiNode)
				}
			}
		}
	}

	return len(antiNode)
}
