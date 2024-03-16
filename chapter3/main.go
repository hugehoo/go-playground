package main

import (
	"errors"
	"fmt"
	"net/http"
)

var errRequestFailed = errors.New("Request failed")

func main() {
	//var results = map[string]string{}
	var results = make(map[string]string)

	urls := []string{"https://www.naver.com", "https://google.com", "https://academy.nomadcoders.co/"}
	results["gello"] = "hello"

	for _, url := range urls {
		err := hitWeb(url)
	}
}

func hitWeb(url string) error {
	fmt.Println("checking: ", url)
	resp, err := http.Get(url)
	if err == nil || resp.StatusCode >= 400 {
		return errRequestFailed
	}
	return nil
}
