package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	re, err := regexp.Compile(`(^|_)[a-zA-Z]`)
	if err != nil {
		// エラー処理
		fmt.Println("エラー:", err)
		return
	}

	// 置換対象となる文字列
	s := "hello_world"

	// 置換処理:あんだースコアか最初の文字を大文字にする
	result := re.ReplaceAllStringFunc(s, func(s string) string {
		// _を消して大文字に変換
		return strings.ToUpper(strings.TrimLeft(s, "_"))
	})
	fmt.Println(result)
}