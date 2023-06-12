package main

import (
	"context"
	"fmt"
	"net"
	"time"

	"grpc_example/pkg/logger"
	"grpc_example/pkg/pb/service"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const grpcPort = 50051

type authServer struct {
	service.UnimplementedAuthServiceServer
}

func (as *authServer) Authenticate(ctx context.Context, userRequest *service.UserRequest) (*service.ServiceResponse, error) {
	var (
		maxValidId    uint64 = 10000
		validAuthCode uint32 = 4956
		timeAfter            = time.Date(2000, 01, 01, 21, 00, 00, 00, time.UTC)
	)

	logger.Logger.Sugar().Infof("New request, id: %d", userRequest.GetId())

	if userRequest.GetId() > maxValidId {
		return &service.ServiceResponse{
			Status:      false,
			HelpMessage: "Invalid user ID",
		}, nil
	}

	if userRequest.GetAuthCode() != validAuthCode {
		return &service.ServiceResponse{
			Status:      false,
			HelpMessage: "Invalid auth code",
		}, nil
	}

	if !userRequest.GetRequestTime().AsTime().After(timeAfter) {
		return &service.ServiceResponse{
			Status:      false,
			HelpMessage: "Invalid time was given",
		}, nil
	}

	return &service.ServiceResponse{
		Status:      true,
		HelpMessage: "Ok",
	}, nil
}

func main() {
	err := logger.SetupLogger()
	if err != nil {
		panic(err)
	}
	defer func() { _ = logger.Logger.Sync() }()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		logger.Logger.Sugar().Fatal(err)
	}

	serv := grpc.NewServer()
	reflection.Register(serv)
	service.RegisterAuthServiceServer(serv, &authServer{})

	logger.Logger.Sugar().Infof("server listening at %s", lis.Addr())

	if err := serv.Serve(lis); err != nil {
		logger.Logger.Sugar().Fatalf("error while starting server: %s\n", err)
	}
}
