package database

import (
	"time"

	"github.com/google/uuid"
	"github.com/pjchender/go-snippets/template/model"
	"gorm.io/gorm/clause"
)

// CreateProduct 會建立新的 Product
func (g *GormDatabase) CreateProduct(product *model.Product) error {
	return g.DB.Create(product).Error
}

// UpdateProductWithoutZero 會更新 Product 中 non-zero 的欄位
// 透過 struct 來更新 Product 時，只會更新 non-zero 的欄位
func (g *GormDatabase) UpdateProductWithoutZero(product *model.Product) error {
	return g.DB.Updates(product).Error
}

// UpdateProductWithZero 會更新有在 map 中列出的欄位（包含 zero-value）
func (g *GormDatabase) UpdateProductWithZero(product *model.Product) error {
	return g.DB.Model(&product).Updates(map[string]interface{}{
		"name":       product.Name,
		"price":      product.Price,
		"is_publish": false,
	}).Error
}

// UpsertProductByProviderWithoutZero 會以 Upsert 的方式更新 non-zero 的欄位
func (g *GormDatabase) UpsertProductByProviderWithoutZero(productInfo *model.Product) (*model.Product, error) {
	var product model.Product

	err := g.DB.Where("provider_unique_id = ?", productInfo.ProviderUniqueID).
		Assign(productInfo).
		FirstOrCreate(&product).Error
	if err != nil {
		return nil, err
	}

	return &product, nil
}

// UpsertProductByProviderWithZero 會以 Upsert 的方式更新有列出的欄位（包含 zero-value）
// OnConflict 中的 Column 需要是 index
func (g *GormDatabase) UpsertProductByProviderWithZero(product *model.Product) (*model.Product, error) {
	var err error

	err = g.DB.Clauses(clause.OnConflict{
		Columns: []clause.Column{{
			Name: "provider_unique_id",
		}},
		DoUpdates: clause.AssignmentColumns([]string{
			"name", "price", "is_publish",
		}),
	}).Create(product).Error
	if err != nil {
		return nil, err
	}

	var updatedProduct model.Product
	err = g.DB.Preload(clause.Associations).Where("provider_unique_id = ?", product.ProviderUniqueID).Take(&updatedProduct).Error
	if err != nil {
		return nil, err
	}

	return &updatedProduct, nil
}

// GetProducts 會取得 DB 中所有未被刪除的 Products
func (g *GormDatabase) GetProducts() ([]*model.Product, error) {
	var products []*model.Product

	err := g.DB.Find(&products).Error
	if err != nil {
		return nil, err
	}

	return products, nil
}

// GetProductsWithConditions 可以根據條件篩選 Products
func (g *GormDatabase) GetProductsWithConditions(
	beginDate, endDate time.Time,
	conditions ...interface{},
) ([]*model.Product, error) {
	var products []*model.Product
	tx := g.DB.Table("products").Preload(clause.Associations)

	// query 時間區間
	if !beginDate.IsZero() {
		tx.Where("created_at >= ?", beginDate)
	}

	if !endDate.IsZero() {
		tx.Where("created_at <= ?", endDate)
	}

	if len(conditions) == 1 {
		tx.Where(conditions[0])
	} else if len(conditions) > 1 {
		tx.Where(conditions[0], conditions[1:]...)
	}

	err := tx.Find(&products).Error
	if err != nil {
		return nil, err
	}

	return products, nil
}

// GetProductByID 會根據 Product 的 ID 來取得 Product
func (g *GormDatabase) GetProductByID(productID uuid.UUID) (*model.Product, error) {
	product := model.Product{}
	err := g.DB.Preload(clause.Associations).Take(&product, productID).Error
	if err != nil {
		return nil, err
	}

	return &product, nil
}

// GetProductByProvider 會根據 ProviderUniqueID 這個欄位來取得 Product
func (g *GormDatabase) GetProductByProvider(providerUniqueID string) (*model.Product, error) {
	product := model.Product{}
	err := g.DB.Preload(clause.Associations).Where("provider_unique_id = ?", providerUniqueID).Take(&product).Error
	if err != nil {
		return nil, err
	}

	return &product, nil
}

// DeleteProductByID 會刪除該 Product 及與其相關連的 records
func (g *GormDatabase) DeleteProductByID(productID uuid.UUID) error {
	return g.DB.Select(clause.Associations).Delete(&model.Product{ID: productID}).Error
}
