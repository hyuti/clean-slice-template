package queries

import (
	"github.com/hyuti/clean-slice-template/services/product/internal/requirement/product/repo"
)

func New(r repo.IProductRepo) *Handler {
	return &Handler{
		repo: r,
	}
}
