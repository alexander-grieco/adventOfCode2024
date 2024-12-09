package main

import (
	_ "embed"
	"fmt"
	"sort"
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

// /////////////////////////////////////// Part 2 ////////////////////////////////////////
func part2() {

	idxStrArr := buildIndexStrArr(strings.Trim(input, "\n"), 2)
	idxArr, idxMap := buildIndexArr2(idxStrArr)

	// sort Maps keys in descending order
	keys := []int{}
	for key := range idxMap {
		keys = append(keys, key)
	}

	// Step 2: Sort keys in descending order
	sort.Sort(sort.Reverse(sort.IntSlice(keys)))

	for _, k := range keys {
		cur := 0
		numLen := idxMap[k].num
		numPos := idxMap[k].pos
		for cur < numPos {
			// Navigate to the next free space
			for idxArr[cur] != "." {
				cur++
			}

			// If the free space is to the right of the entry, exit
			if cur > numPos {
				break
			}

			// Calucate the size of the free space
			dotLen := 1
			for idxArr[cur+dotLen] == "." {
				dotLen++
			}

			// If a space is big enough, swap the values
			if numLen <= dotLen {
				for i := 0; i < numLen; i++ {
					temp := idxArr[numPos+i]
					idxArr[numPos+i] = "."
					idxArr[cur+i] = temp
				}
				break // We don't want to switch if we already did

			} else {
				// If space isn't big enough, go check the next space
				cur += dotLen
			}
		}
	}

	fmt.Printf("%d\n", calcSha(idxArr))
}

type arrMap struct {
	num int
	pos int
}

func buildIndexArr2(str []string) ([]string, map[int]arrMap) {
	var idxArr []string
	idxMap := map[int]arrMap{}
	totIdx := 0
	for idx, s := range str {
		id, _ := strconv.Atoi(string(s[0]))
		free, _ := strconv.Atoi(string(s[1]))
		idxMap[idx] = arrMap{num: id, pos: totIdx}
		for i := 0; i < id; i++ {
			totIdx++
			idxArr = append(idxArr, strconv.Itoa(idx))
		}
		for i := 0; i < free; i++ {
			totIdx++
			idxArr = append(idxArr, ".")
		}
	}
	return idxArr, idxMap
}

// /////////////////////////////////////// Part 1 ////////////////////////////////////////
func part1() {
	idxStrArr := buildIndexStrArr(strings.Trim(input, "\n"), 2)
	idxArr := buildIndexArr(idxStrArr)
	// fmt.Printf("%#v", idxArr)

	cur := 0
	curEnd := len(idxArr) - 1
	for cur < curEnd {
		if idxArr[cur] == "." {
			temp := idxArr[curEnd]
			idxArr[curEnd] = "."
			idxArr[cur] = temp
			curEnd--
		}
		for idxArr[curEnd] == "." {
			curEnd--
		}
		cur++
	}
	fmt.Printf("%d\n", calcSha(idxArr))
}

func calcSha(str []string) int {
	count := 0
	for idx, s := range str {
		num, err := strconv.Atoi(s)
		if err != nil {
			continue
		}
		count += idx * num
	}

	return count
}

func buildIndexArr(str []string) []string {
	var idxArr []string
	for idx, s := range str {
		id, _ := strconv.Atoi(string(s[0]))
		free, _ := strconv.Atoi(string(s[1]))
		// fmt.Printf("Num: %d, free: %d", id, free)
		for i := 0; i < id; i++ {
			idxArr = append(idxArr, strconv.Itoa(idx))
		}
		for i := 0; i < free; i++ {
			idxArr = append(idxArr, ".")
		}
	}
	return idxArr
}

func buildIndexStrArr(s string, size int) []string {
	var ret []string
	for i := 0; i < len(s)-1; i += size {
		ret = append(ret, input[i:i+size])
	}
	ret = append(ret, s[len(s)-1:]+"0")

	return ret
}
