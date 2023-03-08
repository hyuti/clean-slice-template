package features

import (
	"github.com/kataras/iris/v12"

	"github.com/hyuti/clean-slice-template/services/product/ent"
	getProductById "github.com/hyuti/clean-slice-template/services/product/internal/features/get-product-by-id"
	"github.com/hyuti/clean-slice-template/services/product/internal/product/repo"
)

func newProductRepo(entC *ent.Client) repo.IProductRepo {
	return repo.New(entC)
}

func newProductGroup(h iris.Party, r repo.IProductRepo) {
	gr := h.Party("/products")
	getProductById.New(gr, r)
}
