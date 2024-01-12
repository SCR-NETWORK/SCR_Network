package appmessage

<<<<<<< HEAD
// MsgDonePruningPointUTXOSetChunks represents a SCR DonePruningPointUTXOSetChunks message
=======
// MsgDonePruningPointUTXOSetChunks represents a SCR DonePruningPointUTXOSetChunks message
>>>>>>> 0e8ed786da4e31df71edebffb326da64b0a6c3b4
type MsgDonePruningPointUTXOSetChunks struct {
	baseMessage
}

// Command returns the protocol command string for the message
func (msg *MsgDonePruningPointUTXOSetChunks) Command() MessageCommand {
	return CmdDonePruningPointUTXOSetChunks
}

// NewMsgDonePruningPointUTXOSetChunks returns a new MsgDonePruningPointUTXOSetChunks.
func NewMsgDonePruningPointUTXOSetChunks() *MsgDonePruningPointUTXOSetChunks {
	return &MsgDonePruningPointUTXOSetChunks{}
}
