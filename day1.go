package main

import (
	"strconv"
)

func day1p1() string {
	depths := readNums("input/day1.txt")

	curDepth := depths[0]
	increases := 0
	for _, depth := range depths {
		if depth > curDepth {
			increases++
		}
		curDepth = depth
	}

	return strconv.Itoa(increases)
}

func day1p2() string {
	depths := readNums("input/day1.txt")

	lastPos := len(depths) - 3
	lastSum := sumSlice(depths[0:3])
	increases := 0
	for i := 1; i <= lastPos; i++ {
		sum := sumSlice(depths[i : i+3])
		if sum > lastSum {
			increases++
		}
		lastSum = sum
	}

	return strconv.Itoa(increases)
}
