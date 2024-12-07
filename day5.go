package main

import (
	"strconv"
	"strings"
)

func day5Part1() int {
	s := getData("5")
	totalSafeInstructions := 0
	totalOfMiddleIndexes := 0
	splitRulesUpdates := strings.Split(s, "\n\n")
	updateInstructions := parseLines(splitRulesUpdates[1], ",")
	r := strings.FieldsFunc(splitRulesUpdates[0], func(r rune) bool { return r == '\n' })

	// create a "set" to allow contant time lookups of variables.
	rulesSet := make(map[string]struct{})
	for _, currRule := range r {
		rulesSet[currRule] = struct{}{}
	}

	// for each update instruction, we need to iterate over
	// every value, and check if that value combined with any
	// other values in the instruction
	for _, updateInstruction := range updateInstructions {
		updateIsValid := true
	outer:
		for i, currStep := range updateInstruction {
			// only iterate from the current value forward to save from repetitive checks
			for _, nextStep := range updateInstruction[i+1:] {
				if _, ok := rulesSet[nextStep+"|"+currStep]; ok {
					updateIsValid = false
					// fmt.Println("update is invalid", updateInstruction, "because of rule", nextStep+"|"+currStep)
					break outer
				}
			}
		}
		if updateIsValid {
			if value, ok := findMiddleValue(updateInstruction); ok {
				totalOfMiddleIndexes += value
			}
			totalSafeInstructions++
		}
	}

	return totalOfMiddleIndexes
}

func parseLines(s, del string) [][]string {
	var result [][]string
	rLines := strings.FieldsFunc(s, func(r rune) bool { return r == '\n' })
	for _, line := range rLines {
		result = append(result, strings.Split(line, del))
	}
	return result
}

func findMiddleValue(s []string) (int, bool) {
	if len(s)%2 == 0 {
		return 0, false
	}
	midIdx := len(s) / 2
	return must(strconv.Atoi(s[midIdx])), true
}
