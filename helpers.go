package main

import (
	"bufio"
	"os"
	"strconv"
)

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, scanner.Err()
}

func readNums(path string) []int {
	var lines, _ = readLines(path)
	var nums []int

	for _, num := range lines {
		num, _ := strconv.Atoi(num)
		nums = append(nums, num)
	}

	return nums
}

func sumSlice(nums []int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
}
