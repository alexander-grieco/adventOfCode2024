package main

import (
	_ "embed"
	"errors"
	"fmt"
	"image"
	"strings"
)

// switch to input.txt to use actual input
//
//go:embed input.txt
var input string

func main() {
	part1()
	part2()
}

// /////////////////////////////////////// Part 2 ////////////////////////////////////////
func part2() {
	delta := []image.Point{
		{0, -1}, {1, 0}, {0, 1}, {-1, 0},
	}

	cur := image.Point{}
	dir := 5
	search := []image.Point{}
	grid := map[image.Point]rune{}
	for y, s := range strings.Split(strings.TrimSpace(string(input)), "\n") {
		for x, r := range s {
			if strings.ContainsRune("^<>v", r) {
				cur = image.Point{x, y}
				grid[image.Point{x, y}] = rune('X')
				switch r {
				case rune('^'):
					dir = 0
				case rune('>'):
					dir = 1
				case rune('v'):
					dir = 2
				case rune('<'):
					dir = 3
				}
			} else if r == rune('.') {
				search = append(search, image.Point{x, y})
				grid[image.Point{x, y}] = r
			} else {
				grid[image.Point{x, y}] = r
			}
		}
	}

	points := []image.Point{}
	for _, pt := range search {
		ret, err := moveGrid2(grid, cur, delta, dir, pt)
		if err == nil {
			points = append(points, ret)
		}
	}
	fmt.Println(len(points))
}

type loop struct {
	p image.Point
	d int
}

func moveGrid2(grid map[image.Point]rune, cur image.Point, delta []image.Point, dir int, pt image.Point) (image.Point, error) {
	grid[pt] = rune('#')
	defer func() {
		grid[pt] = rune('.')
	}()

	visited := map[loop]int{}
	for {
		if _, ok := visited[loop{
			cur,
			dir,
		}]; ok {
			return pt, nil
		}
		r, ok := grid[cur]
		if !ok {
			return image.Point{}, errors.New("Nope")
		}
		switch r {
		case rune('#'):
			cur = cur.Sub(delta[dir])
			dir = (dir + 1) % 4
		default:
			visited[loop{
				cur,
				dir,
			}] += 1
		}
		cur = cur.Add(delta[dir])
	}
}

// /////////////////////////////////////// Part 1 ////////////////////////////////////////
func part1() {
	delta := []image.Point{
		{0, -1}, {1, 0}, {0, 1}, {-1, 0},
	}

	cur := image.Point{}
	dir := 5
	grid := map[image.Point]rune{}
	for y, s := range strings.Split(strings.TrimSpace(string(input)), "\n") {
		for x, r := range s {
			if strings.ContainsRune("^<>v", r) {
				cur = image.Point{x, y}
				grid[image.Point{x, y}] = rune('X')
				switch r {
				case rune('^'):
					dir = 0
				case rune('>'):
					dir = 1
				case rune('v'):
					dir = 2
				case rune('<'):
					dir = 3
				}
			} else {
				grid[image.Point{x, y}] = r
			}
		}
	}
	moveGrid(grid, cur, delta, dir)

}

func moveGrid(grid map[image.Point]rune, cur image.Point, delta []image.Point, dir int) {
	visited := map[image.Point]int{}
	for {
		r, ok := grid[cur]
		if !ok {
			break // exit
		}
		switch r {
		case rune('#'):
			cur = cur.Sub(delta[dir])
			dir = (dir + 1) % 4
		default:
			visited[cur] += 1
		}
		cur = cur.Add(delta[dir])
	}
	fmt.Println(len(visited))
}
