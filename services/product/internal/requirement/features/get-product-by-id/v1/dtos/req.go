package dtos

import (
	"github.com/google/uuid"
	irisUtils "github.com/hyuti/clean-slice-template/services/product/pkg/utils/iris"
)

type (
	DetailReq = irisUtils.DetailReq
)

func NewDetailReq(id uuid.UUID) *DetailReq {
	return &DetailReq{
		ID: id,
	}
}
