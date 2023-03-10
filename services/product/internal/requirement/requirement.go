package requirement

import (
	"github.com/hyuti/clean-slice-template/services/product/ent"
	"github.com/hyuti/clean-slice-template/services/product/internal/infra/rpc"
	"github.com/hyuti/clean-slice-template/services/product/internal/requirement/features"
	prodRepo "github.com/hyuti/clean-slice-template/services/product/internal/requirement/product/repo"
	"github.com/kataras/iris/v12"
)

func New(h iris.Party, entC *ent.Client, r *rpc.GRPCServer) {
	// config features
	features.New(h, prodRepo.New(entC), r)
}
