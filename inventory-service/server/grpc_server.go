package server

import (
	"context"
	"fmt"
	"log/slog"
	"net"
	"os"
	"os/signal"
	"regexp"
	"runtime/debug"
	"sync"
	"syscall"

	"github.com/budsx/synapsis/inventory-service/config"
	"github.com/budsx/synapsis/inventory-service/handler"
	inventory "github.com/budsx/synapsis/inventory-service/proto"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/recovery"
	"go.elastic.co/apm/module/apmgrpc/v2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
)

func RunGRPCServer(conf *config.Config, handler *handler.InventoryHandler) (*grpc.Server, error) {
	grpcConn, err := net.Listen("tcp", fmt.Sprintf(":%d", conf.GRPCPort))
	if err != nil {
		slog.Error("Failed to listen", "error", err)
		return nil, err
	}

	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(interceptor),
		grpc.ChainUnaryInterceptor(
			apmgrpc.NewUnaryServerInterceptor(),
			recovery.UnaryServerInterceptor(recovery.WithRecoveryHandlerContext(grpcRecoveryHandler)),
		),
		grpc.ChainStreamInterceptor(
			apmgrpc.NewStreamServerInterceptor(),
			recovery.StreamServerInterceptor(recovery.WithRecoveryHandlerContext(grpcRecoveryHandler)),
		),
	)

	inventory.RegisterInventoryServiceServer(grpcServer, handler)
	reflection.Register(grpcServer)

	go grpcServer.Serve(grpcConn)
	return grpcServer, nil
}

func grpcRecoveryHandler(ctx context.Context, panic interface{}) error {
	newLineRegex := regexp.MustCompile(`\r?\n`)
	stackTrace := newLineRegex.ReplaceAllString(string(debug.Stack()), " ")
	slog.Error("panic happened",
		"panic_message", panic,
		"panic_stacktrace", stackTrace)
	return status.Errorf(codes.Internal, "server error happened")
}

func interceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	newCtx := metadata.NewIncomingContext(ctx, md)
	return handler(newCtx, req)
}

type Operation func(context.Context) error

func GracefulShutdown(ctx context.Context, ops map[string]Operation) <-chan struct{} {
	s := make(chan os.Signal, 1)
	signal.Notify(s, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	<-s

	wait := make(chan struct{})
	var wg sync.WaitGroup

	for _, op := range ops {
		wg.Add(1)
		go func(op Operation) {
			defer wg.Done()
			if err := op(ctx); err != nil {
				slog.Error("Failed to perform operation", "error", err)
			}
		}(op)
	}
	wg.Wait()
	close(wait)

	return wait
}
