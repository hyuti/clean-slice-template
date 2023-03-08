package endpoints

import (
	"github.com/hyuti/clean-slice-template/services/product/internal/features/get-product-by-id/v1/dtos"
	irisUtils "github.com/hyuti/clean-slice-template/services/product/pkg/utils/iris"
	"github.com/kataras/iris/v12"
	"github.com/mehdihadeli/go-mediatr"
)

func Register(gr iris.Party) {
	gr.Get("/{id:string}", func(ctx iris.Context) {
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
}
