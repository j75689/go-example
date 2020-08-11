package main

import (
	"fmt"

	"github.com/shopspring/decimal"
)

func main() {
	v := float64(0)
	fmt.Println(v)
	fmt.Println(-v)

	fmt.Println(decimal.NewFromFloat(-v).IsZero())
}
