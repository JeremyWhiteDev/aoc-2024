package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"slices"
	"strings"
)


func getEnv() string {
	loadEnv()
	return os.Getenv("SESSION")
}

func loadEnv() {
	bytes := must(os.ReadFile(".env"))
	os.Setenv("SESSION", strings.Split(string(bytes), "=")[1])
}

func getData(day string) string {
	req := must(http.NewRequest("GET",fmt.Sprintf("https://adventofcode.com/2024/day/%s/input", day), nil))
	req.AddCookie(&http.Cookie{Name: "session", Value: getEnv()})

	res := must(http.DefaultClient.Do(req))
	body  := must(io.ReadAll(res.Body))
	res.Body.Close()

	return string(body)
}

func must[T any](obj T, err error) T {
    if err != nil {
        panic(err)
    }
    return obj
}

func Map[T, V any](slice []T, fn func(T) V) []V {
    result := make([]V, len(slice))
    for i, t := range slice {
        result[i] = fn(t)
    }
    return result
}

func SortToCopy(slice []int, asc bool) []int {
	sliceCopy := make([]int, len(slice))
	copy(sliceCopy, slice)
	if asc {
		slices.Sort(sliceCopy)
		return sliceCopy
	}
	
	slices.Sort(sliceCopy)
	slices.Reverse(sliceCopy)
	return sliceCopy
}