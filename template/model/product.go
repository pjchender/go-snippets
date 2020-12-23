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

	// Belongs to Category
	CategoryID uuid.UUID
	Category   Category
}

// ProductExternal 是給外部 API 檢視和編輯 Category 用
type ProductExternal struct {
	ProductID uuid.UUID `json:"productId"`
	Name      string    `json:"name"`
	Price     int64     `json:"price"`
	IsPublish bool      `json:"isPublish"`

	// 由於是要給外部 API 使用，所以 nested 在內的 model 也要轉成 external
	CategoryID uuid.UUID        `json:"categoryId"`
	Category   CategoryExternal `json:"category"`
}

// ToExternal 會將 Product 轉成外部使用和檢視的 struct
func (p *Product) ToExternal() ProductExternal {
	productExternal := ProductExternal{
		ProductID:  p.ID,
		Name:       p.Name,
		Price:      p.Price,
		IsPublish:  p.IsPublish,
		CategoryID: p.CategoryID,
		Category:   p.Category.ToExternal(),
	}

	return productExternal
}

// ProductQuery 提供可以在 url 後帶入使用的 queryString
type ProductQuery struct {
	ProductID string `form:"productId"`
	Name      string `form:"name"`
	IsPublish string `form:"isPublish"`

	// categoryID 雖然是 uuid 但透過 queryString 傳的時候只能是字串
	CategoryID string `form:"categoryId"`
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
