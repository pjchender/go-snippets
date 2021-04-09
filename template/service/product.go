package service

import (
	"time"

	"github.com/google/uuid"
	"github.com/pjchender/go-snippets/template/model"
)

// CreateProductRequest 是使用者在建立 Product 時需要帶入的參數
type CreateProductRequest struct {
	Name      string `form:"name" binding:"required"`
	IsPublish string `form:"isPublish" binding:"required"`
	Price     int64  `form:"price" binding:"required,gte=1"`

	// categoryID 雖然是 uuid 但透過 queryString 傳的時候只能是字串
	CategoryID string `form:"categoryId" binding:"required"`
}

// UpdateProductRequest 會對應到 database/product 中的 UpdateProductWithZero 方法
// 之所以需要這個 struct 是因為當 UpdateProductWithZero 沒有對應欄位的 value 時，會用 zero value 覆蓋
// 因此在更新時一定要把對應的值都給進去
type UpdateProductRequest struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Price     int64     `json:"price"`
	IsPublish bool      `json:"isPublish"`
}

// UpsertProductRequest 會從 JSON 來
type UpsertProductRequest struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Price     int64     `json:"price"`
	IsPublish bool      `json:"isPublish"`
}

// ListProductRequest 提供可以在 url 後帶入使用的 queryString
type ListProductRequest struct {
	ProductID  uuid.UUID
	Name       string
	IsPublish  bool
	CategoryID uuid.UUID
	CreatedAt  time.Time
	BeginDate  time.Time
	EndDate    time.Time
}

func UpsertProductsToInternal(upsertProducts []*UpsertProductRequest) []*model.Product {
	products := make([]*model.Product, len(upsertProducts))

	for i, upsertProduct := range upsertProducts {
		internalProduct := upsertProduct.ToInternal()
		products[i] = &internalProduct
	}

	return products
}

func (d *UpsertProductRequest) ToInternal() model.Product {
	return model.Product{
		ID:        d.ID,
		Name:      d.Name,
		Price:     d.Price,
		IsPublish: d.IsPublish,
	}
}

func (svc *Service) ListProducts(
	param ListProductRequest,
) (
	[]*model.Product, error,
) {
	products, err := svc.db.GetProductsWithConditions(
		param.BeginDate, param.EndDate, model.Product{
			ID:         param.ProductID,
			Name:       param.Name,
			IsPublish:  param.IsPublish,
			CategoryID: param.CategoryID,
			CreatedAt:  param.CreatedAt,
		})

	if err != nil {
		return nil, err
	}

	return products, nil
}
