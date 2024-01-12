package appmessage

<<<<<<< HEAD
// MsgUnexpectedPruningPoint represents a SCR UnexpectedPruningPoint message
=======
// MsgUnexpectedPruningPoint represents a SCR UnexpectedPruningPoint message
>>>>>>> 0e8ed786da4e31df71edebffb326da64b0a6c3b4
type MsgUnexpectedPruningPoint struct {
	baseMessage
}

// Command returns the protocol command string for the message
func (msg *MsgUnexpectedPruningPoint) Command() MessageCommand {
	return CmdUnexpectedPruningPoint
}

<<<<<<< HEAD
// NewMsgUnexpectedPruningPoint returns a new SCR UnexpectedPruningPoint message
=======
// NewMsgUnexpectedPruningPoint returns a new SCR UnexpectedPruningPoint message
>>>>>>> 0e8ed786da4e31df71edebffb326da64b0a6c3b4
func NewMsgUnexpectedPruningPoint() *MsgUnexpectedPruningPoint {
	return &MsgUnexpectedPruningPoint{}
}
