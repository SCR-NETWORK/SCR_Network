package appmessage

<<<<<<< HEAD
// MsgRequestPruningPointAndItsAnticone represents a SCR RequestPruningPointAndItsAnticone message
=======
// MsgRequestPruningPointAndItsAnticone represents a SCR RequestPruningPointAndItsAnticone message
>>>>>>> 0e8ed786da4e31df71edebffb326da64b0a6c3b4
type MsgRequestPruningPointAndItsAnticone struct {
	baseMessage
}

// Command returns the protocol command string for the message
func (msg *MsgRequestPruningPointAndItsAnticone) Command() MessageCommand {
	return CmdRequestPruningPointAndItsAnticone
}

// NewMsgRequestPruningPointAndItsAnticone returns a new MsgRequestPruningPointAndItsAnticone.
func NewMsgRequestPruningPointAndItsAnticone() *MsgRequestPruningPointAndItsAnticone {
	return &MsgRequestPruningPointAndItsAnticone{}
}
