package rpc

import (
	"context"

	productv1 "github.com/hyuti/clean-slice-template/services/product/pkg/proto/gen/product/v1"
	"google.golang.org/grpc"
)

var _ productv1.IGRPCServer = (*GRPCServer)(nil)

type GRPCServer struct {
	productv1.UnimplementedProductServiceServer
	getProductByIDHandler func(context.Context, *productv1.GetProductByIDRequest) (*productv1.GetProductByIDResponse, error)
	srv                   *grpc.Server
}

func New() *GRPCServer {
	s := grpc.NewServer()
	return new(s)
}

func (s *GRPCServer) RegisterGetProductByID(h func(context.Context, *productv1.GetProductByIDRequest) (*productv1.GetProductByIDResponse, error)) {
	s.getProductByIDHandler = h
}
func (s *GRPCServer) GetProductByID(ctx context.Context, req *productv1.GetProductByIDRequest) (*productv1.GetProductByIDResponse, error) {
	return s.getProductByIDHandler(ctx, req)
}
func (s *GRPCServer) GetGRPCServer() *grpc.Server {
	return s.srv
}
func new(s *grpc.Server) *GRPCServer {
	return &GRPCServer{srv: s}
}
