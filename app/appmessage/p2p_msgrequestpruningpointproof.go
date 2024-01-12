package appmessage

<<<<<<< HEAD
// MsgRequestPruningPointProof represents a SCR RequestPruningPointProof message
=======
// MsgRequestPruningPointProof represents a SCR RequestPruningPointProof message
>>>>>>> 0e8ed786da4e31df71edebffb326da64b0a6c3b4
type MsgRequestPruningPointProof struct {
	baseMessage
}

// Command returns the protocol command string for the message
func (msg *MsgRequestPruningPointProof) Command() MessageCommand {
	return CmdRequestPruningPointProof
}

// NewMsgRequestPruningPointProof returns a new MsgRequestPruningPointProof.
func NewMsgRequestPruningPointProof() *MsgRequestPruningPointProof {
	return &MsgRequestPruningPointProof{}
}
