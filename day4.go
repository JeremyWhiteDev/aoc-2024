package main

import (
	"strings"
)

func day4Part1() int {
	s := getData("4")
	total := 0
	lines := strings.FieldsFunc(s, func(r rune) bool { return r == '\n' })
	// iterate over each line
	for i, line := range lines {
		var indexesOfX []int
		for j, c := range line {
			if c == 'X' {
				indexesOfX = append(indexesOfX, j)
			}
		}
		// day4 brute force

		// todo, refactor to cursor that "moves"
		// provide cursor with x/y directions. if cursor is given direction it can't go (is at it's limit), don't move it.

		// for each index of x, do the search for backwards, forwards, up, down, and 4 diagnoal directions (safely)
		for _, indexOfX := range indexesOfX {
			// check east
			if indexOfX+3 < len(line) {
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
			if indexOfX-3 >= 0 {
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
			if i+3 < len(lines) {
				var sb strings.Builder
				sb.WriteByte(line[indexOfX])
				sb.WriteByte(lines[i+1][indexOfX])
				sb.WriteByte(lines[i+2][indexOfX])
				sb.WriteByte(lines[i+3][indexOfX])
				if sb.String() == "XMAS" {
					total++
				}
			}
			// check north
			if i-3 >= 0 {
				var sb strings.Builder
				sb.WriteByte(line[indexOfX])
				sb.WriteByte(lines[i-1][indexOfX])
				sb.WriteByte(lines[i-2][indexOfX])
				sb.WriteByte(lines[i-3][indexOfX])
				if sb.String() == "XMAS" {
					total++
				}
			}

			// DIAGONAL
			// Check northeast
			if i-3 >= 0 && indexOfX+3 < len(line) {
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
			if i+3 < len(lines) && indexOfX+3 < len(line) {
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
			if i-3 >= 0 && indexOfX-3 >= 0 {
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
			if i+3 < len(lines) && indexOfX-3 >= 0 {
				var sb strings.Builder
				sb.WriteByte(line[indexOfX])
				sb.WriteByte(lines[i+1][indexOfX-1])
				sb.WriteByte(lines[i+2][indexOfX-2])
				sb.WriteByte(lines[i+3][indexOfX-3])
				if sb.String() == "XMAS" {
					total++
				}
			}
		}

	}
	return total

}

type cursor struct {
	x     int
	y     int
	lines []string
}

// get value relative to cursor. If attempted value is out of range, then return the value at the cursor
func (c *cursor) getRelativeToCursor(x, y int) byte {
	// get the value at the current cursor
	if x+c.x >= len(c.lines[c.y]) || x+c.x < 0 {
		return c.get()
	}
	if y+c.y >= len(c.lines) || y+c.y < 0 {
		return c.get()
	}
	return c.lines[c.y+y][c.x+x]
}

// set the cursor's position
func (c *cursor) setPosition(x, y int) {
	c.x = x
	c.y = y
}

// get the value at the cursor's current coordinates
func (c *cursor) get() byte {
	return c.lines[c.y][c.x]
}

// This is incorrect??
func day4Part2() int {
	s := getData("4")
	total := 0
	lines := strings.FieldsFunc(s, func(r rune) bool { return r == '\n' })
	var coordOfA [][]int

	for y, line := range lines {
		for x, c := range line {
			if c == 'A' {
				coordOfA = append(coordOfA, []int{x, y})
			}
		}
	}

	cursor := cursor{lines: lines}

	for _, coord := range coordOfA {
		grc := cursor.getRelativeToCursor
		cursor.setPosition(coord[0], coord[1])
		var tlbr, trbl bool

		// top left/bottom right
		if (grc(-1, -1) == 'M' && grc(1, 1) == 'S') || (grc(-1, -1) == 'S' && grc(1, 1) == 'M') {
			tlbr = true
		}
		// top right/bottem left
		if (grc(1, -1) == 'M' && grc(-1, 1) == 'S') || (grc(1, -1) == 'S' && grc(-1, 1) == 'M') {
			trbl = true
		}
		if tlbr && trbl {
			total++
		}
	}

	return total
}
