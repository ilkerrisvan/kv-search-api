package model

import "go.mongodb.org/mongo-driver/bson"

type KeyValueDbData struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type RecordsRequest struct {
	StartDate string  `json:"startDate"`
	EndDate   string  `json:"endDate"`
	MinCount  float64 `json:"minCount"`
	MaxCount  float64 `json:"maxCount"`
}

type RecordsResponse struct {
	Code    int      `json:"code"`
	Msg     string   `json:"msg"`
	Records []bson.M `json:"records"`
}
