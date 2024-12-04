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
func dayTwo() int {
	s := getData("1")

	allLocations := strings.FieldsFunc(s, func (r rune) bool {
		return r == ' ' || r == '\n'
	})

	var locationOneArr, locationTwoArr []int

	// split single slice into two seperate slices representing each location
	for i, location := range allLocations {
		locationId := must(strconv.Atoi(location))
        if i % 2 == 0 {
            locationOneArr = append(locationOneArr, locationId)
			continue
        }
		locationTwoArr = append(locationTwoArr, locationId)
    }

	// Sorting is probably the heaviest part of this problem
	sort.Ints(locationOneArr)
	sort.Ints(locationTwoArr)

	// reduce both slices down to a total "distance" between all the ids
	var total int
	for i, locationOne := range locationOneArr {
		locationTwo := locationTwoArr[i]
		total += int(math.Abs(float64(locationTwo - locationOne)))
    }

	return total
}