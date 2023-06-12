package main

import (
	"context"
	"math/rand"
	"time"

	"grpc_example/pkg/logger"
	"grpc_example/pkg/pb/service"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/timestamppb"
)

const serverAddr = "localhost:50051"

func main() {
	err := logger.SetupLogger()
	if err != nil {
		panic(err)
	}
	defer func() { _ = logger.Logger.Sync() }()

	conn, err := grpc.Dial(serverAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.Logger.Sugar().Fatalf("error while creating connection: %s\n", err)
	}
	defer conn.Close()

	client := service.NewAuthServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	rand.Seed(time.Now().UnixNano())
	id := uint64(rand.Intn(1001))

	response, err := client.Authenticate(ctx, &service.UserRequest{
		Id:          id,
		AuthCode:    4956,
		RequestTime: timestamppb.New(time.Date(2001, 01, 01, 21, 00, 00, 00, time.UTC)),
	})

	if err != nil {
		logger.Logger.Sugar().Panicf("error occurred while making auth, err: %s", err)
	}

	logger.Logger.Sugar().Infof("status: %t, message: %s", response.Status, response.HelpMessage)
}
