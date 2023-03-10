package data

import (
	"github.com/hyuti/clean-slice-template/services/product/config"
	"github.com/hyuti/clean-slice-template/services/product/ent"
	"github.com/hyuti/clean-slice-template/services/product/pkg/datastore"
)

func New(cfg *config.DB) (*ent.Client, error) {
	return datastore.NewClient(cfg.URL, cfg.PoolMax)
}
