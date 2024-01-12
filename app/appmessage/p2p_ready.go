package appmessage

<<<<<<< HEAD
// MsgReady implements the Message interface and represents a SCR
=======
// MsgReady implements the Message interface and represents a SCR
>>>>>>> 0e8ed786da4e31df71edebffb326da64b0a6c3b4
// Ready message. It is used to notify that the peer is ready to receive
// messages.
//
// This message has no payload.
type MsgReady struct {
	baseMessage
}

// Command returns the protocol command string for the message. This is part
// of the Message interface implementation.
func (msg *MsgReady) Command() MessageCommand {
	return CmdReady
}

<<<<<<< HEAD
// NewMsgReady returns a new SCR Ready message that conforms to the
=======
// NewMsgReady returns a new SCR Ready message that conforms to the
>>>>>>> 0e8ed786da4e31df71edebffb326da64b0a6c3b4
// Message interface.
func NewMsgReady() *MsgReady {
	return &MsgReady{}
}
