package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
)

// switch to testInput.txt to use test input
//
//go:embed input.txt
var input string

func main() {
	part1()
	part2()
}

var (
	wordRe = regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)`)
	numRe  = regexp.MustCompile(`\d{1,3}`)
	dontRe = regexp.MustCompile(`don\'t\(\).*?do\(\)`) // have to do this non-greedily (note the `?`)
)

func getSum2(s string) int {
	// remove all sections between a don't() and a do()
	newStr := dontRe.ReplaceAllString(s, "")

	// Only find matches on the new string
	matches := wordRe.FindAllString(newStr, -1)

	// initialize sum
	sum := 0
	for _, m := range matches {
		sum += calcMult2(m)
	}

	return sum
}

func part2() {
	sum := 0

	// Have to make input all one line for my solution to work
	newInput := strings.ReplaceAll(input, "\n", "")

	// Same as part1
	f := strings.NewReader(newInput)
	s := bufio.NewScanner(f)
	for s.Scan() {
		sum += getSum2(s.Text())
	}
	fmt.Println(sum)
}

func calcMult2(m string) int {
	nums := numRe.FindAllString(m, -1)
	n1, err := strconv.Atoi(nums[0])
	if err != nil {
		log.Fatalf("Error converting number 1: %s", err)
	}
	n2, err := strconv.Atoi(nums[1])
	if err != nil {
		log.Fatalf("Error converting number 2: %s", err)
	}
	return n1 * n2
}

// ///////////////////////////////////////// Part 1 ////////////////////////////
func part1() {
	sum := 0
	f := strings.NewReader(input)
	s := bufio.NewScanner(f)
	for s.Scan() {
		sum += getSum(s.Text())
	}
	fmt.Println(sum)
}

func getSum(s string) int {
	matches := wordRe.FindAllString(s, -1)

	// initialize sum
	sum := 0
	for _, m := range matches {
		sum += calcMult2(m)
	}

	return sum
}

func calcMult(m string) int {
	nums := numRe.FindAllString(m, -1)
	n1, err := strconv.Atoi(nums[0])
	if err != nil {
		log.Fatalf("Error converting number 1: %s", err)
	}
	n2, err := strconv.Atoi(nums[1])
	if err != nil {
		log.Fatalf("Error converting number 2: %s", err)
	}
	return n1 * n2
}
