package appmessage

<<<<<<< HEAD
// MsgPruningPointProof represents a SCR PruningPointProof message
=======
// MsgPruningPointProof represents a SCR PruningPointProof message
>>>>>>> 0e8ed786da4e31df71edebffb326da64b0a6c3b4
type MsgPruningPointProof struct {
	baseMessage

	Headers [][]*MsgBlockHeader
}

// Command returns the protocol command string for the message
func (msg *MsgPruningPointProof) Command() MessageCommand {
	return CmdPruningPointProof
}

// NewMsgPruningPointProof returns a new MsgPruningPointProof.
func NewMsgPruningPointProof(headers [][]*MsgBlockHeader) *MsgPruningPointProof {
	return &MsgPruningPointProof{
		Headers: headers,
	}
}
