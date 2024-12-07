package main

import (
	_ "embed"
	"fmt"
	"math"
	"strconv"
	"strings"
)

// switch to input.txt to use actual input
//
//go:embed input.txt
var input string

func main() {
	// Ingest input
	testVals, nums := []int{}, [][]int{}
	ops_p1 := []string{"+", "*"}
	for _, l := range strings.Split(strings.TrimSpace(string(input)), "\n") {
		vals := strings.Split(l, ": ")
		testVals = append(testVals, func() int {
			val, _ := strconv.Atoi(vals[0])
			return val
		}())
		nums = append(nums, getNums(vals[1]))
	}

	// Count the number of correct answer
	// Problem 1
	count := 0
	for idx, tv := range testVals {
		num := calcCalibrations(tv, nums[idx], ops_p1)
		count += num
	}
	fmt.Println(count)

	// Problem 2
	count = 0
	ops_p2 := []string{"+", "*", "||"}
	for idx, tv := range testVals {
		num := calcCalibrations(tv, nums[idx], ops_p2)
		count += num
	}
	fmt.Println(count)
}

func calcCalibrations(tv int, nums []int, ops []string) int {
	mem := map[int][]int{}
	mem[0] = []int{nums[0]}
	for idx := 1; idx < len(nums); idx++ {
		for j := range len(ops) {
			switch ops[j] {
			case "+":
				for _, x := range mem[idx-1] {
					mem[idx] = append(mem[idx], x+nums[idx])
				}
			case "*":
				for _, x := range mem[idx-1] {
					mem[idx] = append(mem[idx], x*nums[idx])
				}
			case "||":
				for _, x := range mem[idx-1] {
					numDigits := len(strconv.Itoa(nums[idx]))
					mem[idx] = append(mem[idx], x*int(math.Pow(10, float64(numDigits)))+nums[idx])
				}
			}
		}
	}

	// Test if one of the final nums is correct
	for _, res := range mem[len(nums)-1] {
		if res == tv {
			return tv
		}
	}
	return 0
}

func getNums(s string) []int {
	vals := strings.Split(s, " ")
	var ret []int
	for _, val := range vals {
		ret = append(ret, func() int {
			ival, _ := strconv.Atoi(val)
			return ival
		}())
	}

	return ret
}
