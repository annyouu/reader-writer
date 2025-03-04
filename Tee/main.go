package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	var buf bytes.Buffer // 出力用のバッファ
	r := strings.NewReader("Hello, 世界\n")

	// io.TeeReaderを使って、読み込むと同時にbufにも書き込む
	tee := io.TeeReader(r, &buf)

	// os.Stdout（標準出力）にteeの内容をコピー
	io.Copy(os.Stdout, tee)

	// bufにも同じ内容が書かれている
	fmt.Println(buf.String())
}