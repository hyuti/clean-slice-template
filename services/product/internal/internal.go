package internal

import (
	"fmt"
	"net"

	dupCfg1 "github.com/hyuti/clean-slice-template/services/product/config"
	"github.com/hyuti/clean-slice-template/services/product/internal/config"
	"github.com/hyuti/clean-slice-template/services/product/internal/features"
	"github.com/hyuti/clean-slice-template/services/product/pkg/logger"
	productv1 "github.com/hyuti/clean-slice-template/services/product/pkg/proto/gen/product/v1"
)

func New(cfg *dupCfg1.Config) {
	l := logger.New("debug")

	app := config.New(cfg, l)
	defer app.DataClient.Close()

	features.New(app.Rest, app.DataClient, app.RPC)
	productv1.RegisterProductServiceServer(app.RPC.GetGRPCServer(), app.RPC)
	app.Rest.Listen(fmt.Sprintf(":%s", cfg.HTTP.Port))

	go func() {
		defer app.RPC.GetGRPCServer().GracefulStop()
	}()
	address := fmt.Sprintf("localhost:%s", cfg.HTTP.Port)
	network := "tcp"

	n, err := net.Listen(network, address)
	if err != nil {
		l.Fatal("failed to listen %v", err)
	}
	err = app.RPC.GetGRPCServer().Serve(n)
	if err != nil {
		l.Fatal("failed to start rpc server %v", err)
	}
}
