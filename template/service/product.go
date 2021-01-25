package service

import (
	"github.com/google/uuid"
	"github.com/pjchender/go-snippets/template/model"
)

// UpsertProductRequest 會從 JSON 來
type UpsertProductRequest struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Price     int64     `json:"price"`
	IsPublish bool      `json:"isPublish"`
}

func ToInternalProducts(upsertProducts []*UpsertProductRequest) []*model.Product {
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
