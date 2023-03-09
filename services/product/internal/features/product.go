package features

import (
	"github.com/kataras/iris/v12"

	"github.com/hyuti/clean-slice-template/services/product/ent"
	"github.com/hyuti/clean-slice-template/services/product/internal/config/rpc"
	getProductById "github.com/hyuti/clean-slice-template/services/product/internal/features/get-product-by-id"
	"github.com/hyuti/clean-slice-template/services/product/internal/product/repo"
)

func newProductRepo(entC *ent.Client) repo.IProductRepo {
	return repo.New(entC)
}

func newProductGroup(h iris.Party, repo repo.IProductRepo, rpc *rpc.GRPCServer) {
	gr := h.Party("/products")
	getProductById.New(gr, repo, rpc)
}
