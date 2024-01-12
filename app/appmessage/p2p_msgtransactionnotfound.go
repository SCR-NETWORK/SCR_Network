// Copyright (c) 2013-2015 The btcsuite developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package appmessage

import (
	"github.com/SCR-NETWORK/SCR_Network/domain/consensus/model/externalapi"
)

<<<<<<< HEAD
// MsgTransactionNotFound defines a SCR TransactionNotFound message which is sent in response to
=======
// MsgTransactionNotFound defines a SCR TransactionNotFound message which is sent in response to
>>>>>>> 0e8ed786da4e31df71edebffb326da64b0a6c3b4
// a RequestTransactions message if any of the requested data in not available on the peer.
type MsgTransactionNotFound struct {
	baseMessage
	ID *externalapi.DomainTransactionID
}

// Command returns the protocol command string for the message. This is part
// of the Message interface implementation.
func (msg *MsgTransactionNotFound) Command() MessageCommand {
	return CmdTransactionNotFound
}

<<<<<<< HEAD
// NewMsgTransactionNotFound returns a new SCR transactionsnotfound message that conforms to the
=======
// NewMsgTransactionNotFound returns a new SCR transactionsnotfound message that conforms to the
>>>>>>> 0e8ed786da4e31df71edebffb326da64b0a6c3b4
// Message interface. See MsgTransactionNotFound for details.
func NewMsgTransactionNotFound(id *externalapi.DomainTransactionID) *MsgTransactionNotFound {
	return &MsgTransactionNotFound{
		ID: id,
	}
}
