package main

import "golang.org/x/text/message"

func main() {
	p := message.NewPrinter(message.MatchLanguage("en"))
	p.Printf("%.2f\n", 123123.12312)
}
