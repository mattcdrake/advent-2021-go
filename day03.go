package main

import (
	"math"
	"strconv"
)

func day3p1() string {
	lines, _ := readLines("input/day03.txt")
	freqs := make([]int, len(lines[0]))

	// Add 1 for "1", subtract 1 for "0". Positive values indicate "1" was the
	// most common digit in that place...assuming no ties. :)
	for _, line := range lines {
		for i := 0; i < len(line); i++ {
			if string(line[i]) == "1" {
				freqs[i]++
			} else {
				freqs[i]--
			}
		}
	}

	power := len(freqs) - 1
	gammaRate := 0
	for i := 0; i < len(freqs); i++ {
		if freqs[i] > 0 {
			gammaRate += int(math.Pow(2, float64(power)))
		}
		power--
	}

	// Find complement
	power = len(freqs)
	limit := int(math.Pow(2, float64(power)) - 1.0)
	epsilonRate := limit - gammaRate

	return strconv.Itoa(gammaRate * epsilonRate)
}
