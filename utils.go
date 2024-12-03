package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)


func getEnv() string {
	loadEnv()
	return os.Getenv("SESSION")
}

func loadEnv() {
	bytes, err := os.ReadFile(".env")
	if (err != nil) {
		panic(err)
	}


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