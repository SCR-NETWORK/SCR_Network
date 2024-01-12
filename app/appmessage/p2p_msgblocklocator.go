package appmessage

import (
	"github.com/SCR-NETWORK/SCR_Network/domain/consensus/model/externalapi"
)

// MaxBlockLocatorsPerMsg is the maximum number of block locator hashes allowed
// per message.
const MaxBlockLocatorsPerMsg = 500

<<<<<<< HEAD
// MsgBlockLocator implements the Message interface and represents a SCR
=======
// MsgBlockLocator implements the Message interface and represents a SCR
>>>>>>> 0e8ed786da4e31df71edebffb326da64b0a6c3b4
// locator message. It is used to find the blockLocator of a peer that is
// syncing with you.
type MsgBlockLocator struct {
	baseMessage
	BlockLocatorHashes []*externalapi.DomainHash
}

// Command returns the protocol command string for the message. This is part
// of the Message interface implementation.
func (msg *MsgBlockLocator) Command() MessageCommand {
	return CmdBlockLocator
}

<<<<<<< HEAD
// NewMsgBlockLocator returns a new SCR locator message that conforms to
=======
// NewMsgBlockLocator returns a new SCR locator message that conforms to
>>>>>>> 0e8ed786da4e31df71edebffb326da64b0a6c3b4
// the Message interface. See MsgBlockLocator for details.
func NewMsgBlockLocator(locatorHashes []*externalapi.DomainHash) *MsgBlockLocator {
	return &MsgBlockLocator{
		BlockLocatorHashes: locatorHashes,
	}
}
