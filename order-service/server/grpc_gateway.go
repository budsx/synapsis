package server

import (
	"context"
	"fmt"
	"net/http"

	"github.com/budsx/synapsis/order-service/config"
	"github.com/budsx/synapsis/order-service/handler"
	order "github.com/budsx/synapsis/order-service/proto"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func RunGRPCGatewayServer(ctx context.Context, conf *config.Config, handler *handler.OrderHandler) error {
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}
	err := order.RegisterOrderServiceHandlerFromEndpoint(ctx, mux, fmt.Sprintf(":%d", conf.GRPCPort), opts)
	if err != nil {
		return fmt.Errorf("failed to register order service handler: %v", err)
	}

	if err := http.ListenAndServe(fmt.Sprintf(":%d", conf.RESTPort), mux); err != nil {
		return fmt.Errorf("failed to serve: %v", err)
	}
	return nil
}
