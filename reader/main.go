package main

import (
	"fmt"
	"io"
	"os"
)

// ファイルの内容を読み取り、標準出力に書き込む
func main() {
	// 読み込むファイルを作成
	fileName := "example.txt"
	content := "Hello, Go Reader and Writer!"
	err := os.WriteFile(fileName, []byte(content), 0644)
	if err != nil {
		fmt.Println("ファイルの作成失敗:", err)
		return
	}

	// ファイルを開く
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("ファイルを開くのに失敗:", err)
		return
	}
	defer file.Close()

	// 読み込む用のバッファを作成
	buf := make([]byte, 1024)
	for {
		n, err := file.Read(buf) // Readメソッドを使用
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("読み込みエラー:", err)
			return
		}

		// 読み込んだデータを標準出力に書き込む
		os.Stdout.Write(buf[:n])   // Writeメソッドを使用
	}
}