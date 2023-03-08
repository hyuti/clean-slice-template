package repo

import (
	"github.com/hyuti/clean-slice-template/services/product/internal/product/models"
	"github.com/hyuti/clean-slice-template/services/product/pkg/crud"
)

type (
	IProductRepo interface {
		crud.RetrieveModelRepository[*models.Product, *models.ProductOrderInput, *models.ProductWhereInput]
	}
)
