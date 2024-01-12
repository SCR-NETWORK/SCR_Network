package rpchandlers

import (
	"github.com/SCR-NETWORK/SCR_Network/app/appmessage"
	"github.com/SCR-NETWORK/SCR_Network/app/rpc/rpccontext"
	"github.com/SCR-NETWORK/SCR_Network/infrastructure/network/netadapter/router"
)

// HandleNotifyPruningPointUTXOSetOverrideRequest handles the respectively named RPC command
func HandleNotifyPruningPointUTXOSetOverrideRequest(context *rpccontext.Context, router *router.Router, _ appmessage.Message) (appmessage.Message, error) {
	listener, err := context.NotificationManager.Listener(router)
	if err != nil {
		return nil, err
	}
	listener.PropagatePruningPointUTXOSetOverrideNotifications()

	response := appmessage.NewNotifyPruningPointUTXOSetOverrideResponseMessage()
	return response, nil
}
