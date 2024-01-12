package appmessage

import (
	"github.com/SCR-NETWORK/SCR_Network/domain/consensus/model/externalapi"
)

<<<<<<< HEAD
// MsgIBDBlockLocator represents a SCR ibdBlockLocator message
=======
// MsgIBDBlockLocator represents a SCR ibdBlockLocator message
>>>>>>> 0e8ed786da4e31df71edebffb326da64b0a6c3b4
type MsgIBDBlockLocator struct {
	baseMessage
	TargetHash         *externalapi.DomainHash
	BlockLocatorHashes []*externalapi.DomainHash
}

// Command returns the protocol command string for the message
func (msg *MsgIBDBlockLocator) Command() MessageCommand {
	return CmdIBDBlockLocator
}

<<<<<<< HEAD
// NewMsgIBDBlockLocator returns a new SCR ibdBlockLocator message
=======
// NewMsgIBDBlockLocator returns a new SCR ibdBlockLocator message
>>>>>>> 0e8ed786da4e31df71edebffb326da64b0a6c3b4
func NewMsgIBDBlockLocator(targetHash *externalapi.DomainHash,
	blockLocatorHashes []*externalapi.DomainHash) *MsgIBDBlockLocator {

	return &MsgIBDBlockLocator{
		TargetHash:         targetHash,
		BlockLocatorHashes: blockLocatorHashes,
	}
}
