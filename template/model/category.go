package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Category 商品類別
type Category struct {
	ID        uuid.UUID `gorm:"primaryKey;uniqueIndex;type:uuid;default:uuid_generate_v4()"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
	Name      string

	// Category has many Products
	Products []*Product
}

// CategoryExternal 是給外部 API 檢視和編輯 Category 用
type CategoryExternal struct {
	CategoryID uuid.UUID          `json:"categoryId"`
	Name       string             `json:"name"`
	Products   []*ProductExternal `json:"products"`
}

// ToExternal 可以將 Category 轉成 CategoryExternal
func (c *Category) ToExternal() CategoryExternal {
	categoryExternal := CategoryExternal{
		CategoryID: c.ID,
		Name:       c.Name,
		Products:   ToExternalProducts(c.Products),
	}

	return categoryExternal
}

func ToExternalCategory(categories []*Category) []*CategoryExternal {
	externalCategories := make([]*CategoryExternal, len(categories))
	for i, category := range categories {
		categoryExternal := category.ToExternal()
		externalCategories[i] = &categoryExternal
	}
	return externalCategories
}
