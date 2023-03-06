package dtos

import (
	"github.com/google/uuid"
)

type ProductResp struct {
	ID    uuid.UUID `json:"id"`
	Price float64   `json:"price"`
	Name  string    `json:"name"`
}
