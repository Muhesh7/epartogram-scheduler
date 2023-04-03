package middlewares

import (
	"context"
	"fmt"
	"time"

	"google.golang.org/grpc"
)

func WithClientUnaryInterceptor() grpc.DialOption {
	return grpc.WithUnaryInterceptor(clientInterceptor)
}

func clientInterceptor(
	ctx context.Context,
	method string,
	req interface{},
	reply interface{},
	cc *grpc.ClientConn,
	invoker grpc.UnaryInvoker,
	opts ...grpc.CallOption,
) error {
	var err error
	start := time.Now()
	// Calls the invoker to execute RPC
	err = invoker(ctx, method, req, reply, cc)
	// Logic after invoking the invoker
	fmt.Printf("Invoked RPC method=%s; Duration=%s; Error=%v", method,
		time.Since(start), err)
	return err
}
