package main

import (
	"bytes"
	"fmt"
	"io"
)

func main() {
	// buf1とbuf2はbytes.Buffer型の変数
	var buf1, buf2 bytes.Buffer

	// io.MultiWriterは、複数のデータを書き込む
	w := io.MultiWriter(&buf1, &buf2)

	// Fprintを使って、MultiWriterにデータを書き込む
	fmt.Fprint(w, "Hello, 世界")

	fmt.Println("buf1:", buf1.String())
	fmt.Println("buf2:", buf2.String())
}