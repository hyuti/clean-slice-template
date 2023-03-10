package features

import (
	"github.com/hyuti/clean-slice-template/services/product/internal/requirement/product/repo"
	productv1 "github.com/hyuti/clean-slice-template/services/product/pkg/proto/gen/product/v1"
	"github.com/kataras/iris/v12"
)

func New(h iris.Party, prodRepo repo.IProductRepo, r productv1.IRegister) {
	// config prod endpoint
	newProductGroup(h, prodRepo, r)
}
