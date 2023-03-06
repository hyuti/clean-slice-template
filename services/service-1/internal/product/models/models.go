package models

import (
	"github.com/google/uuid"
	"github.com/hyuti/clean-slice-template/services/service-1/ent"
)

type (
	Product           = ent.Product
	ProductWhereInput struct {
		ID *uuid.UUID
	}
	ProductOrderInput struct{}
)