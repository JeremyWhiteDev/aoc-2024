package main

import (
	"strings"
)

func day4Part1() int {
	s := getData("4")
	total := 0
	lines := strings.FieldsFunc(s, func (r rune) bool { return r == '\n' })
	// iterate over each line
	for i, line := range lines {
		var indexesOfX []int
		for j, c := range line {
			if c == 'X' {indexesOfX = append(indexesOfX, j)}
		}
		// day4 brute force

		// todo, refactor to cursor that "moves"
		// provide cursor with x/y directions. if cursor is given direction it can't go (is at it's limit), don't move it.


		// for each index of x, do the search for backwards, forwards, up, down, and 4 diagnoal directions (safely)
		for _, indexOfX := range indexesOfX {
			// check east
			if indexOfX + 3 < len(line) {
				var sb strings.Builder
				sb.WriteByte(line[indexOfX])
				sb.WriteByte(line[indexOfX+1])
				sb.WriteByte(line[indexOfX+2])
				sb.WriteByte(line[indexOfX+3])
				if sb.String() == "XMAS" {
					total++
				}
			}
			// check west
			if indexOfX - 3 >= 0 {
				var sb strings.Builder
				sb.WriteByte(line[indexOfX])
				sb.WriteByte(line[indexOfX-1])
				sb.WriteByte(line[indexOfX-2])
				sb.WriteByte(line[indexOfX-3])
				if sb.String() == "XMAS" {
					total++
				}
			}
			// check south
			if i + 3 < len(lines) {
				var sb strings.Builder
				sb.WriteByte(line[indexOfX])
				sb.WriteByte(lines[i + 1][indexOfX])
				sb.WriteByte(lines[i + 2][indexOfX])
				sb.WriteByte(lines[i + 3][indexOfX])
				if sb.String() == "XMAS" {
					total++
				}
			}
			// check north
			if i - 3 >= 0 {
				var sb strings.Builder
				sb.WriteByte(line[indexOfX])
				sb.WriteByte(lines[i - 1][indexOfX])
				sb.WriteByte(lines[i - 2][indexOfX])
				sb.WriteByte(lines[i - 3][indexOfX])
				if sb.String() == "XMAS" {
					total++
				}
			}

			// DIAGONAL
			// Check northeast
			if i - 3 >= 0  && indexOfX + 3 < len(line) {
				var sb strings.Builder
				sb.WriteByte(line[indexOfX])
				sb.WriteByte(lines[i-1][indexOfX+1])
				sb.WriteByte(lines[i-2][indexOfX+2])
				sb.WriteByte(lines[i-3][indexOfX+3])
				if sb.String() == "XMAS" {
					total++
				}
			}

			// check southeast
			if i + 3 < len(lines) && indexOfX + 3 < len(line) {
				var sb strings.Builder
				sb.WriteByte(line[indexOfX])
				sb.WriteByte(lines[i+1][indexOfX+1])
				sb.WriteByte(lines[i+2][indexOfX+2])
				sb.WriteByte(lines[i+3][indexOfX+3])
				if sb.String() == "XMAS" {
					total++
				}
			}

			// check northwest
			if i - 3 >= 0 && indexOfX - 3 >= 0 {
				var sb strings.Builder
				sb.WriteByte(line[indexOfX])
				sb.WriteByte(lines[i-1][indexOfX-1])
				sb.WriteByte(lines[i-2][indexOfX-2])
				sb.WriteByte(lines[i-3][indexOfX-3])
				if sb.String() == "XMAS" {
					total++
				}
			}

			// check southwest
			if i + 3 < len(lines) && indexOfX - 3 >= 0 {
				var sb strings.Builder
				sb.WriteByte(line[indexOfX])
				sb.WriteByte(lines[i + 1][indexOfX-1])
				sb.WriteByte(lines[i + 2][indexOfX-2])
				sb.WriteByte(lines[i + 3][indexOfX-3])
				if sb.String() == "XMAS" {
					total++
				}
			}
		}

		
	}
	return total

}