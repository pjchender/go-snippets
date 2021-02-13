package service

import (
	"github.com/google/uuid"
	"github.com/pjchender/go-snippets/template/model"
	"github.com/pjchender/go-snippets/template/pkg/app"
)

type CategoryQuery struct {
	Name string `form:"name"`
}

// 這裏要根據 Database 需要接受的參數來定義型別
// GetCategoryRequest 是從路由的 param 取得
type GetCategoryRequest struct {
	ID uuid.UUID
}

// CountCategoryRequest 是從路由的 queryString 取得
type CountCategoryRequest struct {
	Name string `form:"name"`
}

// ListCategoryRequest 是從路由的 queryString 取得
type ListCategoryRequest struct {
	Name string `form:"name"`
}

// CreateCategoryRequest 是從 JSON 取得
type CreateCategoryRequest struct {
	Name     string                  `json:"name" binding:"required"`
	Products []*UpsertProductRequest `json:"products"`
}

// UpdateCategoryRequest 是從 JSON 取得
type UpdateCategoryRequest struct {
	ID       uuid.UUID               `json:"id"`
	Name     string                  `json:"name"`
	Products []*UpsertProductRequest `json:"products"`
}

// DeleteCategoryRequest 是從路由的 param 取得
type DeleteCategoryRequest struct {
	ID uuid.UUID
}

func (svc *Service) CountCategory(param CountCategoryRequest) (int64, error) {
	return svc.db.CountCategory(param)
}

func (svc *Service) GetCategoryByID(param GetCategoryRequest) (*model.Category, error) {
	category, err := svc.db.GetCategoryByID(param.ID)
	if err != nil {
		return nil, err
	}
	return category, nil
}

func (svc *Service) GetCategoryList(param ListCategoryRequest, paging *app.Paging) ([]*model.Category, error) {
	categories, err := svc.db.ListCategory(paging.Page, paging.PageSize, model.Category{
		Name: param.Name,
	})

	if err != nil {
		return nil, err
	}

	return categories, nil
}

func (svc *Service) CreateCategory(param CreateCategoryRequest) (*model.Category, error) {
	category := model.Category{
		Name:     param.Name,
		Products: ToInternalProducts(param.Products),
	}

	err := svc.db.CreateCategory(&category)

	if err != nil {
		return nil, err
	}

	return &category, nil
}

// CreateCategories 可以一次建立多筆 Category
func (svc *Service) CreateCategories(params []*CreateCategoryRequest) ([]*model.Category, error) {
	categories := make([]*model.Category, len(params))
	for i, param := range params {
		categories[i] = &model.Category{
			Name:     param.Name,
			Products: ToInternalProducts(param.Products),
		}
	}

	err := svc.db.CreateCategories(categories)
	if err != nil {
		return nil, err
	}

	return categories, nil
}

// UpdateCategory 會先檢查該 CategoryID 是否存在，不存在則拋錯
// 並回傳更新後的 Category
func (svc *Service) UpdateCategory(param UpdateCategoryRequest) (*model.Category, error) {
	var err error

	category := model.Category{
		ID:       param.ID,
		Name:     param.Name,
		Products: ToInternalProducts(param.Products),
	}

	// 先確認該筆 record 存在，找不到該 record 則回傳 404
	_, err = svc.db.GetCategoryByID(param.ID)
	if err != nil {
		return nil, err
	}

	// 更新 Category
	err = svc.db.UpdateCategory(&category)
	if err != nil {
		return nil, err
	}

	// 取得更新後的 category
	updatedCategory, err := svc.db.GetCategoryByID(param.ID)
	if err != nil {
		return nil, err
	}

	return updatedCategory, nil
}

// DeleteCategory 會先檢查該 CategoryID 是否存在
func (svc *Service) DeleteCategory(param DeleteCategoryRequest) error {
	var err error

	// 先確認該筆 record 存在，找不到該 record 則回傳 404
	_, err = svc.db.GetCategoryByID(param.ID)
	if err != nil {
		return err
	}

	err = svc.db.DeleteCategoryByID(param.ID)
	if err != nil {
		return err
	}

	return nil
}
