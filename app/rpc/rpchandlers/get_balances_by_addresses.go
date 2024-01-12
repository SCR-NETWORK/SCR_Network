package rpchandlers

import (
	"github.com/SCR-NETWORK/SCR_Network/app/appmessage"
	"github.com/SCR-NETWORK/SCR_Network/app/rpc/rpccontext"
	"github.com/SCR-NETWORK/SCR_Network/infrastructure/network/netadapter/router"
	"github.com/pkg/errors"
)

// HandleGetBalancesByAddresses handles the respectively named RPC command
func HandleGetBalancesByAddresses(context *rpccontext.Context, _ *router.Router, request appmessage.Message) (appmessage.Message, error) {
	if !context.Config.UTXOIndex {
		errorMessage := &appmessage.GetBalancesByAddressesResponseMessage{}
<<<<<<< HEAD
		errorMessage.Error = appmessage.RPCErrorf("Method unavailable when SCR-NETWORK is run without --utxoindex")
=======
		errorMessage.Error = appmessage.RPCErrorf("Method unavailable when SCR_Network is run without --utxoindex")
>>>>>>> 0e8ed786da4e31df71edebffb326da64b0a6c3b4
		return errorMessage, nil
	}

	getBalancesByAddressesRequest := request.(*appmessage.GetBalancesByAddressesRequestMessage)

	allEntries := make([]*appmessage.BalancesByAddressesEntry, len(getBalancesByAddressesRequest.Addresses))
	for i, address := range getBalancesByAddressesRequest.Addresses {
		balance, err := getBalanceByAddress(context, address)

		if err != nil {
			rpcError := &appmessage.RPCError{}
			if !errors.As(err, &rpcError) {
				return nil, err
			}
			errorMessage := &appmessage.GetUTXOsByAddressesResponseMessage{}
			errorMessage.Error = rpcError
			return errorMessage, nil
		}
		allEntries[i] = &appmessage.BalancesByAddressesEntry{
			Address: address,
			Balance: balance,
		}
	}

	response := appmessage.NewGetBalancesByAddressesResponse(allEntries)
	return response, nil
}
