package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	r := strings.NewReader("Hello, 世界")

	// 先頭から2バイト移動("He"をスキップ)
	r.Seek(2, io.SeekStart)
	io.CopyN(os.Stdout, r, 2)
	fmt.Println() // 改行

	// 現在の位置から-4バイト戻る
	r.Seek(-4, io.SeekCurrent)
	io.CopyN(os.Stdout, r, 7)
	fmt.Println() // 改行

	// 末尾から-6バイト戻る
	r.Seek(-6, io.SeekEnd)
	io.Copy(os.Stdout, r)
	fmt.Println() // 改行
}