package main

import (
	"errors"
	"fmt"
	"net/http"
)

var errReqeustFailed = errors.New("Request failed")

func main() {
	urls := []string{"https://www.naver.com", "https://google.com", "https://academy.nomadcoders.co/"}

	for _, url := range urls {
		hitWeb(url)
	}
}

func hitWeb(url string) error {
	fmt.Println("checking: ", url)
	resp, err := http.Get(url)
	if err == nil || resp.StatusCode >= 400 {
		return errReqeustFailed
	}
	return nil
}
