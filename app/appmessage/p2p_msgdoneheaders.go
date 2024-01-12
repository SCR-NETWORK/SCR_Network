package appmessage

<<<<<<< HEAD
// MsgDoneHeaders implements the Message interface and represents a SCR
=======
// MsgDoneHeaders implements the Message interface and represents a SCR
>>>>>>> 0e8ed786da4e31df71edebffb326da64b0a6c3b4
// DoneHeaders message. It is used to notify the IBD syncing peer that the
// syncer sent all the requested headers.
//
// This message has no payload.
type MsgDoneHeaders struct {
	baseMessage
}

// Command returns the protocol command string for the message. This is part
// of the Message interface implementation.
func (msg *MsgDoneHeaders) Command() MessageCommand {
	return CmdDoneHeaders
}

<<<<<<< HEAD
// NewMsgDoneHeaders returns a new SCR DoneIBDBlocks message that conforms to the
=======
// NewMsgDoneHeaders returns a new SCR DoneIBDBlocks message that conforms to the
>>>>>>> 0e8ed786da4e31df71edebffb326da64b0a6c3b4
// Message interface.
func NewMsgDoneHeaders() *MsgDoneHeaders {
	return &MsgDoneHeaders{}
}
