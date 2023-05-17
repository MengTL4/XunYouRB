package Book

import (
	"CiYuanJi/Encrypt"
	"fmt"
	"github.com/google/uuid"
	"github.com/imroc/req/v3"
	"strconv"
	"strings"
	"time"
)

const (
	ChapterListUrl    = "/api/ciyuanji/client/chapter/getChapterListByBookId"
	BookDetailUrl     = "/api/ciyuanji/client/book/getBookDetail"
	ChapterContentUrl = "/api/ciyuanji/client/chapter/getChapterContent"
	UserBookRackList  = "/api/ciyuanji/client/bookrack/getUserBookRackList"
)

var (
	headers = map[string]string{
		"Channel":     "35",
		"User-Agent":  "Mozilla/5.0 (Linux; Android 11; Pixel 4 XL Build/RP1A.200720.009; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/92.0.4515.115 Mobile Safari/537.36",
		"Targetmodel": "SM-N9700",
		"Platform":    "1",
		"Deviceno":    "d0b7cef20c3c6b5f",
		"Version":     "3.3.2",
		"Token":       "",
	}
)

func GetID() string {
	return strings.ReplaceAll(uuid.New().String(), "-", "")
}

func CurrentTimeMillis() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

type Detail struct {
	Data struct {
		Book struct {
			BookID            int    `json:"bookId"`
			BookName          string `json:"bookName"`
			AuthorName        string `json:"authorName"`
			WordCount         int    `json:"wordCount"`
			ImgURL            string `json:"imgUrl"`
			Notes             string `json:"notes"`
			LatestUpdateTime  string `json:"latestUpdateTime"`
			LatestChapterId   int    `json:"latestChapterId"`
			LatestChapterName string `json:"latestChapterName"`
			ChapterCount      int    `json:"chapterCount"`
			TagList           []struct {
				TagName string `json:"tagName"`
			} `json:"tagList"`
			FirstChapterID int `json:"firstChapterId"`
		} `json:"book"`
	} `json:"data"`
}

func GetBook(bookId string) (string, string) {
	timestamp = CurrentTimeMillis()
	requestId = GetID()
	paramInfo = GetInfoJson(bookId, timestamp)
	encryptParamInfo = Encrypt.Param(paramInfo, key)
	infoSign := "param=" + encryptParamInfo + "&requestId=" + requestId + "&timestamp=" + strconv.FormatInt(timestamp, 10) + "&key=" + keyParam
	getBookInfoSign := Encrypt.Sign(infoSign)
	var response Detail
	client := req.C()
	request, _ := client.R().
		SetHeaders(headers).
		SetPathParams(map[string]string{
			"timestamp": strconv.FormatInt(timestamp, 10),
			"requestId": requestId,
			"sign":      getBookInfoSign,
			"param":     encryptParamInfo,
		}).
		SetSuccessResult(&response).
		Get("https://api.hwnovel.com" + BookDetailUrl + "?timestamp={timestamp}&requestId={requestId}&sign={sign}&param={param}")
	request.Close = true
	var tagNames []string
	for _, tag := range response.Data.Book.TagList {
		tagNames = append(tagNames, tag.TagName)
	}
	tagsString := strings.Join(tagNames, ",")
	str := "书名:" + response.Data.Book.BookName + "\n" +
		"作者:" + response.Data.Book.AuthorName + "\n" +
		"标签:" + tagsString + "\n" +
		"字数统计:" + strconv.Itoa(response.Data.Book.WordCount) + "\n" +
		"章节统计:" + strconv.Itoa(response.Data.Book.ChapterCount) + "\n" +
		"最后更新:" + response.Data.Book.LatestUpdateTime + "\n" +
		"简介:" + "\n" + response.Data.Book.Notes
	fmt.Println(str)
	return str, response.Data.Book.BookName
}

type AllChapterList struct {
	Data struct {
		BookChapter struct {
			ChapterList []struct {
				ChapterID   int    `json:"chapterId"`
				ChapterName string `json:"chapterName"`
			} `json:"chapterList"`
		} `json:"bookChapter"`
	} `json:"data"`
}

