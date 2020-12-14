---
id: go-pkg-time-rate
title: '[Go] pkg - time/rate'
date: 2020-11-16 10:10:10
updated: 2020-11-16 10:10:10
categories:
  - Go
tags:
  - golang
  - build-in-pkg
---

# [Go] pkg - time/rate

<!-- markdownlint-disable MD010 -->

> - [time/rate](https://pkg.go.dev/golang.org/x/time/rate) @ go.pkg
> - [Golang 標準庫限流器 time/rate 使用介紹](https://www.cyhone.com/articles/usage-of-golang-rate/)
> - [Golang time/rate 限速器](https://www.jianshu.com/p/1ecb513f7632) @ 簡書

## 重要概念

在 Golang 中使用 Limiter 來控制在特定頻率內，某個事件是否允許被執行。這個 **Limiter** 是實作 **Token Bucket** 的方式來達到限流的目的，也就是會先設定：

- **event rate（r）**：**將 token 放入桶子的頻率**，例如每秒將放入 n 個 token 到桶子（bucket）中。
- **burst size（b）**：**一個桶子（bucket）中能夠容納的 token 數量**

一開始桶子會是滿的，只要桶子中有剩餘的 Token 就可以取用，若沒有剩餘的 Token 則需要等待後才能取用。

## 建立 Limiter：NewLimiter

##### keywords: `NewLimiter`

使用 `NewLimiter` 來建立一個 non-zero Limiter：

```go
func NewLimiter(r Limit, b int) *Limiter
```

Limiter 包含兩個主要的屬性：

- `r`：rate，型別是 `Limit`（ `Limit` 的型別是 `float64`）， 是用來定義**「每秒」內某事件可以發生的次數**，`zero` 的話表示不允許任何事件發生。可以透過 [`Every(interval time.Duration) Limit`](https://pkg.go.dev/golang.org/x/time/rate#Every) 這個方法來取得 Limit。
- `b`：burst size，表示桶子的大小，也就是桶子中可以放入多少 Token

```go
// r：rate，每秒會放入 10 個 token
// b：burst size，桶子的大小只能容納 1 個 token
limiter := rate.NewLimiter(10, 1)

fmt.Println(limiter.Limit(), limiter.Burst()) // 10, 1
```

也可使用 `Every()` 來產生 `Limit`：

```go
// func Every(interval time.Duration) Limit
//
// r：每 100 毫秒會放入 1 個 token（同樣也是每秒會有 10 個 token）
// b：桶子的大小只能容納 1 個 token
limit := rate.Every(100 * time.Millisecond)
limiter := rate.NewLimiter(limit, 1)

fmt.Println(limiter.Limit(), limiter.Burst()) // 10, 1
```

## 使用 Limiter

##### keywords: `Allow`, `Reserve`, `Wait`, `AllowN`, `ReserveN`, `WaitN`

Limiter 主要有三種方法，分別是 `Allow`, `Reserve` 和 `Wait`，**一般來說最常使用到的是 `Wait`**。這三種方法都需要消耗「一個」 token，**差別在於當 token 不足的時候所採取的行為**。

當 Token 不足時：

- Allow：會回傳 `false`
- Reserve：會回傳 `Reservation`，表示預約未來的 Token 並告知要等多久後才能再次使用
- Wait：會等待那裡（阻塞），直到有足夠的 Token 或該 context 被取消。

如果需要一次消耗多個 Token，則使用 `AllowN`, `ReserveN` 和 `WaitN`。

### Wait/WaitN

```go
func (lim *Limiter) Wait(ctx context.Context) (err error)  // 等同於 WaitN(ctx, 1)
func (lim *Limiter) WaitN(ctx context.Context, n int) (err error)
```

- `WaitN` 會阻塞住，每次執行需要消耗 `n` 個 token，也就是直到有足夠（n）的 token 時才會繼續往後執行
- 在下述情況發生時會回傳錯誤
  - 如果需要消耗的 token 數目（ `n`） 超過 Limiter 水桶的數量（burst size）時
  - Context 被取消（canceled）
  - Context 的等待時間超過 Deadline 時

```go
// 範例程式碼：https://www.jianshu.com/p/1ecb513f7632
func main() {
	counter := 0
	ctx := context.Background()

	// 每 200 毫秒會放一次 token 到桶子（每秒會放 5 個 token 到桶子），bucket 最多容納 1 個 token
	limit := rate.Every(time.Millisecond * 200)
	limiter := rate.NewLimiter(limit, 1)
	fmt.Println(limiter.Limit(), limiter.Burst()) // 5，1

	for {
		counter++
		limiter.Wait(ctx)
		fmt.Printf("counter: %v, %v \n", counter, time.Now().Format(time.RFC3339))
	}
}
```

### Allow/AllowN

```go
func (lim *Limiter) Allow() bool      // 等同於 AllowN(time.Now(), 1)
func (lim *Limiter) AllowN(now time.Time, n int) bool
```

- `AllowN` 表示在某個的時間點時，每次需要消耗 `n` 個 token，若桶子中的 token 數目是否滿足 `n`，則會回傳 `true` 並消耗掉桶子中的 token，否則回傳 `false`
- 只有在你想要 drop / skip 超過 rate limit 的事件時使用，否則使用 `Reserve` 或 `Wait`

```go
// 範例程式碼：https://www.jianshu.com/p/1ecb513f7632
func main() {
	counter := 0

	// event rate：每 200 毫秒會放一次 token 到桶子（每秒會放 5 個 token 桶子）
	// burst size：bucket 最多容納 4 個 token
	limit := rate.Every(time.Millisecond * 200)
	limiter := rate.NewLimiter(limit, 4)
	fmt.Println(limiter.Limit(), limiter.Burst()) // 5，4

	for {
		counter++

		// 每次需要 3 個 token
		if isAllowed := limiter.AllowN(time.Now(), 3); isAllowed {
			fmt.Printf("counter: %v, %v \n", counter, time.Now().Format(time.RFC3339))
		} else {
			fmt.Printf("counter: %v, not allow \n", counter)
			time.Sleep(100 * time.Millisecond)
		}
	}
}
```

### Reserve/ReserveN

```go
func (lim *Limiter) Reserve() *Reservation   // 等同於 ReserveN(time.Now(), 1)
func (lim *Limiter) ReserveN(now time.Time, n int) *Reservation
```

- `ReserveN` 會回傳 `Reservation`，用來指稱還需要等多久才能有足夠的 token 讓事件發生；後續的 Limiter 會把 Reservation 納入考量
- 當 n 超過桶子能夠容納的 token 數量時（即，Limiters 的 burst size），Reservation 的 OK 方法將會回傳 `false`

```go
func main() {
	counter := 0

	// event rate：每 200 毫秒會放一次 token 到桶子（每秒會放 5 個 token 桶子）
	// burst size：bucket 最多容納 3 個 token
	limit := rate.Every(time.Millisecond * 200)
	limiter := rate.NewLimiter(limit, 3)
	fmt.Println(limiter.Limit(), limiter.Burst()) // 5，3

	for {
		counter++
		// 每次執行需要 2 個 token
		tokensNeed := 2
		reserve := limiter.ReserveN(time.Now(), tokensNeed)

		// r.OK() 是 false 表示 n 的數量大於桶子能容納的數量（lim.burst）
		if !reserve.OK() {
			fmt.Printf("一次所需的 token 數（%v）大於桶子能容納 token 的數（%v）\n", tokensNeed, limiter.Burst())
			return
		}

		// reserve.Delay() 可以取得需要等待的時間
		time.Sleep(reserve.Delay())

		// 等待完後做事...
		fmt.Printf("counter: %v, %v \n", counter, time.Now().Format(time.RFC3339))
	}
}
```

- `r.Delay()`：可以得到需要等待的時間，`0` 則表示不用等待

## 調整 Limiter

##### keywords: `SetBurst`, `SetLimit`, `SetBurstAt`, `SetLimitAt`

如果需要動態調整 Limiter 的數率和桶子的大小，則可以使用 `SetBurst` 或 `SetLimit` 的方法。

## 整合 GIN 限制向 client 發送 Request 的次數

### 限制特定 usecase / API 中的 limiter

#### usecase (API)

- 在 `PostUsecase` 的 struct 中定義 `Limiter` 的型別
- 在 `router/post.go` 中使用 `NewLimiter()` 來建立 Limiter
- 在 `GetPost` 中透過 `Limiter.Wait()` 來限制發送請求的頻率

```go
// usecase/post.go

// STEP 1：在 struct 中定義 limiter，並在 router/post.go 中建立 Limiter
type PostUsecase struct {
	Limiter *rate.Limiter
}

func (p *PostUsecase) GetPost(ctx *gin.Context) {
	id := ctx.Param("id")

  // STEP 3：使用 Limiter.Wait，每次會消耗桶子中的一個 token
	p.Limiter.Wait(ctx)

  // STEP 4：實際發送請求
	post := getPost(id)

	ctx.JSON(http.StatusOK, post)
}
```

#### router

- 使用 `NewLimiter()` 來建立 Limiter
  - `rate.Every(200 * time.Millisecond)`：每 200 毫秒會放入一個 token 到桶子（bucket）中
  - `rate.NewLimiter(limit, 1)`：桶子的容量（burst size）為 1 個 token

```go
// router/post.go

func registerPosts(router *gin.Engine) {

  // STEP 2：使用 NewLimiter() 來建立 Limiter
	limit := rate.Every(1000 * time.Millisecond)
	limiter := rate.NewLimiter(limit, 1)
	postHandler := &usecase.PostUsecase{
		Limiter: limiter,
	}

	router.GET("/posts/:id", postHandler.GetPost)
}
```

### 限制多支 usecase / API 的 limiter

#### 撰寫 limiter package

如果是很多不同支 API 都需要限制流量的話，則可以建立一個獨立的 package：

```go
// ./pkg/limiter/limiter.go
package limiter

// STEP 1：建立 limiter
// rate：每秒會放 1 個 token 到 bucket 中
// burst size：桶子最多可以容納 1 個 bucket
var RateLimiter = rate.Every(time.Millisecond * 1000)
var RequestLimiter = rate.NewLimiter(RateLimiter, 1)
```

#### 在 API 中使用 limiter

並在需要限流的 limiter 中使用它：

```go
// ./usecase/post.go
package usecase

import "sandbox/gin-sandbox/pkg/limiter"

func (p *PostUsecase) GetPost(ctx *gin.Context) {

  // STEP 3：使用建立好的 limiter
  // 每次需要消耗桶子中的 1 個 bucket
	limiter.RequestLimiter.Wait(ctx)

	post := getPost(id)

	ctx.JSON(http.StatusOK, post)
}
```

在另一支需要限流的 API 中使用寫好的 limiter：

```go
// ./usecase/healthcheck.go

import "sandbox/gin-sandbox/pkg/limiter"

package usecase

func (h *HealthCheckUsecase) Pong(ctx *gin.Context) {

	limiter.RequestLimiter.Wait(ctx)

	ctx.JSON(http.StatusOK, gin.H{
		"message":    "pong",
		"threadNum,": threadNum,
		"counter":    counter,
	})
}

```

### 使用 JMeter 測試結果

若我們的 Limiter 限制每秒給一個 token 到 bucket 中，且 bucket 的 burst size（能夠容納的 token 數量）為 1 時，表示每秒只能處理一個請求。

若以 JMeter 進行測試，可以看到 Throughput（流量）的欄位即為 `1.0/sec`：

![Screen Shot 2020-11-16 at 4.31.57 PM](https://i.imgur.com/iwBGT82.png)

## 範例程式碼

- [gin-limiter-example](https://github.com/pjchender/go-gin-example/tree/gin-limiter-example) @ github

## 參考

- [time/rate](https://pkg.go.dev/golang.org/x/time/rate) @ go.pkg
- [Golang 標準庫限流器 time/rate 使用介紹](https://www.cyhone.com/articles/usage-of-golang-rate/)
- [Golang time/rate 限速器](https://www.jianshu.com/p/1ecb513f7632) @ 簡書
