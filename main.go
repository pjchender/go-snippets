package main

import (
	"fmt"
	"github.com/pjchender/go-snippets/lib"
)

func main() {
	exist := []int{2, 4, 6}
	update := []int{1, 3, 6}

	deleted := lib.Filter(exist, func(i int) bool {
		return !lib.Contains(update, i)
	})

	fmt.Println(deleted)
}
