package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)


func GetEnv() string {
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

func GetData(day string) string {

	req, err := http.NewRequest("GET",fmt.Sprintf("https://adventofcode.com/2024/day/%s/input", day), nil)
	if (err != nil) {
		panic(err)
	}

	
	req.AddCookie(&http.Cookie{Name: "session", Value: GetEnv()})

	res, err := http.DefaultClient.Do(req)

	if (err != nil) {
		panic(err)
	}

	body, err := io.ReadAll(res.Body)

	if (err != nil) {
		panic(err)
	}

	res.Body.Close()

	return string(body)
}