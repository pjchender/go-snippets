package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Category struct {
	ID        uuid.UUID `gorm:"primaryKey;uniqueIndex;type:uuid;default:uuid_generate_v4()"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
	Name      string
}

// CategoryExternal 是給外部 API 檢視和編輯 Category 用
type CategoryExternal struct {
	CategoryID uuid.UUID `json:"categoryId"`
	Name       string    `json:"name"`
}

// ToExternal 可以將 Category 轉成 CategoryExternal
func (c *Category) ToExternal() CategoryExternal {
	categoryExternal := CategoryExternal{
		CategoryID: c.ID,
		Name:       c.Name,
	}

	return categoryExternal
}
