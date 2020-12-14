---
id: golang-pkg-viper
title: '[go-pkg] Viper'
date: 2020-12-14 10:10:10
updated: 2020-12-14 10:10:10
categories:
  - Go
tags:
  - golang
  - pkg
---

# [go-pkg] Viper

> [spf13/viper](https://github.com/spf13/viper)  @ github

Viper 是 Golang App 中用來使用設定檔的工具，它可以用來：

- 讀取 JSON、TOML、YAML、HCL、evnfile、Java 屬性設定檔等等
- 設定預設值
- 自動重新讀取並載入設定檔
- 透過  CLI 中的 flags 來覆蓋設定
- 讀取 buffer 檔
- 讀取遠端的設定檔

## 使用步驟

1. **建立設定檔**：在 `configs` 資料夾中建立 `config.yaml`
2. **讀取設定檔**：在 `setting` 資料夾中建立 `setting.go` 透過 viper 讀取設定檔
3. **轉成 go struct**：將 viper 讀取好的設定檔透過 `Marshal` 轉成 golang 中可使用的 struct
4. **放置設定於全域**：在 `global` 資料夾中建立 `setting.go`，透過 `ReadSection` 將解析過後的設定檔放置於全域以供使用
5. **載入設定檔**：在 `main.go` 中將撰寫好的設定載入 `global` 套件
6. **使用設定檔**：在需要使用設定檔的地方，可以使用 `global` 取得存放於全域的設定

## 範例程式碼

> [go-snippets/pkg/viper](https://github.com/pjchender/go-snippets/tree/master/pkg/viper) @ pjchender github

## 資料來源

- [再強一點：用 Go語言完成六個大型專案](https://www.tenlong.com.tw/products/9789865501501)
