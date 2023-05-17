package Book

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	key      = "ZUreQN0E"                                                                                            // 解密密钥(DES加密取前八位)
	keyParam = "NpkTYvpvhJjEog8Y051gQDHmReY54z5t3F0zSd9QEFuxWGqfC8g8Y4GPuabq0KPdxArlji4dSnnHCARHnkqYBLu7iIw55ibTo18" //param参数后缀密钥
)

var (
	requestId        string
	paramInfo        string
	encryptParamInfo string
	timestamp        int64
)

func getProgressBar(current, total int) string {
	width := 50
	progress := float64(current) / float64(total)
	completeWidth := int(progress * float64(width))
	remainingWidth := width - completeWidth
	return strings.Repeat("█", completeWidth) + strings.Repeat(" ", remainingWidth)
}

func Process(bookId string) {
	notes, bookName := GetBook(bookId)
	fileName := bookName + ".txt"
	chapterIDs := ChapterList(bookId)
	file, err := os.OpenFile(fileName, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("无法打开文件:", err)
		return
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	_, err = writer.WriteString(notes + "\n")
	totalChapters := len(chapterIDs)
	fmt.Println("开始下载章节...")
	for i, chapterID := range chapterIDs {
		_, content := GetContent(bookId, chapterID)
		_, err = writer.WriteString(content)
		if err != nil {
			fmt.Println("写入文件失败:", err)
			return
		}
		fmt.Printf("\r正在下载章节 %d/%d [%s]", i+1, totalChapters, getProgressBar(i+1, totalChapters))
	}
	err = writer.Flush()
	if err != nil {
		fmt.Println("刷新缓冲区失败:", err)
		return
	}
	fmt.Println("\n全部下载任务已完成!")
}
