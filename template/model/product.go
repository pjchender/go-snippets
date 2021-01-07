package model

import (
	"errors"
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

// Count 會根據 Product 的 Name, IsPublish 和
func (p *Product) Count(db *gorm.DB) (int64, error) {
	var count int64
	if p.Name != "" {
		db = db.Where("name = ?", p.Name)
	}

	db = db.Where("is_publish = ?", p.IsPublish)
	err := db.Model(&p).Where("deleted_at IS NULL").Count(&count).Error
	if err != nil {
		return 0, err
	}

	return count, nil
}

// List 會根據 pageOffset 和 pageSize 以及 name, state 的條件回傳對應的資料
func (p *Product) List(db *gorm.DB, pageOffset, pageSize int) ([]*Product, error) {
	var products []*Product
	var err error
	if pageOffset >= 0 && pageSize > 0 {
		db = db.Offset(pageOffset).Limit(pageSize)
	}
	if p.Name != "" {
		db = db.Where("name = ?", p.Name)
	}
	db = db.Where("is_publish = ?", p.IsPublish)
	if err = db.Where("deleted_at IS NULL").Find(&products).Error; err != nil {
		return nil, err
	}

	return products, nil
}

// ListByIDs 會根據給的 ids 回傳對應的資料
func (p *Product) ListByIDs(db *gorm.DB, ids []uint32) ([]*Product, error) {
	var products []*Product
	db = db.Where("deleted_at IS NULL")
	err := db.Where("id IN (?)", ids).Find(&products).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	return products, nil
}

// Get 會根據 id 取得某一筆 Product 的資料
func (p Product) Get(db *gorm.DB) (Product, error) {
	var product Product
	err := db.Where("id = ? AND is_publish = ? AND deleted_at IS NULL", p.ID, p.IsPublish).First(&product).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return product, err
	}

	return product, nil
}

// Create 會建立新的 Product
func (p *Product) Create(db *gorm.DB) error {
	return db.Create(&p).Error
}

// Update 會更新某筆 Product 的資料
func (p *Product) Update(db *gorm.DB, values interface{}) error {
	return db.Model(&p).Where("id = ? AND deleted_at IS NULL", p.ID).Updates(values).Error
}

// Delete 會刪除某筆 Product
func (p *Product) Delete(db *gorm.DB) error {
	return db.Where("deleted_at IS NULL").Delete(&p).Error
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
func (p *Product) ToExternal() *ProductExternal {
	productExternal := ProductExternal{
		ProductID:  p.ID,
		Name:       p.Name,
		Price:      p.Price,
		IsPublish:  p.IsPublish,
		CategoryID: p.CategoryID,
		Category:   p.Category.ToExternal(),
	}

	return &productExternal
}

// ProductQuery 提供可以在 url 後帶入使用的 queryString
type ProductQuery struct {
	ProductID string `form:"productId"`
	Name      string `form:"name"`
	IsPublish string `form:"isPublish"`

	// categoryID 雖然是 uuid 但透過 queryString 傳的時候只能是字串
	CategoryID string `form:"categoryId"`

	// query 時間的話可以用 int64 後續轉成 Unix
	CreatedAt int64 `form:"createdAt"`
	BeginDate int64 `form:"beginDate"`
	EndDate   int64 `form:"endDate"`
}

// FilterProducts 會保留 handler 回傳 true 的 Product
func FilterProducts(searchElement []*Product, handler func(product *Product) bool) []*Product {
	n := 0
	for _, element := range searchElement {
		if handler(element) {
			searchElement[n] = element
			n++
		}
	}
	return searchElement[:n]
}

// ProductForUpdate 會對應到 database/product 中的 UpdateProductWithZero 方法
// 之所以需要這個 struct 是因為當 UpdateProductWithZero 沒有對應欄位的 value 時，會用 zero value 覆蓋
// 因此在更新時一定要把對應的值都給進去
type ProductForUpdate struct {
	ID        uuid.UUID
	Name      string
	Price     int64
	IsPublish bool
}
