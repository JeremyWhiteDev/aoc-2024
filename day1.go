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
func dayOnePartOne() int {
	s := getData("1")

	allLocations := strings.FieldsFunc(s, Split)

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

// Given two lists of numbers, calculate a "similarity score"
// that compares the first list to the second list.
// 
// The calculation should be based on the total sum of locationIds in the first
// multiplied by the frequency of those locationIds in the second list.
func dayOnePartTwo() int {
	s := getData("1")

	allLocations := strings.FieldsFunc(s, Split)

	// create two maps, with locationId:frequency
	locationOneMap := make(map[int]int)
	locationTwoMap := make(map[int]int)

	for i, location := range allLocations {
		locationId := must(strconv.Atoi(location))
        if i % 2 == 0 {
			locationOneMap[locationId] += 1
			continue
        }
		locationTwoMap[locationId] += 1
    }

	var similarityScore int
	// iterate over the first location map, calculating the similarity score based
	// based on whether the locationId occurs in the second map
	for locationId := range locationOneMap {
		similarityScore += (locationId * locationTwoMap[locationId])
	}

	return similarityScore
}

func Split(r rune) bool {
	return r == ' ' || r == '\n'
}
