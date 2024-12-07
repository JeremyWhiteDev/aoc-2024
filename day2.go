package main

import (
	"math"
	"reflect"
	"strconv"
	"strings"
)

// Given two lists of numbers, order the lists from lowest to highest,
// and compare the difference of the numbers at matching indexes,
// collecting the differences into a total "difference" between the two lists
func dayTwoPartOne() int {
	s := getData("2")

	allReports := strings.FieldsFunc(s, func(r rune) bool {
		return r == '\n'
	})

	safeReports := 0

	// report is safe if
	// all increasing
	// all decreasing
	// adjacent levels differ by minimum 1 and maximum 3
	reportIsNotSafe := func(currentLevel, nextLevel int, reportShouldIncrease bool) bool {
		nextLevelDif := math.Abs(float64(currentLevel - nextLevel))
		return ((reportShouldIncrease && currentLevel > nextLevel) ||
			(!reportShouldIncrease && currentLevel < nextLevel) ||
			(nextLevelDif < 1 || nextLevelDif > 3))
	}

	for _, report := range allReports {
		reportArr := strings.Fields(report)
		reportIsSafe := true
		reportShouldIncrease := false

		for i, currentLevel := range reportArr {
			// we're done with this report
			if len(reportArr)-1 == i {
				continue
			}

			currentLevel := must(strconv.Atoi(currentLevel))
			nextLevel := must(strconv.Atoi(reportArr[i+1]))
			// set whether we expect the report to increase based on the first
			// two levels
			if i == 0 {
				reportShouldIncrease = currentLevel < nextLevel
			}

			if reportIsNotSafe(currentLevel, nextLevel, reportShouldIncrease) {
				reportIsSafe = false
				break
			}
		}

		if reportIsSafe {
			safeReports++
		}
	}
	return safeReports
}

// incorrect currently. I couldn't get this to work as I'm over engineering this problem.
// I looked up a solution and have comments on that solution's implementation.
// Solution is below
func dayTwoPartTwo() int {
	s := getData("2")

	allReports := strings.FieldsFunc(s, func(r rune) bool {
		return r == '\n'
	})

	safeReports := 0

	for _, report := range allReports {
		reportArr := strings.Fields(report)

		if reportIsSafe(reportArr, false) {
			safeReports++
		}
	}
	return safeReports
}

func reportIsSafe(reportArr []string, skipRecursion bool) bool {
	reportShouldIncrease := false
	numOfLevelIncreases := 0
	numOfLevelDecreases := 0
	var unsafeLevels []int

	// Determine whether a report is expected to increase or decrease.
	// 1 1 2
	// 20 22 19 18 17
	// 20 21 19 23 24
	for i, currentLevel := range reportArr {
		if i == 0 {
			continue
		}
		currentLevel := must(strconv.Atoi(currentLevel))
		prevLevel := must(strconv.Atoi(reportArr[i-1]))
		if currentLevel > prevLevel {
			numOfLevelIncreases++
		}
		if currentLevel < prevLevel {
			numOfLevelDecreases++
		}
	}

	reportShouldIncrease = numOfLevelIncreases > numOfLevelDecreases

	// iterate over levels and find those levels which are unsafe
	for i, currentLevel := range reportArr {
		if i == 0 {
			continue
		}
		currentLevel := must(strconv.Atoi(currentLevel))
		prevLevel := must(strconv.Atoi(reportArr[i-1]))

		levelIsNotSafe, _ := levelIsNotSafe(currentLevel, prevLevel, reportShouldIncrease)

		if levelIsNotSafe {
			unsafeLevels = append(unsafeLevels, i)
		}
	}

	// TODO unsafe reason?
	if len(unsafeLevels) == 0 {
		return true
	}

	// brute force retry

	if skipRecursion {
		return false
	}

	// check if removing any of the levels makes the report safe
	for _, unsafeLevelIndex := range unsafeLevels {
		newReport := RemoveIndex(reportArr, unsafeLevelIndex)
		if reportIsSafe(newReport, true) {
			return true
		}
	}
	return false
}

// level is safe if
// The level increases/decreases as expected
// adjacent levels differ by minimum 1 and maximum 3
func levelIsNotSafe(currentLevel, prevLevel int, reportShouldIncrease bool) (bool, string) {
	prevLevelDiff := math.Abs(float64(currentLevel - prevLevel))

	if reportShouldIncrease && currentLevel < prevLevel {
		return true, "Levels decrease instead of increasing"
	}

	if !reportShouldIncrease && currentLevel > prevLevel {
		return true, "Levels increase instead of decrease"

	}

	if prevLevelDiff < 1 {
		return true, "Levels are the same"
	}

	if prevLevelDiff > 3 {
		return true, "Level difference greater than 3"
	}

	return false, ""
}

func RemoveIndex(s []string, index int) []string {
	ret := make([]string, 0)
	ret = append(ret, s[:index]...)
	return append(ret, s[index+1:]...)
}

// This was taken from youtuber Jonathan Paulson and rewritten in Go.
// I added comments on the failings of my previous enderstanding
func dayTwoPartTwoSolution() int {
	s := getData("2")
	safeReports := 0
	allReports := strings.FieldsFunc(s, func(r rune) bool { return r == '\n' })

	for _, report := range allReports {
		reportArr := Map(strings.Fields(report), func(item string) int { return must(strconv.Atoi(item)) })
		if reportIsSafe2(reportArr) {
			safeReports++
			continue
		}

		for i := range reportArr {
			// Create a modified report slice excluding the element at index i
			modReport := append(reportArr[:i:i], reportArr[i+1:]...)
			if reportIsSafe2(modReport) {
				safeReports++
				break
			}
		}
	}
	return safeReports
}

func reportIsSafe2(reportArr []int) bool {
	incOrDec := (reflect.DeepEqual(reportArr, SortToCopy(reportArr, true)) ||
		reflect.DeepEqual(reportArr, SortToCopy(reportArr, false)))
	allLevelsSafe := true

	for j, level := range reportArr {
		if j == len(reportArr)-1 {
			continue
		}
		diff := int(math.Abs(float64(level - reportArr[j+1])))
		if diff < 1 || diff > 3 {
			allLevelsSafe = false
			break
		}
	}
	return incOrDec && allLevelsSafe
}
