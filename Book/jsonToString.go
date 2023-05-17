package Book

import (
	"encoding/json"
	"fmt"
)

type getInfo struct {
	SortType  string `json:"sortType"`
	PageNo    string `json:"pageNo"`
	PageSize  string `json:"pageSize"`
	BookID    string `json:"bookId"`
	Timestamp int64  `json:"timestamp"`
}

func GetInfoJson(bookId string, timestamp int64) string {
	jsonStruct := getInfo{
		SortType:  "1",
		PageNo:    "1",
		PageSize:  "9999",
		BookID:    bookId,
		Timestamp: timestamp,
	}
	jsonData, err := json.Marshal(jsonStruct)
	if err != nil {
		fmt.Println("JSON encoding error:", err)
		return ""
	}
	jsonString := string(jsonData)
	return jsonString
}

type getContent struct {
	BookID    string `json:"bookId"`
	ChapterId string `json:"chapterId"`
	Timestamp int64  `json:"timestamp"`
}

func GetContentJson(bookId, chapterId string, timestamp int64) string {
	jsonStruct := getContent{
		BookID:    bookId,
		ChapterId: chapterId,
		Timestamp: timestamp,
	}
	jsonData, err := json.Marshal(jsonStruct)
	if err != nil {
		fmt.Println("JSON encoding error:", err)
		return ""
	}
	jsonString := string(jsonData)
	return jsonString
}

type GetBookRackList struct {
	RankType  string `json:"rankType"`
	PageNo    string `json:"pageNo"`
	PageSize  string `json:"pageSize"`
	Timestamp int64  `json:"timestamp"`
}

func GetBookRackListJson(timestamp int64) string {
	jsonStruct := GetBookRackList{
		RankType:  "1",
		PageNo:    "1",
		PageSize:  "9999",
		Timestamp: timestamp,
	}
	jsonData, err := json.Marshal(jsonStruct)
	if err != nil {
		fmt.Println("JSON encoding error:", err)
		return ""
	}
	jsonString := string(jsonData)
	return jsonString
}
