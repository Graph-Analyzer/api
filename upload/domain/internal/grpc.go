package internal

import (
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/keepalive"
)

func NewConnection(url string) (*grpc.ClientConn, error) {
	kacp := keepalive.ClientParameters{
		Time:    10 * time.Second, // send ping every 10 seconds if there is no activity
		Timeout: time.Second,      // wait 1 second for ping back
	}

	connection, err := grpc.Dial(url,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithKeepaliveParams(kacp))
	if err != nil {
		return nil, err
	}

	return connection, nil
}
