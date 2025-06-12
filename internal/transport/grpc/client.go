package transportgrpc

import (
	"fmt"

	userpb "github.com/IbadT/project-protos/proto/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewUserClient(addr string) (userpb.UserServiceClient, *grpc.ClientConn, error) {
	// conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, nil, fmt.Errorf("failed to connect to %s: %w", addr, err)
	}

	client := userpb.NewUserServiceClient(conn)
	return client, conn, nil
}
