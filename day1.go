package main

import (
	"strconv"
)

func day1p1() string {
	var lines, _ = readLines("input/day1.txt")
	var depths []int

	for _, num := range lines {
		depth, _ := strconv.Atoi(num)
		depths = append(depths, depth)
	}

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
