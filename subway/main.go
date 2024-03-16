package main

import (
	"fmt"
	"log"
	"sync"
)

//type ArrivalInfo struct {
//	ErrorMessage struct {
//		Status  int    `json:"status"`
//		Code    string `json:"code"`
//		Message string `json:"message"`
//	} `json:"errorMessage"`
//	RealtimeArrivalList []struct {
//		SubwayID    string `json:"subwayId"`
//		SubwaysNm   string `json:"subwaysNm"`
//		StatnNm     string `json:"statnNm"`
//		ArvlMsg2    string `json:"arvlMsg2"`
//		ArvlMsg3    string `json:"arvlMsg3"`
//		TrainLineNm string `json:"trainLineNm"`
//		StatnFid    string `json:"statnFid"`
//		StatnTid    string `json:"statnTid"`
//		StatnId     string `json:"statnId"`
//		UpdnLine    string `json:"updnLine"`
//	} `json:"realtimeArrivalList"`
//}

//func fetchArrivalInfo(url string, wg *sync.WaitGroup, ch chan<- *ArrivalInfo, errCh chan<- error) {
//	defer wg.Done()
//
//	resp, err := http.Get(url)
//	if err != nil {
//		errCh <- err
//		return
//	}
//	defer resp.Body.Close()
//
//	var arrivalInfo ArrivalInfo
//	if err := json.NewDecoder(resp.Body).Decode(&arrivalInfo); err != nil {
//		errCh <- err
//		return
//	}
//
//	ch <- &arrivalInfo
//}

func main() {
	baseUrl := "http://swopenapi.seoul.go.kr/api/subway/764a4a554974626e35305252644367/json/realtimeStationArrival/0/8/"
	lineTwoStations := []string{"역삼", "강남", "교대", "서초", "방배", "사당"}
	var urls []string
	for _, station := range lineTwoStations {
		urls = append(urls, fmt.Sprint(baseUrl, station))
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
			if arrivalInfo.ErrorMessage.Status != 200 {
				return
			}
			// Display the arrival information
			for _, arrival := range arrivalInfo.RealtimeArrivalList {
				fmt.Printf("Subway ID: %s, Line: %s, Station: %s, Arrival Message: %s, , Arrival Message3: %s, TrainLineNm: %s, StatnFid: %s, StatnTid: %s, StatnId: %s, UpdnLine: %s\n",
					arrival.SubwayID,
					arrival.SubwaysNm,
					arrival.StatnNm,
					arrival.ArvlMsg2,
					arrival.ArvlMsg3,
					arrival.TrainLineNm,
					arrival.StatnFid,
					arrival.StatnTid,
					arrival.StatnId,
					arrival.UpdnLine)
			}
			fmt.Println("--------------------------------------")
		case err := <-errCh:
			log.Println("Error occurred:", err)
		}
	}
}
