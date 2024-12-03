package main

import (
	"math"
	"sort"
	"strconv"
	"strings"
)

// Given two lists of numbers, order the lists from lowest to highest,
// and compare the difference of the numbers at matching indexes,
// collecting the differences into a total "difference" between the two lists
func dayOne() int {
	s := getData("1")

	allLocations := strings.FieldsFunc(s, Split)

	// naive approach
	var locationOneArr, locationTwoArr []int

	for i, location := range allLocations {
		distance := must(strconv.Atoi(location))
        if i % 2 == 0 {
            locationOneArr = append(locationOneArr, distance)
			continue
        }
		locationTwoArr = append(locationTwoArr, distance)
    }

	// Sorting is probably the heaviest part of this problem
	sort.Ints(locationOneArr)
	sort.Ints(locationTwoArr)

	var total int
	for i, locationOne := range locationOneArr {
		locationTwo := locationTwoArr[i]
		total += int(math.Abs(float64(locationTwo - locationOne)))
    }

	return total
}

func Split(r rune) bool {
	return r == ' ' || r == '\n'
}
