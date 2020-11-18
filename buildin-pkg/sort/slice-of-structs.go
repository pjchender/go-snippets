package sort

import (
	"fmt"
	"sort"
)

type Programmer struct {
	Name string
	Age  int
}

/* 建立一個符合 sort.interface 的  type */
type byAge []Programmer

func (p byAge) Len() int {
	return len(p)
}

func (p byAge) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p byAge) Less(i, j int) bool {
	return p[i].Age < p[j].Age
}

func SortSliceOfStructs() {
	programmers := []Programmer{
		{Name: "Aaron", Age: 30},
		{Name: "Bruce", Age: 20},
		{Name: "Candy", Age: 50},
		{Name: "Derek", Age: 1000},
	}

	sort.Sort(byAge(programmers))               // ascending
	sort.Sort(sort.Reverse(byAge(programmers))) // descending

	fmt.Println(programmers)
}
