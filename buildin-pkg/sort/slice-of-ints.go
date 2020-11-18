package sort

import (
	"fmt"
	"sort"
)

func SortSliceOfInts() {
	numbers := []int{1, 5, 3, 6, 2}
	sort.Ints(numbers)
	fmt.Println(numbers) // [1 2 3 5 6]，ascending

	sort.Sort(sort.Reverse(sort.IntSlice(numbers)))
	fmt.Println(numbers) // [6 5 3 2 1]，descending
}
