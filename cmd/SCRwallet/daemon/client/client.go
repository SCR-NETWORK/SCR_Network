package client

import (
	"context"
	"time"

	"github.com/SCR-NETWORK/SCR_Network/cmd/SCRwallet/daemon/server"

	"github.com/pkg/errors"

	"github.com/SCR-NETWORK/SCR_Network/cmd/SCRwallet/daemon/pb"
	"google.golang.org/grpc"
)

// Connect connects to the SCRwalletd server, and returns the client instance
func Connect(address string) (pb.SCRwalletdClient, func(), error) {
	// Connection is local, so 1 second timeout is sufficient
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	conn, err := grpc.DialContext(ctx, address, grpc.WithInsecure(), grpc.WithBlock(), grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(server.MaxDaemonSendMsgSize)))
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			return nil, nil, errors.New("SCRwallet daemon is not running, start it with `SCRwallet start-daemon`")
		}
		return nil, nil, err
	}

	return pb.NewSCRwalletdClient(conn), func() {
		conn.Close()
	}, nil
}
