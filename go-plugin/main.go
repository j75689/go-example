package main

import (
	"go-plugin/plugin"
	"log"
)

func main() {
	plugin.Load("./plugin")
	variable, err := plugin.Execute("equal", struct {
		Target interface{} `json:"target"`
		Value  interface{} `json:"value"`
	}{
		Target: 1,
		Value:  1,
	}, make(map[string]interface{}))

	log.Println(variable, err)
}
