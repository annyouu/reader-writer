package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// io.PipeでReaderとWriterを作成
	r, w := io.Pipe()

	// 非同期でWriterにデータを書き込む
	go func() {
		// w (pipeWriter)にデータを書き込む
		fmt.Fprint(w, "Hello, 世界\n")
		w.Close()  //書き込み完了後にClose
	}()

	// r (pipeReader)からデータを読み取り、os.Stdoutにコピー
	io.Copy(os.Stdout, r) // rから読み込んだデータをos.Stdoutに書き込む
}