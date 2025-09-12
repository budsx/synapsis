package server

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/budsx/synapsis/inventory-service/config"
	inventory "github.com/budsx/synapsis/inventory-service/proto"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func RunGRPCGatewayServer(ctx context.Context, conf *config.Config) error {
	mux := runtime.NewServeMux(
		runtime.WithIncomingHeaderMatcher(customHeaderMatcher),
		runtime.WithOutgoingHeaderMatcher(runtime.DefaultHeaderMatcher),
	)

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

func customHeaderMatcher(key string) (string, bool) {
	key = strings.ToLower(key)
	switch key {
	case "x-idempotency-key":
		return "x-idempotency-key", true
	case "x-request-id":
		return "x-request-id", true
	default:
		return runtime.DefaultHeaderMatcher(key)
	}
}

func customOutgoingHeaderMatcher(key string) (string, bool) {
	key = strings.ToLower(key)
	switch key {
	case "x-idempotency-key":
		return "x-idempotency-key", true
	default:
		return runtime.DefaultHeaderMatcher(key)
	}
}
