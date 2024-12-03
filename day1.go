package main

import (
	"fmt"
	"strings"
)

func main() {

	s := GetData("1")

	allLocations := strings.FieldsFunc(s, Split)

	 var locationOneArr, locationTwoArr []string

	 for i, location := range allLocations {
        if i % 2 == 0 {
            locationOneArr = append(locationOneArr, location)
			continue
        }
		locationTwoArr = append(locationTwoArr, location)
	
    }

fmt.Print(locationOneArr)
fmt.Print(locationTwoArr)
}



func Split(r rune) bool {
	return r == ' ' || r == '\n'
}