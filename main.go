package main

import (
	"fmt"

	"github.com/pjchender/go-snippets/helpers/enum"
)

func main() {
	req := "foo"
	err := enum.IsValidOrderStatus(req)
	if err != nil {
		fmt.Println(err)
	}
}
