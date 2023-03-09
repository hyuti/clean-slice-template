package controller

import (
	"context"

	"github.com/google/uuid"
	"github.com/hyuti/clean-slice-template/services/product/internal/features/get-product-by-id/v1/dtos"
	productv1 "github.com/hyuti/clean-slice-template/services/product/pkg/proto/gen/product/v1"
	irisUtils "github.com/hyuti/clean-slice-template/services/product/pkg/utils/iris"
	"github.com/kataras/iris/v12"
	"github.com/mehdihadeli/go-mediatr"
)

func Register(rest iris.Party, rpc productv1.IRegister) {
	rest.Get("/{id:string}", func(ctx iris.Context) {
		req := new(dtos.DetailReq)
		if err := irisUtils.ReadID(ctx, req); err != nil {
			ctx.Problem(iris.NewProblem().Type("/id").Title("Validation problem").Detail("ID param missing").Status(iris.StatusBadRequest))
			return
		}
		resp, err := mediatr.Send[*dtos.DetailReq, *dtos.ProductResp](ctx.Request().Context(), dtos.NewDetailReq(req.ID))
		if err != nil {
			ctx.Problem(iris.NewProblem().Type("/id").Title("Internal problem").Detail(err.Error()).Status(iris.StatusInternalServerError))
			return
		}
		ctx.JSON(resp)
	})

	rpc.RegisterGetProductByID(func(ctx context.Context, req *productv1.GetProductByIDRequest) (*productv1.GetProductByIDResponse, error) {
		id, err := uuid.Parse(req.Id)
		if err != nil {
			return nil, err
		}
		resp, err := mediatr.Send[*dtos.DetailReq, *dtos.ProductResp](ctx, dtos.NewDetailReq(id))
		if err != nil {
			return nil, err
		}
		return &productv1.GetProductByIDResponse{
			Id:    resp.ID.String(),
			Name:  resp.Name,
			Price: resp.Price,
		}, nil
	})
}
