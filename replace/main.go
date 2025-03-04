package main

import (
	"fmt"
	"regexp"
)

func main() {
	re, err := regexp.Compile(`(\d+)年(\d+)月(\d+)日`)
	if err != nil {
		fmt.Println("エラー:", err)
		return
	}

	date := "1986年01月12日"

	// 置換処理
	s := re.ReplaceAllString(date, "${2}/${3} ${1}")

	fmt.Println(s)
}