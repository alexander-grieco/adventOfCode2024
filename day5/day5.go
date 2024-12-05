package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"strconv"
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

var invalid [][]int

// /////////////////////////////////////// Part 2 ////////////////////////////////////////
func part2() {
	inputs := strings.Split(strings.TrimSpace(string(input)), "\n\n")
	rules := getRules(inputs[0])
	fmt.Println(calcInvalid(rules))
}

func sort(inv []int, r map[int][]int) []int {
	for i := 0; i < len(inv)-1; i++ {
		swapped := false
		for j := 0; j < len(inv)-i-1; j++ {
			if cmp(r, inv[j], inv[j+1]) {
				temp := inv[j]
				inv[j] = inv[j+1]
				inv[j+1] = temp
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}
	return inv
}

func calcInvalid(r map[int][]int) int {
	ret := 0
	for _, inv := range invalid {
		inv = sort(inv, r)
		ret += inv[len(inv)/2]
	}
	return ret
}

func cmp(r map[int][]int, i, j int) bool {
	for _, num := range r[j] {
		if i == num {
			return true
		}
	}
	return false
}

// /////////////////////////////////////// Part 1 ////////////////////////////////////////
func part1() {
	inputs := strings.Split(strings.TrimSpace(string(input)), "\n\n")
	rules := getRules(inputs[0])
	updates := getUpdates(inputs[1])

	fmt.Println(getCorrect(rules, updates))
}

func isValid(upd []int, r map[int][]int) bool {
	seen := []int{}
	for _, pg := range upd {
		for _, sn := range seen {
			for _, rpg := range r[pg] {
				if rpg == sn {
					invalid = append(invalid, upd)
					return false
				}
			}
		}
		seen = append(seen, pg)
	}
	return true
}

func getCorrect(r map[int][]int, u [][]int) int {
	count := 0
	for _, upd := range u {
		if isValid(upd, r) {
			count += upd[len(upd)/2]
		}
	}
	return count
}

func getRules(r string) map[int][]int {
	rules := map[int][]int{}
	f := strings.NewReader(r)
	s := bufio.NewScanner(f)
	for s.Scan() {
		nums := strings.Split(s.Text(), "|")
		n1, _ := strconv.Atoi(nums[0])
		n2, _ := strconv.Atoi(nums[1])
		rules[n1] = append(rules[n1], n2)
	}
	return rules
}

func getUpdates(u string) [][]int {
	updates := [][]int{}
	f := strings.NewReader(u)
	s := bufio.NewScanner(f)
	for s.Scan() {
		pages := strings.Split(s.Text(), ",")
		order := []int{}
		for _, page := range pages {
			insert, _ := strconv.Atoi(page)
			order = append(order, insert)
		}
		updates = append(updates, order)
	}
	return updates
}
