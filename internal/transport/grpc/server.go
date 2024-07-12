package grpc

import (
	api "github.com/mephistolie/chefbook-backend-collection/api/proto/implementation/v1"
	"github.com/mephistolie/chefbook-backend-collection/internal/transport/dependencies/service"
)

type CollectionServer struct {
	api.UnsafeCollectionServiceServer
	service service.Service
}

func NewServer(service service.Service) *CollectionServer {
	return &CollectionServer{service: service}
}
