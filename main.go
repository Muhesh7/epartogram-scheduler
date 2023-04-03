package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/go-co-op/gocron"
	"github.com/joho/godotenv"
	"github.com/muhesh7/epartogram-scheduler/middlewares"
	"github.com/muhesh7/epartogram-scheduler/rpcs"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func check(client rpcs.MonitorServiceClient) (*rpcs.Empty, error) {
	return client.Check(context.Background(), &rpcs.Empty{})
}

func runCronJob(client rpcs.MonitorServiceClient) {
	s := gocron.NewScheduler(time.UTC)

	s.Every(30).Minutes().Do(func() {
		go check(client)
	})

	s.StartBlocking()
}

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println(fmt.Println("Error loading .env"))
	}
	var gRPCHost = os.Getenv("GRPC_HOST")

	var opts []grpc.DialOption
	opts = append(
		opts,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		middlewares.WithClientUnaryInterceptor())
	conn, err := grpc.Dial(gRPCHost, opts...)

	if err != nil {
		fmt.Println("error in dial", err)
	}
	defer conn.Close()
	client := rpcs.NewMonitorServiceClient(conn)
	runCronJob(client)
}
