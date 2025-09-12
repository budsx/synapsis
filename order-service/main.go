package main

import (
	"context"

	"github.com/budsx/synapsis/order-service/config"
	"github.com/budsx/synapsis/order-service/handler"
	"github.com/budsx/synapsis/order-service/repository"
	"github.com/budsx/synapsis/order-service/server"
	"github.com/budsx/synapsis/order-service/services"
	"github.com/budsx/synapsis/order-service/transport/messaging"
	"github.com/budsx/synapsis/order-service/utils/common"
)

func main() {
	conf := config.Load()
	logger := common.NewLogger()
	repo, err := initRepository(conf)
	ctx := context.Background()
	if err != nil {
		logger.Error(ctx, "Failed to create repository", "error", err)
		return
	}
	service := services.NewOrderService(repo, logger)
	handler := handler.NewOrderHandler(service)

	grpcServer, err := server.RunGRPCServer(conf, handler)
	if err != nil {
		logger.Error(ctx, "Failed to create gRPC server", "error", err)
		return
	}

	go func() {
		err := server.RunGRPCGatewayServer(ctx, conf)
		if err != nil {
			logger.Error(ctx, "Failed to create gRPC gateway server", "error", err)
			return
		}
	}()

	if err := messaging.NewTransportOrderMessaging(conf, service, repo.MessageQueue); err != nil {
		logger.Error(ctx, "Failed to create transport order messaging", "error", err)
		return
	}

	logger.Info(ctx, "Service orders is running...")
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

func initRepository(conf *config.Config) (*repository.Repository, error) {
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
			RabbitmqURL:       conf.Rabbitmq.RabbitmqURL,
			TopicReserveStock: conf.TopicReserveStock,
		},
		MicroConf: repository.MicroConf{
			InventoryHost: conf.InventoryHost,
			InventoryPort: conf.InventoryPort,
		},
		RedisConf: repository.RedisConf{
			RedisHost:     conf.Redis.RedisHost,
			RedisPassword: conf.Redis.RedisPassword,
			RedisDB:       conf.Redis.RedisDB,
		},
	})
	return repo, err
}
