package features

import (
	"github.com/kataras/iris/v12"

	getProductById "github.com/hyuti/clean-slice-template/services/product/internal/requirement/features/get-product-by-id"
	"github.com/hyuti/clean-slice-template/services/product/internal/requirement/product/repo"
	productv1 "github.com/hyuti/clean-slice-template/services/product/pkg/proto/gen/product/v1"
)

func newProductGroup(h iris.Party, repo repo.IProductRepo, rpc productv1.IRegister) {
	gr := h.Party("/products")
	getProductById.New(gr, repo, rpc)
}
