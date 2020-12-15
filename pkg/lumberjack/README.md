---
id: golang-snippets-lumberjack-logger
title: '[go] Snippets - lumberjack (logger)'
date: 2020-12-14 10:10:10
updated: 2020-12-14 10:10:10
categories:
  - Go
tags:
  - golang
  - pkg
  - logger
---

# [Go] pkg - lumberjack(logger)

## 步驟

1. 在 `logger` 資料夾中建立 `logger.go`
2. 在 `global` 資料夾中建立 `logger.go` 和 `setting.go`
3. 在 `main.go` 中撰寫 `setupLogger` 的方法
4. 在 `main.go` 中的 `init` 呼叫 `setupLogger`，並於 `main` 中使用 Logger

## 使用方式

```go
func main() {
	global.Logger.Infof("%s: go-snippets/%s", "pkg", "blog-service")
}
```

## 資料來源

- [再強一點：用 Go語言完成六個大型專案](https://www.tenlong.com.tw/products/9789865501501)
