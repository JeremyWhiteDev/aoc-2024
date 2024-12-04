package main

import (
	"math"
	"strconv"
	"strings"
)

// Given two lists of numbers, order the lists from lowest to highest,
// and compare the difference of the numbers at matching indexes,
// collecting the differences into a total "difference" between the two lists
func dayTwoPartOne() int {
	s := getData("2")

	allReports := strings.FieldsFunc(s, func (r rune) bool {
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
			(nextLevelDif < 1 ||nextLevelDif > 3))
	}

	for _, report := range allReports {
		reportArr := strings.Fields(report)
		reportIsSafe := true
		reportShouldIncrease := false
		

		for i, currentLevel := range reportArr {
			// we're done with this report
			if len(reportArr) - 1 == i {
				continue
			}
			
			currentLevel := must(strconv.Atoi(currentLevel))
			nextLevel := must(strconv.Atoi(reportArr[i +1]))
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