package getProductById

import (
	"github.com/hyuti/clean-slice-template/services/service-1/internal/features/get-product-by-id/v1/endpoints"
	"github.com/hyuti/clean-slice-template/services/service-1/internal/features/get-product-by-id/v1/mediator"
	"github.com/hyuti/clean-slice-template/services/service-1/internal/product/repo"
	"github.com/kataras/iris/v12"
)

func New(prodGr iris.Party, r repo.IProductRepo) {
	// config endpoint
	endpoints.Register(prodGr)

	// config mediator
	mediator.Register(r)
}
