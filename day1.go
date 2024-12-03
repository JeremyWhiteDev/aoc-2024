package main

import (
	"math"
	"sort"
	"strconv"
	"strings"
)

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

	sort.Ints(locationOneArr)
	sort.Ints(locationTwoArr)

	var distances []int
	for i, locationOne := range locationOneArr {
		locationTwo := locationTwoArr[i]
		distances = append(distances, int(math.Abs(float64(locationTwo - locationOne))))
    }

	var total int
	for _, distance := range distances {
		total += distance
    }

	return total
}

func Split(r rune) bool {
	return r == ' ' || r == '\n'
}
