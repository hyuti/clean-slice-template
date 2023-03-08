package internal

import (
	"fmt"

	dupCfg1 "github.com/hyuti/clean-slice-template/services/product/config"
	"github.com/hyuti/clean-slice-template/services/product/internal/config"
	"github.com/hyuti/clean-slice-template/services/product/internal/features"
	"github.com/hyuti/clean-slice-template/services/product/pkg/logger"
)

func New(cfg *dupCfg1.Config) {
	l := logger.New("debug")

	app := config.New(cfg, l)
	defer app.GetDataClient().Close()

	features.New(app.GetHTTPServerHandler(), app.GetDataClient())

	app.GetHTTPServerHandler().Listen(fmt.Sprintf(":%s", cfg.HTTP.Port))
}
