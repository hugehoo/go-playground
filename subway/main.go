package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
)

type ArrivalInfo struct {
	ErrorMessage struct {
		Status  int    `json:"status"`
		Code    string `json:"code"`
		Message string `json:"message"`
	}
	RealtimeArrivalList []struct {
		SubwayID    string `json:"subwayId"`
		SubwaysNm   string `json:"subwaysNm"`
		StatnNm     string `json:"statnNm"`
		ArvlMsg2    string `json:"arvlMsg2"`
		ArvlMsg3    string `json:"arvlMsg3"`
		TrainLineNm string `json:"trainLineNm"`
		StatnFid    string `json:"statnFid"`
		StatnTid    string `json:"statnTid"`
		StatnId     string `json:"statnId"`
	} `json:"realtimeArrivalList"`
}

func fetchArrivalInfo(url string, wg *sync.WaitGroup, ch chan<- *ArrivalInfo, errCh chan<- error) {
	defer wg.Done()

	resp, err := http.Get(url)
	if err != nil {
		errCh <- err
		return
	}
	defer resp.Body.Close()

	var arrivalInfo ArrivalInfo
	if err := json.NewDecoder(resp.Body).Decode(&arrivalInfo); err != nil {
		errCh <- err
		return
	}

	ch <- &arrivalInfo
}

func main() {
	urls := []string{
		"http://swopenapi.seoul.go.kr/api/subway/764a4a554974626e35305252644367/json/realtimeStationArrival/0/6/역삼",
		"http://swopenapi.seoul.go.kr/api/subway/764a4a554974626e35305252644367/json/realtimeStationArrival/0/6/강남",
		"http://swopenapi.seoul.go.kr/api/subway/764a4a554974626e35305252644367/json/realtimeStationArrival/0/6/교대",
	}

	var wg sync.WaitGroup
	ch := make(chan *ArrivalInfo)
	errCh := make(chan error)

	for _, url := range urls {
		wg.Add(1)
		go fetchArrivalInfo(url, &wg, ch, errCh)
	}

	go func() {
		wg.Wait()
		close(ch)
		close(errCh)
	}()

	for {
		select {
		case arrivalInfo, ok := <-ch:
			if !ok {
				return
			}
			// Display the arrival information
			for _, arrival := range arrivalInfo.RealtimeArrivalList {
				fmt.Printf("Subway ID: %s, Line: %s, Station: %s, Arrival Message: %s\n", arrival.SubwayID, arrival.SubwaysNm, arrival.StatnNm, arrival.ArvlMsg2)
			}
			fmt.Println("--------------------------------------")
		case err := <-errCh:
			log.Println("Error occurred:", err)
		}
	}
}
