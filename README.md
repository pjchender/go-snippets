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

## Template

使用 GIN 框架，並搭配下述架構：

- API：負責把 HTTP request 的參數取出（例如，`ShouldBind`），並帶入 Service 方法
- Service：
  1. 定義呼叫 database 時所需要的參數（例如，`GetCategoryRequest`），提供不同 delivery 進行使用（例如 HTTP 或 gRPC）
  2. 根據 delivery 提供的參數呼叫 database 中的方法，因此除了 API 會使用 service 之外，gRPC 也只需把要用到的參數取出後呼叫 Service 方法即可
- Database：定義直接操作 Database 的方法
- Model：用來定義 GORM 的資料結構，直接與 Database 建立的型別有關
