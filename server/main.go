package main

import (
	"context"
	"fmt"
	at "github.com/abitofhelp/apis/proto/enum/access_tier/v5"
	pb "github.com/abitofhelp/apis/proto/hello_world/v1"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/timestamppb"
	"net"
	"time"
)

const (
	grpcPort = ":50051"
)

// server is used to implement hello_world.GreeterServer.
type server struct {
	pb.UnimplementedGreeterServer
	accessTier at.AccessTier
}

// SayHello implements hello_world.GreeterServer.
func (s *server) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	requestArrivedUtc := time.Now().UTC()
	logger.Sugar().Infof("Received: %v", req)
	return buildResponse(req.Name, requestArrivedUtc), nil
}

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	if lis, err := net.Listen("tcp", grpcPort); err == nil {
		svr := grpc.NewServer()
		pb.RegisterGreeterServer(svr, &server{})
		// Register reflection service on the gRPC server.
		reflection.Register(svr)

		// Start gRPC server.
		logger.Sugar().Infof(">>>> Starting gRPC service on %s...\n", lis.Addr().String())
		if err := svr.Serve(lis); err != nil {
			logger.Sugar().Fatalf("failed to serve: %v", err)
		}
	} else {
		logger.Sugar().Fatalf("failed to listen: %v", err)
	}
}

func buildResponse(name string, requestArrivedUtc time.Time) *pb.HelloResponse {
	messageArrivedTimestamp := timestamppb.Timestamp{
		Seconds: requestArrivedUtc.Unix(),
		Nanos:   int32(requestArrivedUtc.Nanosecond()),
	}

	repliedUtc := time.Now().UTC()
	repliedTimestamp := timestamppb.Timestamp{
		Seconds: repliedUtc.Unix(),
		Nanos:   int32(repliedUtc.Nanosecond()),
	}

	return &pb.HelloResponse{
		Message:           fmt.Sprintf("Howdy %s!", name),
		MessageArrivedUtc: &messageArrivedTimestamp,
		RepliedUtc:        &repliedTimestamp,
	}
}
