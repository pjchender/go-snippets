package slice

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func ExampleContains() {
	s := []int{1, 3, 5}
	fmt.Println("isContain 5", Contains(s, 5))
	// Output: isContain 5 true

	fmt.Println("isContain 6", Contains(s, 6))
	// Output: isContain 6 false
}

func TestContains(t *testing.T) {
	s := []int{1, 3, 5}

	assert.True(t, Contains(s, 5))
	assert.False(t, Contains(s, 6))
}

func ExampleFilter() {
	s := []int{1, 3, 5}
	biggerThenFive := Filter(s, func(i int) bool {
		return i > 2
	})
	fmt.Println(biggerThenFive)
	// Output: [3, 4]
}

func TestFilter(t *testing.T) {
	s := []int{1, 3, 5}
	biggerThenFive := Filter(s, func(i int) bool {
		return i > 2
	})

	expect := []int{3, 4}
	assert.Equal(t, expect, biggerThenFive)
}

func ExampleFilterWithContain() {
	// 保留 exist 中不存在 update 內的元素
	exist := []int{2, 4, 6}
	update := []int{1, 3, 6}

	filtered := Filter(exist, func(i int) bool {
		return !Contains(update, i)
	})
	fmt.Println(filtered)
	// Output: [2, 4]
}

func TestFilterWithContain(t *testing.T) {
	exist := []int{2, 4, 6}
	update := []int{1, 3, 6}

	filtered := Filter(exist, func(i int) bool {
		return !Contains(update, i)
	})

	expect := []int{2, 4}
	assert.Equal(t, expect, filtered)
}
