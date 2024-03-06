package rpchandlers

import (
	"github.com/dmgcoin/dmgcoin/app/appmessage"
	"github.com/dmgcoin/dmgcoin/app/rpc/rpccontext"
	"github.com/dmgcoin/dmgcoin/infrastructure/network/netadapter/router"
)

// HandleGetCurrentNetwork handles the respectively named RPC command
func HandleGetCurrentNetwork(context *rpccontext.Context, _ *router.Router, _ appmessage.Message) (appmessage.Message, error) {
	response := appmessage.NewGetCurrentNetworkResponseMessage(context.Config.ActiveNetParams.Net.String())
	return response, nil
}
