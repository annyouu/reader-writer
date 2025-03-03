package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	// io.Copyの例(全てコピー)
	fmt.Println("=== io.Copy Example ===")
	r1 := strings.NewReader("Hello, 世界")
	if _, err := io.Copy(os.Stdout, r1); err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println() // 改行

	// io.CoyeNの例(5バイトだけコピー)
	fmt.Println("=== io.CopyN Example ===")
	r2 := strings.NewReader("Hello, 世界")
	if _, err := io.CopyN(os.Stdout, r2, 5); err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println() // 改行
}