package config

import (
	"github.com/hyuti/clean-slice-template/services/service-1/ent"
	"github.com/kataras/iris/v12"
)

type App interface {
	GetHTTPServerHandler() *iris.Application
	GetDataClient() *ent.Client
}