func ChapterList(bookId string) []string {
	timestamp = CurrentTimeMillis()
	requestId = GetID()
	paramInfo = GetInfoJson(bookId, timestamp)
	encryptParamInfo = Encrypt.Param(paramInfo, key)
	infoSign := "param=" + encryptParamInfo + "&requestId=" + requestId + "&timestamp=" + strconv.FormatInt(timestamp, 10) + "&key=" + keyParam
	getBookInfoSign := Encrypt.Sign(infoSign)
	var response AllChapterList
	client := req.C()
	request, _ := client.R().
		SetHeaders(headers).
		SetPathParams(map[string]string{
			"timestamp": strconv.FormatInt(timestamp, 10),
			"requestId": requestId,
			"sign":      getBookInfoSign,
			"param":     encryptParamInfo,
		}).
		SetSuccessResult(&response).
		Get("https://api.hwnovel.com" + ChapterListUrl + "?timestamp={timestamp}&requestId={requestId}&sign={sign}&param={param}")
	request.Close = true
	//列举目录
	//var chapterNames []string
	//for _, chapter := range response.Data.BookChapter.ChapterList {
	//	chapterNames = append(chapterNames, chapter.ChapterName)
	//}
	//chapterString := strings.Join(chapterNames, "\n")
	var chapterIDs []string
	for _, chapter := range response.Data.BookChapter.ChapterList {
		chapterIDs = append(chapterIDs, strconv.Itoa(chapter.ChapterID))
	}
	//fmt.Println(chapterString)
	//fmt.Println(chapterIDs)
	return chapterIDs
}

type Content struct {
	Data struct {
		Chapter struct {
			ChapterName string `json:"chapterName"`
			Content     string `json:"content"`
		} `json:"chapter"`
	} `json:"data"`
}

func GetContent(bookId, chapterId string) (string, string) {
	timestamp = CurrentTimeMillis()
	requestId = GetID()
	paramInfo = GetContentJson(bookId, chapterId, timestamp)
	encryptParamInfo = Encrypt.Param(paramInfo, key)
	infoSign := "param=" + encryptParamInfo + "&requestId=" + requestId + "&timestamp=" + strconv.FormatInt(timestamp, 10) + "&key=" + keyParam
	getBookInfoSign := Encrypt.Sign(infoSign)
	var response Content
	client := req.C()
	request, _ := client.R().
		SetHeaders(headers).
		SetPathParams(map[string]string{
			"timestamp": strconv.FormatInt(timestamp, 10),
			"requestId": requestId,
			"sign":      getBookInfoSign,
			"param":     encryptParamInfo,
		}).
		SetSuccessResult(&response).
		Get("https://api.hwnovel.com" + ChapterContentUrl + "?timestamp={timestamp}&requestId={requestId}&sign={sign}&param={param}")
	request.Close = true
	content := response.Data.Chapter.Content
	result := strings.Replace(content, "\n", "", -1)
	decryptContent := Encrypt.DecryptDESECB([]byte(result), []byte(key))
	finalStr := response.Data.Chapter.ChapterName + "\n" + decryptContent + "\n"
	return response.Data.Chapter.ChapterName, finalStr
}

type BookRackList struct {
	Data struct {
		BookRackList []struct {
			BookID   int    `json:"bookId"`
			BookName string `json:"bookName"`
		} `json:"bookRackList"`
	} `json:"data"`
}

func GetUserBookRackList() {
	timestamp = CurrentTimeMillis()
	requestId = GetID()
	paramInfo = GetBookRackListJson(timestamp)
	encryptParamInfo = Encrypt.Param(paramInfo, key)
	infoSign := "param=" + encryptParamInfo + "&requestId=" + requestId + "&timestamp=" + strconv.FormatInt(timestamp, 10) + "&key=" + keyParam
	getBookInfoSign := Encrypt.Sign(infoSign)
	var response BookRackList
	client := req.C()
	request, _ := client.R().
		SetHeaders(headers).
		SetPathParams(map[string]string{
			"timestamp": strconv.FormatInt(timestamp, 10),
			"requestId": requestId,
			"sign":      getBookInfoSign,
			"param":     encryptParamInfo,
		}).
		SetSuccessResult(&response).
		Get("https://api.hwnovel.com" + UserBookRackList + "?timestamp={timestamp}&requestId={requestId}&sign={sign}&param={param}")
	request.Close = true
	for _, book := range response.Data.BookRackList {
		fmt.Println("ID:", book.BookID, "\t", "Name:", book.BookName)
	}
}
