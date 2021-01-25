// service 會根據參數呼叫 database 中的方法
// 因此除了 API 會使用 service 之外，gRPC 也只需把要用到的參數取出後呼叫 Service 方法即可

package service

import (
	"context"

	"github.com/pjchender/go-snippets/template/database"
)

// Service 可以使用 dao
type Service struct {
	ctx context.Context
	db  *database.GormDatabase
}

// New 會回傳 Service 的 instance
func New(ctx context.Context, db *database.GormDatabase) Service {
	service := Service{
		ctx: ctx,
		db:  db,
	}
	return service
}
