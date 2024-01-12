package appmessage

<<<<<<< HEAD
// MsgBlockWithTrustedDataV4 represents a SCR BlockWithTrustedDataV4 message
=======
// MsgBlockWithTrustedDataV4 represents a SCR BlockWithTrustedDataV4 message
>>>>>>> 0e8ed786da4e31df71edebffb326da64b0a6c3b4
type MsgBlockWithTrustedDataV4 struct {
	baseMessage

	Block               *MsgBlock
	DAAWindowIndices    []uint64
	GHOSTDAGDataIndices []uint64
}

// Command returns the protocol command string for the message
func (msg *MsgBlockWithTrustedDataV4) Command() MessageCommand {
	return CmdBlockWithTrustedDataV4
}

// NewMsgBlockWithTrustedDataV4 returns a new MsgBlockWithTrustedDataV4.
func NewMsgBlockWithTrustedDataV4() *MsgBlockWithTrustedDataV4 {
	return &MsgBlockWithTrustedDataV4{}
}
