package main

import (
	"math"
	"strconv"
)

func calcFreqs(lines []string) []int {
	freqs := make([]int, len(lines[0]))

	// Add 1 for "1", subtract 1 for "0". Positive values indicate "1" was the
	// most common digit in that place.
	for _, line := range lines {
		for i := 0; i < len(line); i++ {
			if string(line[i]) == "1" {
				freqs[i]++
			} else {
				freqs[i]--
			}
		}
	}

	return freqs
}

// Filter elements that don't have the given char in the given position
func filterOut(place int, value byte, lines []string) []string {
	for i := 0; i < len(lines); i++ {
		if lines[i][place] != value {
			lines[i] = lines[len(lines)-1]
			lines = lines[0 : len(lines)-1]
			i--
		}
	}
	return lines
}

func freqsToBitCriteria(freqs []int, isOxy bool) []int {
	bitCrit := make([]int, len(freqs))
	for i, freq := range freqs {
		if isOxy {
			if freq >= 0 {
				bitCrit[i] = 1
			} else {
				bitCrit[i] = 0
			}

		} else {
			if freq >= 0 {
				bitCrit[i] = 0
			} else {
				bitCrit[i] = 1
			}

		}
	}
	return bitCrit
}

func findRating(lines []string, isOxy bool) string {
	for i := 0; i < len(lines[0]); i++ {
		freqs := calcFreqs(lines)
		bitCrit := freqsToBitCriteria(freqs, isOxy)
		lines = filterOut(i, byte(strconv.Itoa(bitCrit[i])[0]), lines)
		if len(lines) == 1 {
			break
		}
	}
	return lines[0]
}

func day3p1() string {
	lines, _ := readLines("input/day03.txt")
	freqs := calcFreqs(lines)
	for i, val := range freqs {
		if val > 0 {
			freqs[i] = 1
		} else {
			freqs[i] = 0
		}
	}

	// Find gammaRate
	power := len(freqs) - 1
	gammaRate := 0
	for i := 0; i < len(freqs); i++ {
		gammaRate += int(math.Pow(2, float64(power))) * freqs[i]
		power--
	}

	// Find epsilonRate via complement
	power = len(freqs)
	limit := int(math.Pow(2, float64(power)) - 1.0)
	epsilonRate := limit - gammaRate

	return strconv.Itoa(gammaRate * epsilonRate)
}

func day3p2() string {
	lines, _ := readLines("input/day03.txt")

	linesCopy := make([]string, len(lines))
	copy(linesCopy, lines)
	oxygenRating, _ := strconv.ParseInt(findRating(linesCopy, true), 2, 64)

	linesCopy = make([]string, len(lines))
	copy(linesCopy, lines)
	co2Rating, _ := strconv.ParseInt(findRating(linesCopy, false), 2, 64)

	return strconv.Itoa(int(oxygenRating * co2Rating))
}
