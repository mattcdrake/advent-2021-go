package main

import (
	"strconv"
	"strings"
)

func day2p1() string {
	lines, _ := readLines("input/day02.txt")
	depth := 0
	horPos := 0

	for _, line := range lines {
		parts := strings.Split(line, " ")
		magnitude, _ := strconv.Atoi(parts[1])

		if parts[0] == "forward" {
			horPos += magnitude
		} else if parts[0] == "down" {
			depth += magnitude
		} else {
			depth -= magnitude
		}
	}

	return strconv.Itoa(depth * horPos)
}

func day2p2() string {
	lines, _ := readLines("input/day02.txt")
	aim := 0
	depth := 0
	horPos := 0

	for _, line := range lines {
		parts := strings.Split(line, " ")
		magnitude, _ := strconv.Atoi(parts[1])

		if parts[0] == "forward" {
			horPos += magnitude
			depth += aim * magnitude
		} else if parts[0] == "down" {
			aim += magnitude
		} else {
			aim -= magnitude
		}
	}

	return strconv.Itoa(depth * horPos)
}
