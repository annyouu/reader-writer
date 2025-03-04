package main

import (
	"fmt"
	"regexp"
)

func main() {
	re, err := regexp.Compile(`(?P<Y>\d+)年(?P<M>\d+)月(?P<D>\d+)日`)
	if err != nil {
		fmt.Println("エラー:", err)
		return
	}

	// 対象となる文字列
	content := "1986年01月12日\n2020年03月24日"

	// 置換後のテンプレート
	template := "$Y/$M/$D\n"

	// 置換処理
	var result []byte
	for _, submatches := range re.FindAllStringSubmatchIndex(content, -1) {
		result = re.ExpandString(result, template, content, submatches)
	}
	// 結果を表示
	fmt.Printf("%q", result)
}