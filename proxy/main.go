package main

import (
	"context"
	"fmt"
	pb "github.com/abitofhelp/apis/proto/hello_world/v1"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"net"
	"net/http"
)

const (
	grpcPort = ":50051"
	grpcHost = ""
	httpPort = ":50052"
	httpHost = ""
)

func main() {

	logger, _ := zap.NewProduction()
	defer logger.Sync()

	// Register gRPC server endpoint.
	// Make sure the gRPC server is running properly and is accessible.
	mux, err := buildMux()
	if err != nil {
		logger.Sugar().Fatalf("\nfailed to build the mux for the proxy service: %v", err)
	}

	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	}

	// The RegisterXXXXHandlerFromEndpoint() adds http handlers for annotated rpc calls in the .proto file.
	// It creates a gRPC client instance to forward http requests to the gRPC server.
	// The Context that is passed into this function is ignored.  Each http handler uses
	// the request's context and adds cancellation to it.  No connection timeout can be
	// enabled with this generated code unless it has been set in the request context in
	// an interceptor.  At this time, we don't have a strong requirement for connection
	// timeout, so I will pass in the background context.
	//ctx, cancel := context.WithCancel(context.Background())
	//ctx, cancel = context.WithTimeout(ctx, appcfg.Runtime.ConnectionTimeOut)
	//defer cancel()
	if err := pb.RegisterGreeterHandlerFromEndpoint(
		context.Background(),
		mux,
		fmt.Sprintf("%s%s", grpcHost, grpcPort), opts); err != nil {
		logger.Sugar().Fatal(err)
	}

	if lis, err := net.Listen("tcp", httpPort); err == nil {

		// Start HTTP server (and proxy calls to gRPC server endpoint)
		logger.Sugar().Infof(">>>> Starting HTTP service proxy on %s...\n", lis.Addr().String())

		if err := http.Serve(lis, mux); err != nil {
			logger.Sugar().Fatalf("failed to serve: %v", err)
		}
	} else {
		logger.Sugar().Fatalf("failed to listen: %v", err)
	}
}

func buildMux() (*runtime.ServeMux, error) {
	mux := runtime.NewServeMux(
		// To properly unescape a path containing a uri to a hello_worlds, which contains slashes (http://).
		runtime.WithUnescapingMode(runtime.UnescapingModeAllExceptReserved))

	return mux, nil
}
