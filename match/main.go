package main


import (
	"fmt"
	"io"
	"regexp"
	"strings"
)

func main() {
	// 正規表現をコンパイル
	re, err := regexp.Compile(`(\d+)年(\d+)月(\d+)日`)
	if err != nil {
		fmt.Println("正規表現のコンパイルに失敗:", err)
		return
	}

	// []byte型が正規表現にマッチするか
	fmt.Println(re.Match([]byte("1986年01月12日")))

	// 文字列がマッチするか
	fmt.Println(re.MatchString("1986年01月12日"))

	// io.RuneReaderがマッチするか
	var r io.RuneReader = strings.NewReader("1986年01月12日")
	fmt.Println(re.MatchReader(r))
}


// import (
// 	"fmt"
// 	"io"
// 	"strings"
// )

// func main() {
// 	// こんにちはという文字列を1つずつ読み込む
// 	var reader io.RuneReader = strings.NewReader("こんにちは")

// 	// 1文字ずつで読み込む
// 	for {
// 		r, size, err := reader.ReadRune()
// 		if err == io.EOF {
// 			break // 読み込み完了
// 		}
// 		fmt.Printf("文字: %c, サイズ: %dバイト\n", r, size)
// 	}
// }
