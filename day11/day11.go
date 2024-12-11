package main

import (
	_ "embed"
	"fmt"
	"math"
	"strconv"
	"strings"
)

/*
	- If the stone is engraved with the number 0, it is replaced by a stone engraved with the number 1.

	- If the stone is engraved with a number that has an even number of digits, it is replaced by two stones. The left half of the digits are engraved on the new left stone, and the right half of the digits are engraved on the new right stone. (The new numbers don't keep extra leading zeroes: 1000 would become stones 10 and 0.)

	- If none of the other rules apply, the stone is replaced by a new stone; the old stone's number multiplied by 2024 is engraved on the new stone.

*/

// switch to input.txt to use actual input
//
//go:embed input.txt
var input string

func main() {
	strNums := strings.Split(strings.Trim(input, "\n"), " ")
	numsMap := map[int]int{}
	for _, str := range strNums {
		num, _ := strconv.Atoi(str)
		numsMap[num] += 1
	}
	solution(numsMap, 25)
	numsMap = map[int]int{}
	for _, str := range strNums {
		num, _ := strconv.Atoi(str)
		numsMap[num] += 1
	}
	solution(numsMap, 75)
}

func solution(nums map[int]int, numBlinks int) {
	for i := 0; i < numBlinks; i++ {
		cMap := copyMap(nums)
		for k, v := range cMap {

			// Test if stone is 0
			if k == 0 {
				nums[1] += v
				if nums[k] == v {
					delete(nums, k)
				} else {
					nums[k] -= v
				}
				continue
			}

			// Test if even digits
			numDigits := len(strconv.Itoa(k))
			if numDigits%2 == 0 {
				pow := numDigits / 2
				n1 := k / int(math.Pow10(pow))
				n2 := k - n1*int(math.Pow10(pow))
				if nums[k] == v {
					delete(nums, k)
				} else {
					nums[k] -= v
				}
				nums[n1] += v
				nums[n2] += v
				continue
			}

			// Just multiply it by 2024
			nums[k*2024] += v
			if nums[k] == v {
				delete(nums, k)
			} else {
				nums[k] -= v
			}
		}
	}
	fmt.Println(totNums(nums))
}

// Calculate the total number of values
func totNums(m map[int]int) int {
	count := 0
	for _, val := range m {
		count += val
	}
	return count
}

// Not the most efficient, but helps not mess up the for loop
func copyMap(m map[int]int) map[int]int {
	copyM := map[int]int{}
	for key, val := range m {
		copyM[key] = val
	}
	return copyM
}
