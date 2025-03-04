package main

import (
	"fmt"
	"regexp"
)

func main() {
	// 正規表現をコンパイル
	re, err := regexp.Compile(`\d+`)
	if err != nil {
		// 正規表現にエラーがあれば処理
		fmt.Println("エラー:", err)
		return
	}

	// バイトスライスでFIndIndexを使う
	fmt.Println("FindIndex:")
	// 最初のマッチ部分のインデックスを返す
	fmt.Println(re.FindIndex([]byte("1986年01月12日")))

	// バイトスライスでFindAllIndexを使う
	fmt.Println("\nFindAllIndex:")
	// 全てのマッチ部分のインデックスを返す
	fmt.Println(re.FindAllIndex([]byte("1986年01月12日"), -1))

	// 文字列でFindStringIndexを使う
	fmt.Println("\nFindStringIndex:")
	// 最初のマッチ部分のインデックスを返す
	fmt.Println(re.FindStringIndex("1986年01月12日"))

	fmt.Println(re.FindAllStringIndex("1986年01月12日", -1))
}