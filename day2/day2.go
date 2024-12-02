package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"log"
	"strconv"
	"strings"
)

// switch to testInput.txt to use test input
//
//go:embed input.txt
var input string

func main() {
	// part1()
	part2()
}

func part2() {
	safe := 0
	f := strings.NewReader(input)
	s := bufio.NewScanner(f)
	for s.Scan() {
		safe += isSafe2(s.Text())
	}
	fmt.Println(safe)
}

func isSafe2(l string) int {
	nums, err := getList(l)
	if err != nil {
		log.Fatalf("Error converting line to list of nums: %s\n", err)
	}

	if safeTest2(nums, 1) || safeTest2(nums, -1) {
		return 1
	}

	return 0
}

func safeTest2(l []int, dir int) bool {
	for i := 1; i < len(l); i++ {
		diff := (l[i] - l[i-1]) * dir
		if diff > 3 || diff < 1 {
			// inelegant but need to copy to new slices to not affect underlying array
			new1, new2 := make([]int, len(l)), make([]int, len(l))
			copy(new1, l)
			copy(new2, l)

			/* If there is an issue with the original array of numbers, run the original test (from part 1)
			 against a shortened list of numbers. We check removing both the current number or the preceeding
			number. If neither or these returns true, we can fail this entry, but if the shortened list returns
			true then this is now a valid entry
			*/
			if !safeTest(append(new1[:i], new1[i+1:]...), dir) && !safeTest(append(new2[:i-1], new2[i:]...), dir) {
				return false
			}
		}
	}
	return true
}

func part1() {
	safe := 0
	f := strings.NewReader(input)
	s := bufio.NewScanner(f)
	for s.Scan() {
		safe += isSafe(s.Text())
	}
	fmt.Println(safe)
}

func getList(l string) ([]int, error) {
	var nums []int
	items := strings.Split(l, " ")

	for _, item := range items {
		n, err := strconv.Atoi(item)
		if err != nil {
			return []int{}, err
		}
		nums = append(nums, n)
	}
	return nums, nil
}

func safeTest(l []int, dir int) bool {
	for i := 1; i < len(l); i++ {
		diff := (l[i] - l[i-1]) * dir
		if diff > 3 || diff < 1 {
			return false
		}
	}
	return true
}

func isSafe(l string) int {
	nums, err := getList(l)
	if err != nil {
		log.Fatalf("Error converting line to list of nums: %s\n", err)
	}

	if safeTest(nums, 1) || safeTest(nums, -1) {
		return 1
	}

	return 0
}
