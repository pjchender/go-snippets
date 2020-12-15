---
id: golang-snippets
title: '[go] Snippets'
date: 2020-12-14 10:10:10
updated: 2020-12-14 10:10:10
categories:
  - Go
tags:
  - golang
  - pkg
---

# Golang Snippets

## PKG

- [pkg/configor](https://github.com/pjchender/go-snippets/tree/master/pkg/configor)：和 viper 類似，可以用來讀取 configuration 的檔案
- [pkg/godotenv](https://github.com/pjchender/go-snippets/tree/master/pkg/godotenv)：可以用來載入 `.env` 的檔案
- [pkg/gorm](https://github.com/pjchender/go-snippets/tree/master/pkg/gorm)：用來處理 Database 的 ORM
- [pkg/jwt](https://github.com/pjchender/go-snippets/tree/master/pkg/jwt)：可以用來建立及讀取 JWT
- pkg/lumberjack：可以用來產生 rolling logger
- [pkg/viper](https://github.com/pjchender/go-snippets/tree/master/pkg/viper)：可以用來讀取 App 中的設定檔（configuration）
- [pkg/lumberjack](https://github.com/pjchender/go-snippets/tree/master/pkg/lumberjack)：搭配 `lumberjack` 來建立 log，並保存在硬碟中


## Helpers

- [helpers/errcode](https://github.com/pjchender/go-snippets/tree/master/helpers/errcode)：用來設定 App 內的錯誤類型，並以 httpStatusCode 的方式回傳
