package appmessage

<<<<<<< HEAD
// MsgPruningPoints represents a SCR PruningPoints message
=======
// MsgPruningPoints represents a SCR PruningPoints message
>>>>>>> 0e8ed786da4e31df71edebffb326da64b0a6c3b4
type MsgPruningPoints struct {
	baseMessage

	Headers []*MsgBlockHeader
}

// Command returns the protocol command string for the message
func (msg *MsgPruningPoints) Command() MessageCommand {
	return CmdPruningPoints
}

// NewMsgPruningPoints returns a new MsgPruningPoints.
func NewMsgPruningPoints(headers []*MsgBlockHeader) *MsgPruningPoints {
	return &MsgPruningPoints{
		Headers: headers,
	}
}
