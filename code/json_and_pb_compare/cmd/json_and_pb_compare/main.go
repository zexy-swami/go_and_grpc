package main

import (
	"encoding/json"
	"time"

	"json_and_pb_compare/pkg/logger"
	"json_and_pb_compare/pkg/pb/user"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func main() {
	err := logger.SetupLogger()
	if err != nil {
		panic(err)
	}
	defer func() { _ = logger.Logger.Sync() }()

	userInfo := &user.User{
		Id:             1,
		Name:           "Rob Pike",
		UsedAuthCodes:  []uint32{0000, 0500, 1000, 1500, 2000, 2500, 3000, 3500, 4000, 4500, 5000, 5500, 6000, 6500, 7000, 7500, 8000, 8500, 9000},
		LastTimeActive: timestamppb.New(time.Date(2020, 1, 10, 23, 30, 00, 36, time.UTC)),
	}

	showJSONStatistics(userInfo)
	showPBStatistics(userInfo)
}

func showJSONStatistics(userInfo *user.User) {
	jsonData, err := json.Marshal(userInfo)
	if err != nil {
		logger.Logger.Sugar().Fatal(err)
	}

	logger.Logger.Sugar().Infof("JSON size: %d", len(jsonData))
	logger.Logger.Sugar().Infof("JSON data: %+v", jsonData)
}

func showPBStatistics(userInfo *user.User) {
	pbData, err := proto.Marshal(userInfo)
	if err != nil {
		logger.Logger.Sugar().Fatal(err)
	}

	logger.Logger.Sugar().Infof("protobuf size: %d", len(pbData))
	logger.Logger.Sugar().Infof("protobuf data: %+v", pbData)
}
