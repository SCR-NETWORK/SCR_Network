package appmessage

<<<<<<< HEAD
// MsgRequestNextPruningPointAndItsAnticoneBlocks implements the Message interface and represents a SCR
=======
// MsgRequestNextPruningPointAndItsAnticoneBlocks implements the Message interface and represents a SCR
>>>>>>> 0e8ed786da4e31df71edebffb326da64b0a6c3b4
// RequestNextPruningPointAndItsAnticoneBlocks message. It is used to notify the IBD syncer peer to send
// more blocks from the pruning anticone.
//
// This message has no payload.
type MsgRequestNextPruningPointAndItsAnticoneBlocks struct {
	baseMessage
}

// Command returns the protocol command string for the message. This is part
// of the Message interface implementation.
func (msg *MsgRequestNextPruningPointAndItsAnticoneBlocks) Command() MessageCommand {
	return CmdRequestNextPruningPointAndItsAnticoneBlocks
}

<<<<<<< HEAD
// NewMsgRequestNextPruningPointAndItsAnticoneBlocks returns a new SCR RequestNextPruningPointAndItsAnticoneBlocks message that conforms to the
=======
// NewMsgRequestNextPruningPointAndItsAnticoneBlocks returns a new SCR RequestNextPruningPointAndItsAnticoneBlocks message that conforms to the
>>>>>>> 0e8ed786da4e31df71edebffb326da64b0a6c3b4
// Message interface.
func NewMsgRequestNextPruningPointAndItsAnticoneBlocks() *MsgRequestNextPruningPointAndItsAnticoneBlocks {
	return &MsgRequestNextPruningPointAndItsAnticoneBlocks{}
}
