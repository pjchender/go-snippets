package rate

import (
	"fmt"
	"time"

	"golang.org/x/time/rate"
)

func useBucketToken() {
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
