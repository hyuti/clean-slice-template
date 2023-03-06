package crud

import "context"

type (
	RetrieveModelRepository[ModelType, OrderInputType, WhereInputType any] interface {
		Retrieve(context.Context, OrderInputType, WhereInputType) (ModelType, error)
	}
	ListModelRepository[ModelType, OrderInputType, WhereInputType any] interface {
		List(context.Context, *int, *int, OrderInputType, WhereInputType) ([]ModelType, error)
	}
	CreateModelRepository[ModelType, CreateInputType any] interface {
		Create(context.Context, CreateInputType) (ModelType, error)
	}
	DeleteModelRepository[ModelType any] interface {
		Delete(context.Context, ModelType) error
	}
	UpdateModelRepository[ModelType any, UpdateInputType any] interface {
		Update(context.Context, ModelType, UpdateInputType) (ModelType, error)
	}
	IIsNextModelRepository[ModelType, OrderInputType, WhereInputType any] interface {
		IsNext(context.Context, int, int, OrderInputType, WhereInputType) (bool, error)
	}
)
