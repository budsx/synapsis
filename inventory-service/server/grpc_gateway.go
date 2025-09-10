package server

import (
	"context"
	"fmt"
	"net/http"

	"github.com/budsx/synapsis/inventory-service/config"
	"github.com/budsx/synapsis/inventory-service/handler"
	inventory "github.com/budsx/synapsis/inventory-service/proto"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func RunGRPCGatewayServer(ctx context.Context, conf *config.Config, handler *handler.InventoryHandler) error {
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}
	err := inventory.RegisterInventoryServiceHandlerFromEndpoint(ctx, mux, fmt.Sprintf(":%d", conf.GRPCPort), opts)
	if err != nil {
		return fmt.Errorf("failed to register inventory service handler: %v", err)
	}

	if err := http.ListenAndServe(fmt.Sprintf(":%d", conf.RESTPort), mux); err != nil {
		return fmt.Errorf("failed to serve: %v", err)
	}
	return nil
}
