package appmessage

import (
	"github.com/SCR-NETWORK/SCR_Network/domain/consensus/model/externalapi"
)

// MaxInvPerTxInvMsg is the maximum number of hashes that can
// be in a single CmdInvTransaction message.
const MaxInvPerTxInvMsg = MaxInvPerMsg

<<<<<<< HEAD
// MsgInvTransaction implements the Message interface and represents a SCR
=======
// MsgInvTransaction implements the Message interface and represents a SCR
>>>>>>> 0e8ed786da4e31df71edebffb326da64b0a6c3b4
// TxInv message. It is used to notify the network about new transactions
// by sending their ID, and let the receiving node decide if it needs it.
type MsgInvTransaction struct {
	baseMessage
	TxIDs []*externalapi.DomainTransactionID
}

// Command returns the protocol command string for the message. This is part
// of the Message interface implementation.
func (msg *MsgInvTransaction) Command() MessageCommand {
	return CmdInvTransaction
}

<<<<<<< HEAD
// NewMsgInvTransaction returns a new SCR TxInv message that conforms to
=======
// NewMsgInvTransaction returns a new SCR TxInv message that conforms to
>>>>>>> 0e8ed786da4e31df71edebffb326da64b0a6c3b4
// the Message interface. See MsgInvTransaction for details.
func NewMsgInvTransaction(ids []*externalapi.DomainTransactionID) *MsgInvTransaction {
	return &MsgInvTransaction{
		TxIDs: ids,
	}
}
