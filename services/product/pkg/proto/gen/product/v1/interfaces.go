package productv1

import context "context"

type (
	IRegister interface {
		RegisterGetProductByID(func(context.Context, *GetProductByIDRequest) (*GetProductByIDResponse, error))
	}
	IGRPCServer interface {
		ProductServiceServer
		IRegister
	}
)
