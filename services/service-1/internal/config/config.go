package config

import (
	"fmt"

	dupCfg1 "github.com/hyuti/clean-slice-template/services/service-1/config"
	"github.com/hyuti/clean-slice-template/services/service-1/ent"
	"github.com/hyuti/clean-slice-template/services/service-1/internal/config/data"
	"github.com/hyuti/clean-slice-template/services/service-1/internal/config/rest"
	"github.com/hyuti/clean-slice-template/services/service-1/pkg/logger"
	"github.com/kataras/iris/v12"
)

// A struct hold all infrastructures of the application
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
