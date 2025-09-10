package main

import (
	"context"
	"log/slog"

	"github.com/budsx/synapsis/order-service/config"
	"github.com/budsx/synapsis/order-service/handler"
	"github.com/budsx/synapsis/order-service/repository"
	"github.com/budsx/synapsis/order-service/server"
	"github.com/budsx/synapsis/order-service/services"
)

func main() {
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
		MicroConf: repository.MicroConf{
			InventoryHost: conf.InventoryHost,
			InventoryPort: conf.InventoryPort,
		},
	})
	if err != nil {
		slog.Error("Failed to create repository", "error", err)
		return
	}
	service := services.NewOrderService(repo)
	handler := handler.NewOrderHandler(service)

	ctx := context.Background()
	slog.Info("🌐 Starting gRPC server...", "port", conf.GRPCPort)
	grpcServer, err := server.RunGRPCServer(conf, handler)
	if err != nil {
		slog.Error("Failed to create gRPC server", "error", err)
		return
	}
	slog.Info("✅ gRPC server started", "port", conf.GRPCPort)

	slog.Info("🌐 Starting REST gateway server...", "port", conf.RESTPort)
	go func() {
		err := server.RunGRPCGatewayServer(ctx, conf, handler)
		if err != nil {
			slog.Error("Failed to create gRPC gateway server", "error", err)
			return
		}
	}()

	slog.Info("Service orders is running...")
	server.GracefulShutdown(ctx, map[string]server.Operation{
		"grpc_server": func(ctx context.Context) error {
			grpcServer.GracefulStop()
			return nil
		},
		"repository": func(ctx context.Context) error {
			// repo.Close()
			return nil
		},
	})
}
