package database

import (
	"github.com/google/uuid"
	"github.com/pjchender/go-snippets/template/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// CreateCategory 會建立新的 Category 及 nested 在內的 Product
func (d *GormDatabase) CreateCategory(category *model.Category) error {
	return d.DB.Create(category).Error
}

// CreateCategories 可以一次建立多筆的 Categories
func (d *GormDatabase) CreateCategories(categories []*model.Category) error {
	return d.DB.Create(&categories).Error
}

// UpdateCategory 會更新 Category 及其 nested 在內的 Product
func (d *GormDatabase) UpdateCategory(category *model.Category) error {

	// 定義能被更新的 db 欄位
	values := map[string]interface{}{
		"name": category.Name,
	}

	// 使用 FullSaveAssociations 才會自動更新關聯在 Category 內的 Product
	return d.DB.Session(&gorm.Session{FullSaveAssociations: true}).
		Model(model.Category{ID: category.ID, Products: category.Products}).
		Updates(values).Error
}

// GetCategories 會同時撈出 Categories 和 nested 在內的 Products
func (d *GormDatabase) GetCategories(conditions ...interface{}) ([]*model.Category, error) {
	var categories []*model.Category

	// Preload 中的名稱需要跟著 model.Category 中的定義，所以這裡會用大寫的 Products
	tx := d.DB.Preload("Products")

	if len(conditions) == 1 {
		tx.Where(conditions[0])
	} else if len(conditions) > 1 {
		tx.Where(conditions[0], conditions[1:]...)
	}

	err := tx.Find(&categories).Error
	if err != nil {
		return nil, err
	}

	return categories, nil
}

// GetCategoryByID 會根據 ID 回傳 Category（包含 Products）
func (d *GormDatabase) GetCategoryByID(categoryID uuid.UUID) (*model.Category, error) {
	category := model.Category{}
	err := d.DB.Preload("Products").Take(&category, categoryID).Error
	if err != nil {
		return nil, err
	}

	return &category, nil
}

// DeleteCategoryByID 會根據 ID 刪除 Category，但不會刪除與其關聯的 Product）
func (d *GormDatabase) DeleteCategoryByID(categoryID uuid.UUID) error {
	// 只會刪除 Category，不會刪除與其關聯的 Product
	return d.DB.Delete(&model.Category{ID: categoryID}).Error
}

// DeleteCategoryByIDAndAssociatedProduct 會根據 ID 刪除 Category，同時刪除與其關聯的 Product）
func (d *GormDatabase) DeleteCategoryByIDAndAssociatedProduct(categoryID uuid.UUID) error {
	// 使用 Select 可以把和該 Category 有關聯的 Product 一併刪除
	return d.DB.Select(clause.Associations).Delete(&model.Category{ID: categoryID}).Error
}

// DeleteCategoryByIDAndAssociation 會根據 categoryID 將 Category 刪除
// 同時將與該 Category 關聯的 Product.CategoryID 一併設成 null（移除關係但不刪除 Product）
func (d *GormDatabase) DeleteCategoryByIDAndAssociation(categoryID uuid.UUID) error {
	var err error

	// 找出所有與 categoryID 相關連的 Products
	var associationProducts []*model.Product
	err = d.DB.Model(&model.Category{ID: categoryID}).Association("Products").Find(&associationProducts)
	if err != nil {
		return err
	}

	// 將所有和該 categoryID 相關連的 Products.CategoryID 設為 null（移除關係）
	err = d.DB.Model(&model.Category{ID: categoryID}).Association("Products").Delete(&associationProducts)
	if err != nil {
		return err
	}

	// 刪除該筆 Category
	return d.DB.Delete(&model.Category{ID: categoryID}).Error
}
