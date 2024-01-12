package appmessage

<<<<<<< HEAD
// BlockHeadersMessage represents a SCR BlockHeaders message
=======
// BlockHeadersMessage represents a SCR BlockHeaders message
>>>>>>> 0e8ed786da4e31df71edebffb326da64b0a6c3b4
type BlockHeadersMessage struct {
	baseMessage
	BlockHeaders []*MsgBlockHeader
}

// Command returns the protocol command string for the message
func (msg *BlockHeadersMessage) Command() MessageCommand {
	return CmdBlockHeaders
}

<<<<<<< HEAD
// NewBlockHeadersMessage returns a new SCR BlockHeaders message
=======
// NewBlockHeadersMessage returns a new SCR BlockHeaders message
>>>>>>> 0e8ed786da4e31df71edebffb326da64b0a6c3b4
func NewBlockHeadersMessage(blockHeaders []*MsgBlockHeader) *BlockHeadersMessage {
	return &BlockHeadersMessage{
		BlockHeaders: blockHeaders,
	}
}
