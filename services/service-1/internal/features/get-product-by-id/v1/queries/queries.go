package queries

import (
	"github.com/hyuti/clean-slice-template/services/service-1/internal/product/repo"
)

func New(r repo.IProductRepo) *Handler {
	return &Handler{
		repo: r,
	}
}
