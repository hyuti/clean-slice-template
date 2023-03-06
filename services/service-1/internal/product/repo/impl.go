package repo

import (
	"context"

	"github.com/hyuti/clean-slice-template/services/service-1/ent"
	"github.com/hyuti/clean-slice-template/services/service-1/ent/product"
	"github.com/hyuti/clean-slice-template/services/service-1/internal/product/models"
)

type ProductRepo struct {
	client *ent.Client
}

func (s *ProductRepo) Retrieve(ctx context.Context, o *models.ProductOrderInput, w *models.ProductWhereInput) (*models.Product, error) {
	// for sake of simplicity, ignore OrderInput for now
	query := s.client.Product.Query()

	if w.ID != nil {
		query = query.Where(product.ID(*w.ID))
	}
	return query.Only(ctx)
}
