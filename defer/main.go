package main

import (
	"errors"
	"fmt"
)

func main() {
	getStr()
}

func getStr() (s string, err error) {
	s = "test"
	fmt.Println(s)
	defer func() {
		if err == nil {
			fmt.Println(err, s)
		}
	}()
	return testDefer()
}

func testDefer() (string, error) {
	return "aaaa", errors.New("test error")
}

// func main() {
// 	cache, err := bigcache.NewBigCache(bigcache.DefaultConfig(1 * time.Second))
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	cache.Set("test", []byte("ttttt"))
// 	b, err := cache.Get("test")
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	fmt.Printf("%s\n", b)
// 	time.Sleep(2 * time.Second)
// 	b, err = cache.Get("test")
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	fmt.Printf("%s\n", b)
// }
