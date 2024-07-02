package main

import (
	"errors"
	"net/http"
)

type requestResult struct {
	url    string
	status string
}

var errRequestFailed = errors.New("Request failed")

func main() {
	//var results = map[string]string{}
	var results = make(map[string]string) // make is function that makes map
	c := make(chan requestResult)

	urls := []string{
		"https://tech.kakaopay.com/page/1/",
		"https://tech.kakaobank.com/page/1/",
		"https://toss.tech/tech",
		"https://d2.naver.com/helloworld",
		"https://medium.com/daangn/development/home"}

	for _, url := range urls {
		go hitWeb(url, c)
	}
	for i := 0; i < len(urls); i++ {
		result := <-c
		results[result.url] = result.status
	}

	for url, result := range results {
		println(url, result)
	}
}

func hitWeb(url string, c chan<- requestResult) { // this channel is send only, can't receive anything
	resp, err := http.Get(url)
	status := "OK"
	if err != nil || resp.StatusCode >= 400 {
		status = "FAILED"
	}
	c <- requestResult{url: url, status: status}
}

// 메인 함수가 끝나면 고루틴도 종료된다.
// channel 은 고루틴과 메인함수 사이에 정보를 전달하기 위한 방법이다. 혹은 고루틴 간 커뮤니케이션도 지원한다.
// 채널은 파이프라인 같은걸로, 메시지르 보낼 수도 받을 수도 있따.
