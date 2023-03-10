package queries

import (
	"context"

	"github.com/hyuti/clean-slice-template/services/product/internal/requirement/features/get-product-by-id/v1/dtos"
	"github.com/hyuti/clean-slice-template/services/product/internal/requirement/product/models"
	"github.com/hyuti/clean-slice-template/services/product/internal/requirement/product/repo"
)

type Handler struct {
	repo repo.IProductRepo
}

func (s *Handler) Handle(ctx context.Context, req *dtos.DetailReq) (*dtos.ProductResp, error) {
	e, err := s.repo.Retrieve(ctx, nil, &models.ProductWhereInput{
		ID: &req.ID,
	})
	return &dtos.ProductResp{
		ID:    e.ID,
		Name:  e.Name,
		Price: e.Price,
	}, err
}
