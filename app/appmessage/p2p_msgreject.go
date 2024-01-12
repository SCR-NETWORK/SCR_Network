package appmessage

<<<<<<< HEAD
// MsgReject implements the Message interface and represents a SCR
=======
// MsgReject implements the Message interface and represents a SCR
>>>>>>> 0e8ed786da4e31df71edebffb326da64b0a6c3b4
// Reject message. It is used to notify peers why they are banned.
type MsgReject struct {
	baseMessage
	Reason string
}

// Command returns the protocol command string for the message. This is part
// of the Message interface implementation.
func (msg *MsgReject) Command() MessageCommand {
	return CmdReject
}

<<<<<<< HEAD
// NewMsgReject returns a new SCR Reject message that conforms to the
=======
// NewMsgReject returns a new SCR Reject message that conforms to the
>>>>>>> 0e8ed786da4e31df71edebffb326da64b0a6c3b4
// Message interface.
func NewMsgReject(reason string) *MsgReject {
	return &MsgReject{
		Reason: reason,
	}
}
