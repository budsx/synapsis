package common

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/sony/gobreaker"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

func ReadyToTrip(counts gobreaker.Counts) bool {
	fmt.Printf("CB Request Counts: %v %v\n", counts.Requests, counts.TotalFailures)
	failration := float64(counts.TotalFailures) / float64(counts.Requests)
	return counts.Requests >= 100 && failration >= 0.6
}

func IsSuccess(err error) bool {
	if err == nil || strings.Contains(err.Error(), "desc = OK: HTTP status code 200") {
		return true
	}
	return false
}

func GrpcClientConnection(address string, circuitBreaker *gobreaker.CircuitBreaker) (*grpc.ClientConn, error) {
	var conn *grpc.ClientConn
	var err error
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithUnaryInterceptor(circuitBreakerInterceptor(circuitBreaker)),
	}

	conn, err = grpc.NewClient(address, opts...)
	if err != nil {
		return nil, err
	}
	return conn, nil
}

func circuitBreakerInterceptor(cb *gobreaker.CircuitBreaker) grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply any, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		_, err := cb.Execute(func() (interface{}, error) {
			return nil, invoker(ctx, method, req, reply, cc, opts...)
		})
		return err
	}
}

func SetupCircuitBreaker(timeout time.Duration, serviceName string) *gobreaker.CircuitBreaker {
	return gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name:         serviceName,
		MaxRequests:  10,
		Interval:     2 * timeout,
		Timeout:      timeout,
		ReadyToTrip:  ReadyToTrip,
		IsSuccessful: IsSuccess,
	})
}

func GetRequestHeaderByKey(ctx context.Context, key string) string {
	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		if keys := md.Get(key); len(keys) > 0 {
			return keys[0]
		}
	}
	return ""
}
