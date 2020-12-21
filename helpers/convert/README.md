# pkg/convert

比較需要留意的是，一開始定義的 `StrTo` 這個 type，可以直接透過 `string(StrTo)` 再轉回 string：

```go
type StrTo string

func (s StrTo) String() string {
	return string(s)
}
```