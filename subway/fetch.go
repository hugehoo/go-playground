package main

import (
	"encoding/json"
	"net/http"
	"sync"
)

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
