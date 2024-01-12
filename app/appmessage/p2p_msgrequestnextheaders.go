package appmessage

<<<<<<< HEAD
// MsgRequestNextHeaders implements the Message interface and represents a SCR
=======
// MsgRequestNextHeaders implements the Message interface and represents a SCR
>>>>>>> 0e8ed786da4e31df71edebffb326da64b0a6c3b4
// RequestNextHeaders message. It is used to notify the IBD syncer peer to send
// more headers.
//
// This message has no payload.
type MsgRequestNextHeaders struct {
	baseMessage
}

// Command returns the protocol command string for the message. This is part
// of the Message interface implementation.
func (msg *MsgRequestNextHeaders) Command() MessageCommand {
	return CmdRequestNextHeaders
}

<<<<<<< HEAD
// NewMsgRequestNextHeaders returns a new SCR RequestNextHeaders message that conforms to the
=======
// NewMsgRequestNextHeaders returns a new SCR RequestNextHeaders message that conforms to the
>>>>>>> 0e8ed786da4e31df71edebffb326da64b0a6c3b4
// Message interface.
func NewMsgRequestNextHeaders() *MsgRequestNextHeaders {
	return &MsgRequestNextHeaders{}
}
