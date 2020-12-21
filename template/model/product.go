package model

import "github.com/google/uuid"

type Product struct {
	ID               uuid.UUID
	Name             string
	Price            int64
	IsPublish        bool
	ProviderUniqueID uuid.UUID
}
