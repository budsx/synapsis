package main

import (
	"context"
	"log/slog"

	"github.com/budsx/synapsis/inventory-service/config"
	"github.com/budsx/synapsis/inventory-service/handler"
	"github.com/budsx/synapsis/inventory-service/repository"
	"github.com/budsx/synapsis/inventory-service/server"
	"github.com/budsx/synapsis/inventory-service/services"
	"github.com/budsx/synapsis/inventory-service/transport/messaging"
	"github.com/budsx/synapsis/inventory-service/utils/globalvar"
	"go.elastic.co/apm/v2"
)

func main() {
	globalvar.APMTracer = apm.DefaultTracer()
	conf := config.Load()

	repo, err := repository.NewRepository(repository.RepoConf{
		DBConf: repository.DBConf{
			DBHost:     conf.Database.Host,
			DBPort:     conf.Database.Port,
			DBUser:     conf.Database.Username,
			DBPassword: conf.Database.Password,
			DBName:     conf.Database.DbName,
			DBDriver:   conf.Database.DriverName,
		},
		RabbitmqConf: repository.RabbitmqConf{
			RabbitmqURL: conf.Rabbitmq.RabbitmqURL,
		},
	})
	if err != nil {
		slog.Error("Failed to create repository", "error", err)
		return
	}
	service := services.NewInventoryService(repo)
	handler := handler.NewInventoryHandler(service)

	ctx := context.Background()
	grpcServer, err := server.RunGRPCServer(conf, handler)
	if err != nil {
		slog.Error("Failed to create gRPC server", "error", err)
		return
	}
	go func() {
		err := server.RunGRPCGatewayServer(ctx, conf, handler)
		if err != nil {
			slog.Error("Failed to create gRPC gateway server", "error", err)
			return
		}
	}()

	if err := messaging.NewMessagingListener(conf, service, repo.MessageQueue); err != nil {
		slog.Error("Failed to create messaging listener", "error", err)
		return
	}

	slog.Info("Service inventory is running...")
	server.GracefulShutdown(ctx, map[string]server.Operation{
		"grpc_server": func(ctx context.Context) error {
			grpcServer.GracefulStop()
			return nil
		},
		"repository": func(ctx context.Context) error {
			repo.Close()
			return nil
		},
	})
}
