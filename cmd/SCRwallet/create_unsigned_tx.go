package main

import (
	"context"
	"fmt"
	"os"

	"github.com/SCR-NETWORK/SCR_Network/cmd/SCRwallet/daemon/client"
	"github.com/SCR-NETWORK/SCR_Network/cmd/SCRwallet/daemon/pb"
	"github.com/SCR-NETWORK/SCR_Network/domain/consensus/utils/constants"
)

func createUnsignedTransaction(conf *createUnsignedTransactionConfig) error {
	daemonClient, tearDown, err := client.Connect(conf.DaemonAddress)
	if err != nil {
		return err
	}
	defer tearDown()

	ctx, cancel := context.WithTimeout(context.Background(), daemonTimeout)
	defer cancel()

	sendAmountLeor := uint64(conf.SendAmount * constants.LeorPerSCR)
	response, err := daemonClient.CreateUnsignedTransactions(ctx, &pb.CreateUnsignedTransactionsRequest{
		From:                     conf.FromAddresses,
		Address:                  conf.ToAddress,
		Amount:                   sendAmountLeor,
		IsSendAll:                conf.IsSendAll,
		UseExistingChangeAddress: conf.UseExistingChangeAddress,
	})
	if err != nil {
		return err
	}

	fmt.Fprintln(os.Stderr, "Created unsigned transaction")
	fmt.Println(encodeTransactionsToHex(response.UnsignedTransactions))

	return nil
}
