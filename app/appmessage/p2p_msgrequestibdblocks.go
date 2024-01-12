package appmessage

import (
	"github.com/SCR-NETWORK/SCR_Network/domain/consensus/model/externalapi"
)

<<<<<<< HEAD
// MsgRequestIBDBlocks implements the Message interface and represents a SCR
=======
// MsgRequestIBDBlocks implements the Message interface and represents a SCR
>>>>>>> 0e8ed786da4e31df71edebffb326da64b0a6c3b4
// RequestIBDBlocks message. It is used to request blocks as part of the IBD
// protocol.
type MsgRequestIBDBlocks struct {
	baseMessage
	Hashes []*externalapi.DomainHash
}

// Command returns the protocol command string for the message. This is part
// of the Message interface implementation.
func (msg *MsgRequestIBDBlocks) Command() MessageCommand {
	return CmdRequestIBDBlocks
}

// NewMsgRequestIBDBlocks returns a new MsgRequestIBDBlocks.
func NewMsgRequestIBDBlocks(hashes []*externalapi.DomainHash) *MsgRequestIBDBlocks {
	return &MsgRequestIBDBlocks{
		Hashes: hashes,
	}
}
