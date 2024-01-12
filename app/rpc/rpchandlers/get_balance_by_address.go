package rpchandlers

import (
	"github.com/SCR-NETWORK/SCR_Network/app/appmessage"
	"github.com/SCR-NETWORK/SCR_Network/app/rpc/rpccontext"
	"github.com/SCR-NETWORK/SCR_Network/domain/consensus/utils/txscript"
	"github.com/SCR-NETWORK/SCR_Network/infrastructure/network/netadapter/router"
	"github.com/SCR-NETWORK/SCR_Network/util"
	"github.com/pkg/errors"
)

// HandleGetBalanceByAddress handles the respectively named RPC command
func HandleGetBalanceByAddress(context *rpccontext.Context, _ *router.Router, request appmessage.Message) (appmessage.Message, error) {
	if !context.Config.UTXOIndex {
		errorMessage := &appmessage.GetUTXOsByAddressesResponseMessage{}
<<<<<<< HEAD
		errorMessage.Error = appmessage.RPCErrorf("Method unavailable when SCR-NETWORK is run without --utxoindex")
=======
		errorMessage.Error = appmessage.RPCErrorf("Method unavailable when SCR_Network is run without --utxoindex")
>>>>>>> 0e8ed786da4e31df71edebffb326da64b0a6c3b4
		return errorMessage, nil
	}

	getBalanceByAddressRequest := request.(*appmessage.GetBalanceByAddressRequestMessage)

	balance, err := getBalanceByAddress(context, getBalanceByAddressRequest.Address)
	if err != nil {
		rpcError := &appmessage.RPCError{}
		if !errors.As(err, &rpcError) {
			return nil, err
		}
		errorMessage := &appmessage.GetUTXOsByAddressesResponseMessage{}
		errorMessage.Error = rpcError
		return errorMessage, nil
	}

	response := appmessage.NewGetBalanceByAddressResponse(balance)
	return response, nil
}

func getBalanceByAddress(context *rpccontext.Context, addressString string) (uint64, error) {
	address, err := util.DecodeAddress(addressString, context.Config.ActiveNetParams.Prefix)
	if err != nil {
		return 0, appmessage.RPCErrorf("Couldn't decode address '%s': %s", addressString, err)
	}

	scriptPublicKey, err := txscript.PayToAddrScript(address)
	if err != nil {
		return 0, appmessage.RPCErrorf("Could not create a scriptPublicKey for address '%s': %s", addressString, err)
	}
	utxoOutpointEntryPairs, err := context.UTXOIndex.UTXOs(scriptPublicKey)
	if err != nil {
		return 0, err
	}

	balance := uint64(0)
	for _, utxoOutpointEntryPair := range utxoOutpointEntryPairs {
		balance += utxoOutpointEntryPair.Amount()
	}
	return balance, nil
}
