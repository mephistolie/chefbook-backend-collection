package grpc

import (
	api "github.com/mephistolie/chefbook-backend-category/api/proto/implementation/v1"
	"github.com/mephistolie/chefbook-backend-category/internal/transport/dependencies/service"
)

type CategoryServer struct {
	api.UnsafeCategoryServiceServer
	service service.Service
}

func NewServer(service service.Service) *CategoryServer {
	return &CategoryServer{service: service}
}
