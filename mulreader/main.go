package main

import (
	"io"
	"os"
	"strings"
)

func main() {
	r1 := strings.NewReader("Hello, ")
	r2 := strings.NewReader("世界\n")
	r := io.MultiReader(r1, r2)

	io.Copy(os.Stdout, r)
}