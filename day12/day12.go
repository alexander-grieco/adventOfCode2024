package main

import (
	_ "embed"
	"fmt"
	"image"
	"log"
	"strings"
)

// switch to input.txt to use actual input
//
//go:embed testInput.txt
var input string

type gridPoint struct {
	str      string
	numAdj   int
	visited  bool
	numTurns int
}

type grid map[image.Point]gridPoint

func main() {
	gr := map[image.Point]gridPoint{}
	for y, s := range strings.Split(strings.TrimSpace(input), "\n") {
		for x, r := range s {
			gr[image.Point{y, x}] = gridPoint{
				str:      string(r),
				numAdj:   0,
				visited:  false,
				numTurns: 0,
			}
		}
	}
	p1 := 0
	p2 := 0
	for coord, pt := range gr {
		if !pt.visited {
			p1a, p2a := calcFenceCost(coord, gr)
			p1 += p1a
			p2 += p2a
		}
	}
	fmt.Println(p1)
	fmt.Println(p2)
}

func calcFenceCost(coord image.Point, gr grid) (int, int) {
	groupList := findTotalGroup(coord, gr)
	numSides := 0
	numDiscountedSides := 0
	delta := []image.Point{
		{0, -1}, {1, 0}, {0, 1}, {-1, 0},
	}
	for _, pt := range groupList {
		// part1
		switch gr[pt].numAdj {
		case 0:
			numSides += 4
		case 1:
			numSides += 3
		case 2:
			numSides += 2
		case 3:
			numSides += 1
		case 4:
			numSides += 0
		default:
			log.Fatalf("Invalid number of neighbors: %d", gr[pt].numAdj)
		}

		// part2

		for _, d := range delta {
			// If adjacent squares are different letters
			if adj := pt.Add(d); gr[pt].str != gr[adj].str {
				// Check if point is at a corner - if yes, add a side
				if opp := (image.Point{-d.Y, d.X}); gr[pt.Add(opp)].str != gr[pt].str || gr[pt.Add(opp).Add(d)].str == gr[pt].str {
					numDiscountedSides++
				}
			}
		}
	}

	return numSides * len(groupList), numDiscountedSides * len(groupList)
}

func findTotalGroup(coord image.Point, gr grid) []image.Point {
	delta := []image.Point{
		{0, -1}, {1, 0}, {0, 1}, {-1, 0},
	}
	group := []image.Point{}
	var q queue[image.Point]
	q.add(coord)

	for !q.isEmpty() {
		cur, _ := q.pop()
		curPt := gr[cur]
		if !curPt.visited {
			group = append(group, cur)
			curPt.visited = true
			gr[cur] = curPt
			for _, d := range delta {
				adj := cur.Add(d)
				// check if adj is in grid
				if adjPt, ok := gr[adj]; ok && !adjPt.visited {
					// check if they are same Letter
					if adjPt.str == curPt.str {
						adjPt.numAdj++
						gr[adj] = adjPt
						curPt.numAdj++
						gr[cur] = curPt
						q.add(adj)
					}
				}
			}
		}
	}
	return group
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
