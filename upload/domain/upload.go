package domain

import (
	"context"
	"fmt"
	"graph-analyzer/api/config"
	"graph-analyzer/api/upload/domain/internal"
	"graph-analyzer/api/upload/domain/pb"
	"io"
	"mime/multipart"
	"sync"

	"google.golang.org/grpc"
)

type UploadService interface {
	Invoke(file *multipart.FileHeader) (UploadDTO, error)
}

type upload struct {
	mutex sync.Mutex
}

func NewUploadService() UploadService {
	return &upload{}
}

type UploadDTO struct {
	Status bool
}

func (s *upload) Invoke(file *multipart.FileHeader) (UploadDTO, error) {
	// Only allow one concurrent upload
	// This limitation should later be implemented in the data collector using channels
	s.mutex.Lock()
	defer s.mutex.Unlock()

	fileContent, err := file.Open()
	if err != nil {
		return UploadDTO{}, err
	}
	defer func(fileContent multipart.File) {
		err := fileContent.Close()
		if err != nil {
			return
		}
	}(fileContent)

	byteContainer, err := io.ReadAll(fileContent)
	if err != nil {
		return UploadDTO{}, err
	}

	status, err := gexfService(&pb.GexfRequest{
		FileContent: byteContainer,
		NetworkName: file.Filename,
	})
	if err != nil {
		return UploadDTO{}, err
	}

	return UploadDTO{
		Status: status.Success,
	}, nil
}

func gexfService(request *pb.GexfRequest) (*pb.GexfResponse, error) {
	connection, err := internal.NewConnection(fmt.Sprintf("%s:%s", config.GetEnvVariable("GRPC_HOST"), config.GetEnvVariable("GRPC_PORT")))
	if err != nil {
		return nil, err
	}
	defer func(connection *grpc.ClientConn) {
		errClose := connection.Close()
		if errClose != nil {
			return
		}
	}(connection)

	client := pb.NewGexfServiceClient(connection)
	status, err := client.ProcessGexf(context.Background(), request)
	if err != nil {
		return nil, err
	}
	return status, nil
}
