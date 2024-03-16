package main

// ArrivalInfo struct definition
type ArrivalInfo struct {
	ErrorMessage struct {
		Status  int    `json:"status"`
		Code    string `json:"code"`
		Message string `json:"message"`
	} `json:"errorMessage"`
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
		UpdnLine    string `json:"updnLine"`
	} `json:"realtimeArrivalList"`
}
