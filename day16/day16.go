package main

import (
	_ "embed"
	"fmt"
	"image"
	"math"
	"strings"
)

// switch to input.txt to use actual input
//
//go:embed input.txt
var input string

type gridPoint struct {
	str     string
	visited bool
	dir     image.Point
	cost    int
	path    []image.Point
}

type grid map[image.Point]*gridPoint

var delta = []image.Point{
	{1, 0},
	{0, 1},
	{-1, 0},
	{0, -1},
}

type nodeState struct {
	pos image.Point
	dir image.Point
}

type scoreState struct {
	rd   nodeState
	path []image.Point
	cost int
}

func main() {
	gr := grid{}
	var start, end image.Point
	for y, s := range strings.Split(strings.TrimSpace(input), "\n") {
		for x, r := range s {
			var dir image.Point
			if string(r) == "S" {
				dir = image.Point{1, 0}
				start = image.Point{x, y}
			}
			if string(r) == "E" {
				end = image.Point{x, y}
			}
			gr[image.Point{x, y}] = &gridPoint{
				str:     string(r),
				visited: false,
				dir:     dir,
				cost:    0,
				path:    []image.Point{},
			}
		}
	}
	part1(gr, start, end)
	part2(gr, start, end)
}

// /////////////////////////////////////// Part 2 ////////////////////////////////////////
func part2(gr grid, start, end image.Point) {
	var q queue[scoreState]
	minCost := math.MaxInt
	nodesInShortestPath := map[int][]image.Point{}
	visited := map[nodeState]int{}
	ss := nodeState{pos: start, dir: image.Point{1, 0}}
	q.add(scoreState{
		rd:   ss,
		path: []image.Point{start},
		cost: 0,
	})

	for !q.isEmpty() {
		cur, _ := q.pop()
		if cur.cost > minCost {
			continue
		}
		if cur.rd.pos == end {
			if cur.cost <= minCost {
				minCost = cur.cost
				nodesInShortestPath[minCost] = append(nodesInShortestPath[minCost], cur.path...)
			}
			continue
		}

		for _, d := range delta {
			if d.Mul(-1) == cur.rd.dir {
				continue
			}

			test := cur.rd.pos.Add(d)
			if gr[test].str == "#" {
				continue
			}
			cost := cur.cost + 1
			if d != cur.rd.dir {
				cost += 1000
			}
			nState := nodeState{pos: test, dir: d}
			if oldCost, ok := visited[nState]; ok {
				if oldCost < cost {
					continue
				}
			}
			visited[nState] = cost
			newPath := make([]image.Point, len(cur.path))
			copy(newPath, cur.path)

			q.add(scoreState{
				rd:   nState,
				path: append(newPath, nState.pos),
				cost: cost,
			})
		}
	}

	nodeMap := map[image.Point]int{}
	for _, node := range nodesInShortestPath[minCost] {
		nodeMap[node]++
	}
	// for k, _ := range nodeMap {
	// 	fmt.Println(k.X, k.Y)
	// }
	// fmt.Println(nodeMap)
	fmt.Println(len(nodeMap))
	// fmt.Println(gr[image.Point{3, 8}].path)
	// fmt.Println(visited[nodeState{pos: image.Point{3, 8}, dir: image.Point{1, 0}}])
	// fmt.Println(visited[nodeState{pos: image.Point{3, 8}, dir: image.Point{0, -1}}])

}

// /////////////////////////////////////// Part 1 ////////////////////////////////////////
func part1(gr grid, start, end image.Point) {
	var q queue[image.Point]
	q.add(start)
	gr[start].visited = true

	for !q.isEmpty() {
		cur, _ := q.pop()
		for _, d := range delta {
			test := cur.Add(d)
			if gr[test].str != "#" {
				cost := gr[cur].cost + 1
				if d != gr[cur].dir {
					cost += 1000
				}
				if gr[test].visited == true {
					if gr[test].cost < cost {
						continue
					}
				}
				gr[test].visited = true
				gr[test].cost = cost
				gr[test].dir = d
				q.add(test)
			}
		}
	}
	fmt.Println(gr[end].cost)
}

// //// Queue
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
