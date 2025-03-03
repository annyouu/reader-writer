package main

import (
	"io"
	"os"
	"strings"
)

func main() {
	r := &io.LimitedReader{
		R: strings.NewReader("Hello, 世界"),
		N: 5, //最大5バイト読み込む
	}

	// 最大5バイト読み込んで出力に書き込む
	io.Copy(os.Stdout, r)
}