package getProductById

import (
	"github.com/hyuti/clean-slice-template/services/product/internal/features/get-product-by-id/v1/controller"
	"github.com/hyuti/clean-slice-template/services/product/internal/features/get-product-by-id/v1/mediator"
	"github.com/hyuti/clean-slice-template/services/product/internal/product/repo"
	productv1 "github.com/hyuti/clean-slice-template/services/product/pkg/proto/gen/product/v1"
	"github.com/kataras/iris/v12"
)

func New(prodGr iris.Party, r repo.IProductRepo, rpc productv1.IRegister) {
	// config controller
	controller.Register(prodGr, rpc)

	// config mediator
	mediator.Register(r)
}
