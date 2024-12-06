package main

import (
	"regexp"
	"strconv"
	"strings"
)

func day3Part1() int {
	s := getData("3")
	total := 0

	re := regexp.MustCompile(`mul\([0-9]{1,3},[0-9]{1,3}\)`)
	things := re.FindAllString(s, -1)

	for _, item := range things {
		pair := strings.Split((item[4:(len(item) - 1)]), ",")
		int1 := must(strconv.Atoi(pair[0]))
		int2 := must(strconv.Atoi(pair[1]))
		total += int1 * int2
	}
	return total
}

func day3Part2() int {
	s := getData("3")
	total := 0
	skip := false

	// For part two, I need to add do/dont matching. Then when iterating over the list, just update do/don't
	re := regexp.MustCompile(`mul\([0-9]{1,3},[0-9]{1,3}\)|do\(\)|don't\(\)`)
	things := re.FindAllString(s, -1)

	for _, item := range things {
		if item == "do()" {
			skip = false
			continue
		}
		if item == "don't()" { skip = true }

		if skip { continue }
		pair := strings.Split((item[4:(len(item) - 1)]), ",")

		int1 := must(strconv.Atoi(pair[0]))
		int2 := must(strconv.Atoi(pair[1]))
		total += int1 * int2
	}
	return total
}