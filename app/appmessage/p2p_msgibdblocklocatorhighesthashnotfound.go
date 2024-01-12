package appmessage

<<<<<<< HEAD
// MsgIBDBlockLocatorHighestHashNotFound represents a SCR BlockLocatorHighestHashNotFound message
=======
// MsgIBDBlockLocatorHighestHashNotFound represents a SCR BlockLocatorHighestHashNotFound message
>>>>>>> 0e8ed786da4e31df71edebffb326da64b0a6c3b4
type MsgIBDBlockLocatorHighestHashNotFound struct {
	baseMessage
}

// Command returns the protocol command string for the message
func (msg *MsgIBDBlockLocatorHighestHashNotFound) Command() MessageCommand {
	return CmdIBDBlockLocatorHighestHashNotFound
}

// NewMsgIBDBlockLocatorHighestHashNotFound returns a new IBDBlockLocatorHighestHashNotFound message
func NewMsgIBDBlockLocatorHighestHashNotFound() *MsgIBDBlockLocatorHighestHashNotFound {
	return &MsgIBDBlockLocatorHighestHashNotFound{}
}
