package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Product 商品
type Product struct {
	ID               uuid.UUID
	CreatedAt        time.Time
	UpdatedAt        time.Time
	DeletedAt        gorm.DeletedAt
	Name             string
	Price            int64
	IsPublish        bool
	ProviderUniqueID uuid.UUID
}

// ProductExternal 外部使用和檢視的 struct
type ProductExternal struct {
	ID        uuid.UUID `json:"productId"`
	Name      string    `json:"name"`
	Price     int64     `json:"price"`
	IsPublish bool      `json:"isPublish"`
}

// ToExternal 會將 Product 轉成外部使用和檢視的 struct
func (p *Product) ToExternal() *ProductExternal {
	productExternal := ProductExternal{
		ID:        p.ID,
		Name:      p.Name,
		Price:     p.Price,
		IsPublish: p.IsPublish,
	}

	return &productExternal
}

// ProductQuery 提供可以在 url 後帶入使用的 queryString
type ProductQuery struct {
	Name      string `form:"name"`
	IsPublish string `form:"isPublish"`
}

// FilterProducts 會保留 handler 回傳 true 的 Product
func FilterProducts(searchElement []Product, handler func(product Product) bool) []Product {
	n := 0
	for _, element := range searchElement {
		if handler(element) {
			searchElement[n] = element
			n++
		}
	}
	return searchElement[:n]
}
