package slice

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func ExampleContains() {
	s := []int{1, 3, 5}
	fmt.Println("isContain 5", Contains(toIntPointers(s), 5))
	// Output: isContain 5 true

	fmt.Println("isContain 6", Contains(toIntPointers(s), 6))
	// Output: isContain 6 false
}

func TestContains(t *testing.T) {
	s := []int{1, 3, 5}

	assert.True(t, Contains(toIntPointers(s), 5))
	assert.False(t, Contains(toIntPointers(s), 6))
}

func ExampleFilter() {
	s := []int{1, 3, 5}
	biggerThenFive := Filter(toIntPointers(s), func(i *int) bool {
		number := *i
		return number > 2
	})
	fmt.Println(biggerThenFive)
	// Output: [3, 4]
}

func TestFilter(t *testing.T) {
	s := []int{1, 3, 5}
	biggerThenFive := Filter(toIntPointers(s), func(i *int) bool {
		number := *i
		return number > 2
	})

	expect := []int{3, 5}
	assert.Equal(t, expect, toIntValues(biggerThenFive))
}

func ExampleFilterWithContain() {
	// 保留 exist 中不存在 update 內的元素
	exist := []int{2, 4, 6}
	update := []int{1, 3, 6}

	filtered := Filter(toIntPointers(exist), func(i *int) bool {
		number := *i
		return !Contains(toIntPointers(update), number)
	})
	fmt.Println(filtered)
	// Output: [2, 4]
}

func TestFilterWithContain(t *testing.T) {
	exist := []int{2, 4, 6}
	update := []int{1, 3, 6}

	filtered := Filter(toIntPointers(exist), func(i *int) bool {
		number := *i
		return !Contains(toIntPointers(update), number)
	})

	expect := []int{2, 4}
	assert.Equal(t, expect, toIntValues(filtered))
}

func toIntPointer(number int) *int {
	return &number
}

func toIntValue(intPoint *int) int {
	return *intPoint
}

func toIntPointers(intValues []int) []*int {
	intPointers := make([]*int, len(intValues))
	for i := range intValues {
		intPointers[i] = &intValues[i]
	}

	return intPointers
}

func toIntValues(intPoints []*int) []int {
	intValues := make([]int, len(intPoints))
	for i := range intPoints {
		intValues[i] = *intPoints[i]
	}

	return intValues
}
