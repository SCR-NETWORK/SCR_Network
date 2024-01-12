package rpchandlers

import (
	"github.com/SCR-NETWORK/SCR_Network/app/appmessage"
	"github.com/SCR-NETWORK/SCR_Network/app/rpc/rpccontext"
	"github.com/SCR-NETWORK/SCR_Network/infrastructure/network/netadapter/router"
)

// HandleGetHeaders handles the respectively named RPC command
func HandleGetHeaders(context *rpccontext.Context, _ *router.Router, request appmessage.Message) (appmessage.Message, error) {
	response := &appmessage.GetHeadersResponseMessage{}
	response.Error = appmessage.RPCErrorf("not implemented")
	return response, nil
}
