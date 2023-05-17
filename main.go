package main

import (
	"CiYuanJi/Book"
	"fmt"
)

func main() {
	var bookId string
	Book.GetUserBookRackList()
	fmt.Print("请输入bookId：")
	_, err := fmt.Scanln(&bookId)
	if err != nil {
		fmt.Println("输入发生错误")
	}
	Book.Process(bookId)
}
