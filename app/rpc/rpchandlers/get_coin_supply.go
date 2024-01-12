package rpchandlers

import (
	"github.com/SCR-NETWORK/SCR_Network/app/appmessage"
	"github.com/SCR-NETWORK/SCR_Network/app/rpc/rpccontext"
	"github.com/SCR-NETWORK/SCR_Network/domain/consensus/utils/constants"
	"github.com/SCR-NETWORK/SCR_Network/infrastructure/network/netadapter/router"
)

// HandleGetCoinSupply handles the respectively named RPC command
func HandleGetCoinSupply(context *rpccontext.Context, _ *router.Router, _ appmessage.Message) (appmessage.Message, error) {
	if !context.Config.UTXOIndex {
		errorMessage := &appmessage.GetCoinSupplyResponseMessage{}
<<<<<<< HEAD
		errorMessage.Error = appmessage.RPCErrorf("Method unavailable when SCR-NETWORK is run without --utxoindex")
=======
		errorMessage.Error = appmessage.RPCErrorf("Method unavailable when SCR_Network is run without --utxoindex")
>>>>>>> 0e8ed786da4e31df71edebffb326da64b0a6c3b4
		return errorMessage, nil
	}

	circulatingLeorSupply, err := context.UTXOIndex.GetCirculatingLeorSupply()
	if err != nil {
		return nil, err
	}

	response := appmessage.NewGetCoinSupplyResponseMessage(
		constants.MaxLeor,
		circulatingLeorSupply,
	)

	return response, nil
}
