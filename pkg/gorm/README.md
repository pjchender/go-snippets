---
id: golang-pkg-gorm-db-connect
title: '[go-pkg] Gorm'
date: 2020-12-14 10:10:10
updated: 2020-12-14 10:10:10
categories:
  - Go
tags:
  - golang
  - pkg
---

# [go-pkg] Gorm Connect Database

> [Connecting to a Database](https://gorm.io/docs/connecting_to_the_database.html)

## 使用步驟

1. 在 `setting/setting.go` 中建立與 Database 有關的設定（可以搭配 viper 讀取 config 檔）
2. 在 `model/model.go` 中建立 `NewDBEngine` 的方法，根據設定與 DB 建立連線
3. 在 `global/global.go` 中建立全域變數 `DBEngine`
4. 在 `main.go` 中建立 `setupDBEngine` 的方法，將連線好的 gorm.DB 保存到 `global` package中

## 範例程式碼

> [go-snippets/pkg/gorm](https://github.com/pjchender/go-snippets/tree/master/pkg/gorm) @ pjchender github

## 資料來源

- [再強一點：用 Go語言完成六個大型專案](https://www.tenlong.com.tw/products/9789865501501)
