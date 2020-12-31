package main

import (
	"fmt"

	"github.com/pjchender/go-snippets/helpers/enum"
)

func main() {
	req := "canceled"
	orderStatus, err := enum.ToOrderStatus(req)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(orderStatus)
}
