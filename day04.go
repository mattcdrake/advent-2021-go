package main

import (
	"strconv"
	"strings"
)

type square struct {
	value  int
	marked bool
}

type board struct {
	rows [5][5]square
	won  bool
}

func (b *board) checkWin() {
	// Check rows
	for i := 0; i < len(b.rows); i++ {
		for j := 0; j < len(b.rows[i]); j++ {
			if !b.rows[i][j].marked {
				break
			}
			if j == len(b.rows[i])-1 {
				b.won = true
				return
			}
		}
	}

	// Check cols
	for i := 0; i < len(b.rows); i++ {
		for j := 0; j < len(b.rows[i]); j++ {
			if !b.rows[j][i].marked {
				break
			}
			if j == len(b.rows[i])-1 {
				b.won = true
				return
			}
		}
	}
}

func (b *board) markNumber(num int) {
	for i := 0; i < len(b.rows); i++ {
		for j := 0; j < len(b.rows[i]); j++ {
			if b.rows[i][j].value == num {
				b.rows[i][j].marked = true
				b.checkWin()
			}
		}
	}
}

func (b board) sumUnmarkedNums() int {
	sum := 0
	for i := 0; i < len(b.rows); i++ {
		for j := 0; j < len(b.rows[i]); j++ {
			if !b.rows[i][j].marked {
				sum += b.rows[i][j].value
			}
		}
	}
	return sum
}

func buildRow(rawRow string) [5]square {
	rawNums := strings.Split(rawRow, " ")
	nums := convertRawNums(rawNums)
	var squares [5]square
	for i := 0; i < 5; i++ {
		squares[i] = square{value: nums[i], marked: false}
	}
	return squares
}

func buildBoard(rawRows [5]string) board {
	rows := [5][5]square{
		buildRow(rawRows[0]),
		buildRow(rawRows[1]),
		buildRow(rawRows[2]),
		buildRow(rawRows[3]),
		buildRow(rawRows[4]),
	}
	return board{rows: rows, won: false}
}

func linesToBoards(lines []string) []board {
	boards := make([]board, 0)
	for i := 0; i < len(lines); i += 6 {
		var rawRows [5]string
		for j := 0; j < 5; j++ {
			rawRows[j] = lines[i+j]
		}
		boards = append(boards, buildBoard(rawRows))
	}
	return boards
}

func convertRawNums(rawNums []string) []int {
	nums := make([]int, 0)
	for _, rawNum := range rawNums {
		num, err := strconv.Atoi(rawNum)
		if err == nil {
			nums = append(nums, num)
		}
	}
	return nums
}

func day4p1() string {
	lines, _ := readLines("input/day04.txt")
	rawNums := strings.Split(lines[0], ",")
	nums := convertRawNums(rawNums)
	boards := linesToBoards(lines[2:])

	for _, num := range nums {
		for i, _ := range boards {
			boards[i].markNumber(num)
			if boards[i].won {
				return strconv.Itoa(boards[i].sumUnmarkedNums() * num)
			}
		}
	}

	return "nope"
}

func day4p2() string {
	lines, _ := readLines("input/day04.txt")
	rawNums := strings.Split(lines[0], ",")
	nums := convertRawNums(rawNums)
	boards := linesToBoards(lines[2:])
	boardsLeft := len(boards)
	for _, num := range nums {
		for i, _ := range boards {
			if boards[i].won {
				continue
			}

			boards[i].markNumber(num)
			if boards[i].won && boardsLeft == 1 {
				return strconv.Itoa(boards[i].sumUnmarkedNums() * num)
			} else if boards[i].won {
				boardsLeft--
			}
		}
	}

	return "nope"
}
