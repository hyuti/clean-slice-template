package features

import (
	"github.com/hyuti/clean-slice-template/services/product/ent"
	"github.com/hyuti/clean-slice-template/services/product/internal/config/rpc"
	"github.com/kataras/iris/v12"
)

func New(h iris.Party, entC *ent.Client, r *rpc.GRPCServer) {
	// config prod repo
	prodRepo := newProductRepo(entC)

	// config prod endpoint
	newProductGroup(h, prodRepo, r)
}
