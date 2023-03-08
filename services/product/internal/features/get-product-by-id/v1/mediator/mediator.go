package mediator

import (
	"github.com/hyuti/clean-slice-template/services/product/internal/features/get-product-by-id/v1/dtos"
	"github.com/hyuti/clean-slice-template/services/product/internal/features/get-product-by-id/v1/queries"
	"github.com/hyuti/clean-slice-template/services/product/internal/product/repo"
	"github.com/mehdihadeli/go-mediatr"
)

func Register(
	r repo.IProductRepo,
) error {
	err := mediatr.RegisterRequestHandler[*dtos.DetailReq, *dtos.ProductResp](queries.New(r))
	if err != nil {
		return err
	}
	return nil
}
