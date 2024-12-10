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

type Grid map[image.Point]int
type Slocs []image.Point
type Visited map[image.Point]bool

func main() {

	start := []image.Point{}
	grid := map[image.Point]int{}
	for y, s := range strings.Split(strings.TrimSpace(string(input)), "\n") {
		for x, r := range s {
			num := int(r - '0')
			if num == 0 {
				start = append(start, image.Point{x, y})
			}
			grid[image.Point{x, y}] = num
		}
	}
	part1(grid, start)
	part2(grid, start)
}

// /////////////////////////////////////// Part 2 ////////////////////////////////////////
func part2(g Grid, s Slocs) {
	count := 0
	for _, sPos := range s {
		count += findAllFullPaths(g, sPos)
	}
	fmt.Println(count)

}

func findAllFullPaths(g Grid, sPos image.Point) int {
	delta := []image.Point{
		{0, -1}, {1, 0}, {0, 1}, {-1, 0},
	}
	count := 0
	var q queue[image.Point]
	q.add(sPos)

	for !q.isEmpty() {
		cur, _ := q.pop()
		curNum := g[cur]
		for _, d := range delta {
			adj := cur.Add(d)
			if _, ok := g[adj]; ok {
				adjNum := g[adj]
				if adjNum-curNum == 1 {
					if adjNum == 9 {
						count++
					}
					q.add(adj)
				}
			}
		}
	}

	return count
}

// /////////////////////////////////////// Part 1 ////////////////////////////////////////
func part1(g Grid, s Slocs) {
	count := 0
	for _, sPos := range s {
		v := Visited{}
		count += findFullPaths(g, sPos, v)
	}
	fmt.Println(count)
}

func findFullPaths(g Grid, sPos image.Point, v Visited) int {
	delta := []image.Point{
		{0, -1}, {1, 0}, {0, 1}, {-1, 0},
	}
	count := 0
	var q queue[image.Point]
	q.add(sPos)

	for !q.isEmpty() {
		cur, _ := q.pop()
		curNum := g[cur]
		for _, d := range delta {
			adj := cur.Add(d)
			if _, ok := g[adj]; !v[adj] && ok {
				adjNum := g[adj]
				if adjNum-curNum == 1 {
					v[adj] = true
					if adjNum == 9 {
						count++
					} else {
						q.add(adj)
					}
				}
			}
		}
	}

	return count
}

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
