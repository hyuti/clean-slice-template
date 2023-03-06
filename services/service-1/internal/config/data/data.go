package data

import (
	"github.com/hyuti/clean-slice-template/services/service-1/config"
	"github.com/hyuti/clean-slice-template/services/service-1/ent"
	"github.com/hyuti/clean-slice-template/services/service-1/pkg/datastore"
)

func New(cfg *config.DB) (*ent.Client, error) {
	return datastore.NewClient(cfg.URL, cfg.PoolMax)
}
