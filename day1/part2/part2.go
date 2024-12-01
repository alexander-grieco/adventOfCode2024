package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

/*
Finding sum of distances
Strategy:
- read input to two sorted slices
- Go through slices and add up differences (absolute value here)

Maybe use a structure for my list?
Inserting is important, we want to use a binary insert method, but not sure how
memory efficient that is
*/

func main() {
	// Open file
	f, err := os.Open("./day1/input.txt")
	// f, err := os.Open("./day1/testInput.txt") // for testing
	if err != nil {
		log.Fatalf("Error opening file: %s", err)
	}
	defer f.Close()

	// Create data structs
	l1, l2 := []int{}, map[int]int{}

	s := bufio.NewScanner(f)
	for s.Scan() {
		text := s.Text()
		items := strings.Split(text, "   ")
		if len(items) != 2 {
			log.Fatalf("Invalid line. Expecting two inputs, found %d", len(items))
		}

		num1, err := strconv.Atoi(items[0])
		if err != nil {
			log.Fatalf("Error converting number 1: %s", err)
		}

		num2, err := strconv.Atoi(items[1])
		if err != nil {
			log.Fatalf("Error converting number 2: %s", err)
		}

		l1 = binaryInsert(num1, l1)
		mapInsert(num2, l2)
	}
	fmt.Println(calcSimilarity(l1, l2))
}

func calcSimilarity(l1 []int, l2 map[int]int) int {
	sim := 0
	for _, num := range l1 {
		sim += num * l2[num]
	}
	return sim
}

func mapInsert(num int, m map[int]int) {
	m[num]++
}

func binarySearch(list []int, num, start, end int) int {
	if start == end {
		if list[start] > num {
			return start
		}
		return start + 1
	}

	// Moving off end of array
	if start > end {
		return start
	}

	// Binary search part
	mid := (start + end) / 2
	// fmt.Printf("Mid is %d\n", mid)
	switch {
	case list[mid] < num:
		return binarySearch(list, num, mid+1, end)
	case list[mid] > num:
		return binarySearch(list, num, start, mid-1)
	default:
		return mid
	}
}

func binaryInsert(num int, list []int) []int {
	// If list is empty
	if len(list) == 0 {
		return []int{num}
	}

	idx := binarySearch(list, num, 0, len(list)-1)
	// fmt.Printf("Index: %d, Length of list: %d, List: %#v, To insert: %d\n", idx, len(list), list, num)
	list = append(list[:idx], append([]int{num}, list[idx:]...)...)

	return list
}
