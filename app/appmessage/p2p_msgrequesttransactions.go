package appmessage

import (
	"github.com/SCR-NETWORK/SCR_Network/domain/consensus/model/externalapi"
)

// MaxInvPerRequestTransactionsMsg is the maximum number of hashes that can
// be in a single CmdInvTransaction message.
const MaxInvPerRequestTransactionsMsg = MaxInvPerMsg

<<<<<<< HEAD
// MsgRequestTransactions implements the Message interface and represents a SCR
=======
// MsgRequestTransactions implements the Message interface and represents a SCR
>>>>>>> 0e8ed786da4e31df71edebffb326da64b0a6c3b4
// RequestTransactions message. It is used to request transactions as part of the
// transactions relay protocol.
type MsgRequestTransactions struct {
	baseMessage
	IDs []*externalapi.DomainTransactionID
}

// Command returns the protocol command string for the message. This is part
// of the Message interface implementation.
func (msg *MsgRequestTransactions) Command() MessageCommand {
	return CmdRequestTransactions
}

<<<<<<< HEAD
// NewMsgRequestTransactions returns a new SCR RequestTransactions message that conforms to
=======
// NewMsgRequestTransactions returns a new SCR RequestTransactions message that conforms to
>>>>>>> 0e8ed786da4e31df71edebffb326da64b0a6c3b4
// the Message interface. See MsgRequestTransactions for details.
func NewMsgRequestTransactions(ids []*externalapi.DomainTransactionID) *MsgRequestTransactions {
	return &MsgRequestTransactions{
		IDs: ids,
	}
}
