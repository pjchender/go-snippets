---
id: golang-snippets-errcode
title: '[go] Snippets - errcode'
date: 2020-12-14 10:10:10
updated: 2020-12-14 10:10:10
categories:
  - Go
tags:
  - golang
  - pkg
---

# [Go] helpers - errCode

## 步驟

1. 在 `errcode.go` 中定義 `Error` 的 struct
2. 建立產生 `Error` 的 `NewError` 方法
3. 定義 Error Struct 可以使用的方法
4. 在 `common_code.go` 中透過 `NewError` 方法定義和產生所需要的 Error

## 使用方式

```go
func main() {
	fmt.Printf("%+v \n", errcode.TooManyRequest) // 錯誤 10000007, 錯誤訊息：請求過多
}
```

## 資料來源

- [再強一點：用 Go語言完成六個大型專案](https://www.tenlong.com.tw/products/9789865501501)
