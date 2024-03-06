package client

import (
	"context"
	"github.com/dmgcoin/dmgcoin/cmd/dmgcoinwallet/daemon/server"
	"time"

	"github.com/pkg/errors"

	"github.com/dmgcoin/dmgcoin/cmd/dmgcoinwallet/daemon/pb"
	"google.golang.org/grpc"
)

// Connect connects to the dmgcoinwalletd server, and returns the client instance
func Connect(address string) (pb.DmgcoinwalletdClient, func(), error) {
	// Connection is local, so 1 second timeout is sufficient
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	conn, err := grpc.DialContext(ctx, address, grpc.WithInsecure(), grpc.WithBlock(), grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(server.MaxDaemonSendMsgSize)))
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			return nil, nil, errors.New("dmgcoinwallet daemon is not running, start it with `dmgcoinwallet start-daemon`")
		}
		return nil, nil, err
	}

	return pb.NewDmgcoinwalletdClient(conn), func() {
		conn.Close()
	}, nil
}
