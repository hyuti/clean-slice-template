package iris

import (
	"github.com/google/uuid"
	ir "github.com/kataras/iris/v12"
)

type DetailReq struct {
	ID uuid.UUID `param:"id"`
}

func ReadID(ctx ir.Context, req *DetailReq) error {
	id := ctx.Params().Get("id")
	uid, err := uuid.Parse(id)
	if err != nil {
		return err
	}
	req.ID = uid
	return nil
}
