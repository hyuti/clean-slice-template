package config

import (
	"fmt"

	dupCfg1 "github.com/hyuti/clean-slice-template/services/product/config"
	"github.com/hyuti/clean-slice-template/services/product/ent"
	"github.com/hyuti/clean-slice-template/services/product/internal/config/data"
	"github.com/hyuti/clean-slice-template/services/product/internal/config/rest"
	"github.com/hyuti/clean-slice-template/services/product/pkg/logger"
	"github.com/kataras/iris/v12"
)

// A struct holds all infrastructures of the application
type Infra struct {
	rest *iris.Application
	data *ent.Client
}

func (s *Infra) GetHTTPServerHandler() *iris.Application {
	return s.rest
}
func (s *Infra) GetDataClient() *ent.Client {
	return s.data
}

func New(cfg *dupCfg1.Config, l logger.Interface) App {
	data, err := data.New(&cfg.DB)
	if err != nil {
		l.Fatal(fmt.Errorf("%w", err))
	}
	return &Infra{
		rest: rest.New(),
		data: data,
	}
}
