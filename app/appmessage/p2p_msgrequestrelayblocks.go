package appmessage

import (
	"github.com/SCR-NETWORK/SCR_Network/domain/consensus/model/externalapi"
)

// MaxRequestRelayBlocksHashes is the maximum number of hashes that can
// be in a single RequestRelayBlocks message.
const MaxRequestRelayBlocksHashes = MaxInvPerMsg

<<<<<<< HEAD
// MsgRequestRelayBlocks implements the Message interface and represents a SCR
=======
// MsgRequestRelayBlocks implements the Message interface and represents a SCR
>>>>>>> 0e8ed786da4e31df71edebffb326da64b0a6c3b4
// RequestRelayBlocks message. It is used to request blocks as part of the block
// relay protocol.
type MsgRequestRelayBlocks struct {
	baseMessage
	Hashes []*externalapi.DomainHash
}

// Command returns the protocol command string for the message. This is part
// of the Message interface implementation.
func (msg *MsgRequestRelayBlocks) Command() MessageCommand {
	return CmdRequestRelayBlocks
}

<<<<<<< HEAD
// NewMsgRequestRelayBlocks returns a new SCR RequestRelayBlocks message that conforms to
=======
// NewMsgRequestRelayBlocks returns a new SCR RequestRelayBlocks message that conforms to
>>>>>>> 0e8ed786da4e31df71edebffb326da64b0a6c3b4
// the Message interface. See MsgRequestRelayBlocks for details.
func NewMsgRequestRelayBlocks(hashes []*externalapi.DomainHash) *MsgRequestRelayBlocks {
	return &MsgRequestRelayBlocks{
		Hashes: hashes,
	}
}
