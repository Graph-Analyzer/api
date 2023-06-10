package domain

import (
	"context"
	"fmt"
	"graph-analyzer/api/config"
	"graph-analyzer/api/upload/domain/internal"
	"graph-analyzer/api/upload/domain/pb"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

type UploadStatusService interface {
	Invoke() (UploadStatusDTO, error)
}

type uploadStatus struct {
}

func NewUploadStatusService() UploadStatusService {
	return &uploadStatus{}
}

type UploadStatusDTO struct {
	Healthy bool
}

func (s *uploadStatus) Invoke() (UploadStatusDTO, error) {
	health, err := healthService()
	if err != nil {
		return UploadStatusDTO{
			Healthy: false,
		}, err
	}

	return UploadStatusDTO{
		Healthy: health.Healthy,
	}, nil
}

func healthService() (*pb.HealthCheckResponse, error) {
	connection, err := internal.NewConnection(fmt.Sprintf("%s:%s", config.GetEnvVariable("GRPC_HOST"), config.GetEnvVariable("GRPC_PORT")))
	if err != nil {
		return nil, err
	}
	defer func(connection *grpc.ClientConn) {
		err := connection.Close()
		if err != nil {
			return
		}
	}(connection)

	client := pb.NewHealthCheckServiceClient(connection)
	health, err := client.Check(context.Background(), &emptypb.Empty{})
	if err != nil {
		return nil, err
	}
	return health, nil
}
