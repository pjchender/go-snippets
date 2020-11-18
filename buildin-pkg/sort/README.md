---
title: '[Go] pkg - sort'
date: 2020-11-18 10:10:10
updated: 2020-11-18 10:10:10
categories:
  - Go
tags:
  - golang
  - build-in-pkg
---

# [Go] pkg - sort

- [sort](https://pkg.go.dev/sort)

## 反向排序

##### keywords: `sort.Reverse()`

```go
func main() {
	numbers := []int{1, 5, 3, 6, 2}
	sort.Ints(numbers)
	fmt.Println(numbers) // [1 2 3 5 6]，ascending

	sort.Sort(sort.Reverse(sort.IntSlice(numbers)))
	fmt.Println(numbers) // [6 5 3 2 1]，descending
}
```

## 客製化欄位

##### keywords: `sort.Sort()` `sort.Interface`

只要有實作 `sort.Interface` 的 slice of structs 都可以使用 `sort.Sort()` 進行排序：

```go
// https://tutorialedge.net/golang/go-sorting-with-sort-tutorial/
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

func main() {
	programmers := []Programmer{
		{Name: "Aaron", Age: 30},
		{Name: "Bruce", Age: 20},
		{Name: "Candy", Age: 50},
		{Name: "Derek", Age: 1000},
	}

	sort.Sort(byAge(programmers))

	fmt.Println(programmers)
}
```

## 參考

- [Go Sorting With the sort Package - Tutorial](https://tutorialedge.net/golang/go-sorting-with-sort-tutorial/) @ tutorialEdge
