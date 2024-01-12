package appmessage

import (
	"github.com/SCR-NETWORK/SCR_Network/domain/consensus/model/externalapi"
)

<<<<<<< HEAD
// MsgRequestPruningPointUTXOSet represents a SCR RequestPruningPointUTXOSet message
=======
// MsgRequestPruningPointUTXOSet represents a SCR RequestPruningPointUTXOSet message
>>>>>>> 0e8ed786da4e31df71edebffb326da64b0a6c3b4
type MsgRequestPruningPointUTXOSet struct {
	baseMessage
	PruningPointHash *externalapi.DomainHash
}

// Command returns the protocol command string for the message
func (msg *MsgRequestPruningPointUTXOSet) Command() MessageCommand {
	return CmdRequestPruningPointUTXOSet
}

// NewMsgRequestPruningPointUTXOSet returns a new MsgRequestPruningPointUTXOSet
func NewMsgRequestPruningPointUTXOSet(pruningPointHash *externalapi.DomainHash) *MsgRequestPruningPointUTXOSet {
	return &MsgRequestPruningPointUTXOSet{
		PruningPointHash: pruningPointHash,
	}
}
