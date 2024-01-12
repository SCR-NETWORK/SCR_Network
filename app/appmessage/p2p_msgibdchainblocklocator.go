package appmessage

import (
	"github.com/SCR-NETWORK/SCR_Network/domain/consensus/model/externalapi"
)

<<<<<<< HEAD
// MsgIBDChainBlockLocator implements the Message interface and represents a SCR
=======
// MsgIBDChainBlockLocator implements the Message interface and represents a SCR
>>>>>>> 0e8ed786da4e31df71edebffb326da64b0a6c3b4
// locator message. It is used to find the blockLocator of a peer that is
// syncing with you.
type MsgIBDChainBlockLocator struct {
	baseMessage
	BlockLocatorHashes []*externalapi.DomainHash
}

// Command returns the protocol command string for the message. This is part
// of the Message interface implementation.
func (msg *MsgIBDChainBlockLocator) Command() MessageCommand {
	return CmdIBDChainBlockLocator
}

<<<<<<< HEAD
// NewMsgIBDChainBlockLocator returns a new SCR locator message that conforms to
=======
// NewMsgIBDChainBlockLocator returns a new SCR locator message that conforms to
>>>>>>> 0e8ed786da4e31df71edebffb326da64b0a6c3b4
// the Message interface. See MsgBlockLocator for details.
func NewMsgIBDChainBlockLocator(locatorHashes []*externalapi.DomainHash) *MsgIBDChainBlockLocator {
	return &MsgIBDChainBlockLocator{
		BlockLocatorHashes: locatorHashes,
	}
}
