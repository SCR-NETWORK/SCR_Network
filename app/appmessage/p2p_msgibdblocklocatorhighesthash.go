package appmessage

import (
	"github.com/SCR-NETWORK/SCR_Network/domain/consensus/model/externalapi"
)

<<<<<<< HEAD
// MsgIBDBlockLocatorHighestHash represents a SCR BlockLocatorHighestHash message
=======
// MsgIBDBlockLocatorHighestHash represents a SCR BlockLocatorHighestHash message
>>>>>>> 0e8ed786da4e31df71edebffb326da64b0a6c3b4
type MsgIBDBlockLocatorHighestHash struct {
	baseMessage
	HighestHash *externalapi.DomainHash
}

// Command returns the protocol command string for the message
func (msg *MsgIBDBlockLocatorHighestHash) Command() MessageCommand {
	return CmdIBDBlockLocatorHighestHash
}

// NewMsgIBDBlockLocatorHighestHash returns a new BlockLocatorHighestHash message
func NewMsgIBDBlockLocatorHighestHash(highestHash *externalapi.DomainHash) *MsgIBDBlockLocatorHighestHash {
	return &MsgIBDBlockLocatorHighestHash{
		HighestHash: highestHash,
	}
}
