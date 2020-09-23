package main

import (
	"encoding/json"
	"fmt"

	"github.com/pjchender/go-snippets/config"
)

func main() {
	configuration := config.Get()

	data, _ := json.MarshalIndent(configuration, "", "  ")
	fmt.Println(string(data))
}
