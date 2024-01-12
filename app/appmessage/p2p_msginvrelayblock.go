package appmessage

import (
	"github.com/SCR-NETWORK/SCR_Network/domain/consensus/model/externalapi"
)

<<<<<<< HEAD
// MsgInvRelayBlock implements the Message interface and represents a SCR
=======
// MsgInvRelayBlock implements the Message interface and represents a SCR
>>>>>>> 0e8ed786da4e31df71edebffb326da64b0a6c3b4
// block inventory message. It is used to notify the network about new block
// by sending their hash, and let the receiving node decide if it needs it.
type MsgInvRelayBlock struct {
	baseMessage
	Hash *externalapi.DomainHash
}

// Command returns the protocol command string for the message. This is part
// of the Message interface implementation.
func (msg *MsgInvRelayBlock) Command() MessageCommand {
	return CmdInvRelayBlock
}

<<<<<<< HEAD
// NewMsgInvBlock returns a new SCR invrelblk message that conforms to
=======
// NewMsgInvBlock returns a new SCR invrelblk message that conforms to
>>>>>>> 0e8ed786da4e31df71edebffb326da64b0a6c3b4
// the Message interface. See MsgInvRelayBlock for details.
func NewMsgInvBlock(hash *externalapi.DomainHash) *MsgInvRelayBlock {
	return &MsgInvRelayBlock{
		Hash: hash,
	}
}
