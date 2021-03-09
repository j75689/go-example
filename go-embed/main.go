package main

import (
	_ "embed"
	"fmt"
)

//go:embed Test.pdf
var f []byte

func main() {
	fmt.Println(len(f))
}
