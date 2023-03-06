package features

import (
	"github.com/hyuti/clean-slice-template/services/service-1/ent"
	"github.com/kataras/iris/v12"
)

func New(h iris.Party, entC *ent.Client) {
	// config prod repo
	prodRepo := newProductRepo(entC)

	// config prod endpoint
	newProductGroup(h, prodRepo)
}
