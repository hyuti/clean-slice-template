package repo

import (
	"github.com/hyuti/clean-slice-template/services/service-1/internal/product/models"
	"github.com/hyuti/clean-slice-template/services/service-1/pkg/crud"
)

type (
	IProductRepo interface {
		crud.RetrieveModelRepository[*models.Product, *models.ProductOrderInput, *models.ProductWhereInput]
	}
)
