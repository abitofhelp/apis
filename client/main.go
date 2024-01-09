package main

import (
	"context"
	"fmt"
	pb "github.com/abitofhelp/apis/proto/hello_world/v1"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

const (
	address     = "localhost:50051"
	defaultName = "world"
)

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	// Set up a connection to the server.
	if conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials())); err == nil {
		defer conn.Close()

		client := pb.NewGreeterClient(conn)

		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		// Build the message to send.
		req := buildRequest(defaultName)
		if rsp, err := client.SayHello(ctx, &req); err == nil {
			now := time.Now().UTC()
			requestDuration := rsp.MessageArrivedUtc.AsTime().Sub(req.SentUtc.AsTime())
			responseDuration := now.Sub(rsp.RepliedUtc.AsTime())
			totalDuration := now.Sub(req.SentUtc.AsTime())

			fmt.Printf("%s %s\n %s: %dμs\n %s: %dμs\n %s: %dμs\n",
				rsp.Message,
				"Here are some details about your request:",
				"Request Duration",
				requestDuration.Microseconds(),
				"Response Duration",
				responseDuration.Microseconds(),
				"Total Duration",
				totalDuration.Microseconds(),
			)
		} else {
			logger.Sugar().Fatalf("failed to say hello: %v", err)
		}
	} else {
		logger.Sugar().Fatalf("did not connect: %v", err)
	}
}

func buildRequest(name string) pb.HelloRequest {
	now := time.Now().UTC()
	timestamp := timestamppb.Timestamp{
		Seconds: now.Unix(),
		Nanos:   int32(now.Nanosecond()),
	}

	return pb.HelloRequest{
		Name:    name,
		SentUtc: &timestamp,
	}
}
