package infra

import (
	"fmt"

	"github.com/hyuti/clean-slice-template/services/product/config"
	"github.com/hyuti/clean-slice-template/services/product/ent"
	"github.com/hyuti/clean-slice-template/services/product/internal/infra/data"
	"github.com/hyuti/clean-slice-template/services/product/internal/infra/rest"
	"github.com/hyuti/clean-slice-template/services/product/internal/infra/rpc"
	"github.com/hyuti/clean-slice-template/services/product/pkg/logger"
	"github.com/kataras/iris/v12"
)

// A struct holds all infrastructures of the application
type Infra struct {
	Rest       *iris.Application
	DataClient *ent.Client
	RPC        *rpc.GRPCServer
}

func New(cfg *config.Config, l logger.Interface) *Infra {
	data, err := data.New(&cfg.DB)
	if err != nil {
		l.Fatal(fmt.Errorf("%w", err))
	}
	return &Infra{
		Rest:       rest.New(),
		DataClient: data,
		RPC:        rpc.New(),
	}
}
