package app

import (
	"context"
	"fmt"
	collectionpb "github.com/mephistolie/chefbook-backend-collection/api/proto/implementation/v1"
	"github.com/mephistolie/chefbook-backend-collection/internal/config"
	"github.com/mephistolie/chefbook-backend-collection/internal/repository/postgres"
	"github.com/mephistolie/chefbook-backend-collection/internal/transport/amqp"
	"github.com/mephistolie/chefbook-backend-collection/internal/transport/dependencies/service"
	collection "github.com/mephistolie/chefbook-backend-collection/internal/transport/grpc"
	"github.com/mephistolie/chefbook-backend-common/log"
	"github.com/mephistolie/chefbook-backend-common/shutdown"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
	"net"
	"time"
)

func Run(cfg *config.Config) {
	log.Init(*cfg.LogsPath, *cfg.Environment == config.EnvDev)
	cfg.Print()

	db, err := postgres.Connect(cfg.Database)
	if err != nil {
		log.Fatal(err)
		return
	}

	repository := postgres.NewRepository(db)

	collectionService := service.New(repository)

	var mqServer *amqp.Server = nil
	if len(*cfg.Amqp.Host) > 0 {
		mqServer, err = amqp.NewServer(cfg.Amqp, collectionService.MQ)
		if err != nil {
			return
		}
		if err := mqServer.Start(); err != nil {
			log.Fatal(err)
			return
		}
		log.Info("MQ server initialized")
	}

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *cfg.Port))
	if err != nil {
		log.Fatal(err)
		return
	}

	grpcServer := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			log.UnaryServerInterceptor(),
		),
	)

	healthServer := health.NewServer()
	collectionServer := collection.NewServer(*collectionService)

	go monitorHealthChecking(db, healthServer)

	collectionpb.RegisterCollectionServiceServer(grpcServer, collectionServer)
	healthpb.RegisterHealthServer(grpcServer, healthServer)

	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			log.Errorf("error occurred while running http server: %s\n", err.Error())
		} else {
			log.Info("gRPC server started")
		}
	}()

	wait := shutdown.Graceful(context.Background(), 5*time.Second, map[string]shutdown.Operation{
		"grpc-server": func(ctx context.Context) error {
			grpcServer.GracefulStop()
			return nil
		},
		"database": func(ctx context.Context) error {
			return db.Close()
		},
		"mq": func(ctx context.Context) error {
			if mqServer == nil {
				return nil
			}
			return mqServer.Stop()
		},
	})
	<-wait
}
