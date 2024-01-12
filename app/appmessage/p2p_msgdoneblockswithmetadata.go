package appmessage

<<<<<<< HEAD
// MsgDoneBlocksWithTrustedData implements the Message interface and represents a SCR
=======
// MsgDoneBlocksWithTrustedData implements the Message interface and represents a SCR
>>>>>>> 0e8ed786da4e31df71edebffb326da64b0a6c3b4
// DoneBlocksWithTrustedData message
//
// This message has no payload.
type MsgDoneBlocksWithTrustedData struct {
	baseMessage
}

// Command returns the protocol command string for the message. This is part
// of the Message interface implementation.
func (msg *MsgDoneBlocksWithTrustedData) Command() MessageCommand {
	return CmdDoneBlocksWithTrustedData
}

<<<<<<< HEAD
// NewMsgDoneBlocksWithTrustedData returns a new SCR DoneBlocksWithTrustedData message that conforms to the
=======
// NewMsgDoneBlocksWithTrustedData returns a new SCR DoneBlocksWithTrustedData message that conforms to the
>>>>>>> 0e8ed786da4e31df71edebffb326da64b0a6c3b4
// Message interface.
func NewMsgDoneBlocksWithTrustedData() *MsgDoneBlocksWithTrustedData {
	return &MsgDoneBlocksWithTrustedData{}
}
