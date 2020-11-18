package main

import (
	"encoding/json"
	"fmt"

	"github.com/pjchender/go-snippets/configor"
)

func main() {
	configuration := configor.Get()

	data, _ := json.MarshalIndent(configuration, "", "  ")
	fmt.Println(string(data))
}
