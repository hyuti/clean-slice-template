package main

import (
	"log"

	"github.com/hyuti/clean-slice-template/services/product/config"
	"github.com/hyuti/clean-slice-template/services/product/internal"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	internal.New(cfg)
}
