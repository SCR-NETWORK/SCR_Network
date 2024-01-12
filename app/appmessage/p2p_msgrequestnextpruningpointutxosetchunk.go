package appmessage

<<<<<<< HEAD
// MsgRequestNextPruningPointUTXOSetChunk represents a SCR RequestNextPruningPointUTXOSetChunk message
=======
// MsgRequestNextPruningPointUTXOSetChunk represents a SCR RequestNextPruningPointUTXOSetChunk message
>>>>>>> 0e8ed786da4e31df71edebffb326da64b0a6c3b4
type MsgRequestNextPruningPointUTXOSetChunk struct {
	baseMessage
}

// Command returns the protocol command string for the message
func (msg *MsgRequestNextPruningPointUTXOSetChunk) Command() MessageCommand {
	return CmdRequestNextPruningPointUTXOSetChunk
}

// NewMsgRequestNextPruningPointUTXOSetChunk returns a new MsgRequestNextPruningPointUTXOSetChunk.
func NewMsgRequestNextPruningPointUTXOSetChunk() *MsgRequestNextPruningPointUTXOSetChunk {
	return &MsgRequestNextPruningPointUTXOSetChunk{}
}
