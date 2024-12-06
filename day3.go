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
	matches := re.FindAllString(s, -1)

	for _, item := range matches {
		pair := strings.Split((item[4:(len(item) - 1)]), ",")
		total += must(strconv.Atoi(pair[0])) * must(strconv.Atoi(pair[1]))
	}
	return total
}

func day3Part2() int {
	s := getData("3")
	total := 0
	skip := false

	// For part two, I need to add do/dont matching. Then when iterating over the list, just update do/don't
	re := regexp.MustCompile(`mul\([0-9]{1,3},[0-9]{1,3}\)|do\(\)|don't\(\)`)
	matches := re.FindAllString(s, -1)

	for _, match := range matches {
		if match == "do()" {
			skip = false
			continue
		}
		if match == "don't()" { skip = true }
		if skip { continue }

		pair := strings.Split((match[4:(len(match) - 1)]), ",")
		total += must(strconv.Atoi(pair[0])) * must(strconv.Atoi(pair[1]))
	}
	return total
}