package main

import (
	"fmt"
	"regexp"
)

var validID = regexp.MustCompile(`^[a-z]+\[[0-9]+\]$`)

func main() {
	fmt.Println(validID.MatchString("adam[23]"))

	validID2, err := regexp.Compile(`^[a-z]+\[[0-9]+\]$`)
	if err != nil {
		// エラー処理
		fmt.Println("正規表現失敗:", err)
		return
	}
	fmt.Println(validID2.MatchString("adam[23]"))
}