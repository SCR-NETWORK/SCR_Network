// Copyright (c) 2013-2016 The btcsuite developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package appmessage

<<<<<<< HEAD
// MsgIBDBlock implements the Message interface and represents a SCR
=======
// MsgIBDBlock implements the Message interface and represents a SCR
>>>>>>> 0e8ed786da4e31df71edebffb326da64b0a6c3b4
// ibdblock message. It is used to deliver block and transaction information in
// response to a RequestIBDBlocks message (MsgRequestIBDBlocks).
type MsgIBDBlock struct {
	baseMessage
	*MsgBlock
}

// Command returns the protocol command string for the message. This is part
// of the Message interface implementation.
func (msg *MsgIBDBlock) Command() MessageCommand {
	return CmdIBDBlock
}

// MaxPayloadLength returns the maximum length the payload can be for the
// receiver. This is part of the Message interface implementation.
func (msg *MsgIBDBlock) MaxPayloadLength(pver uint32) uint32 {
	return MaxMessagePayload
}

<<<<<<< HEAD
// NewMsgIBDBlock returns a new SCR ibdblock message that conforms to the
=======
// NewMsgIBDBlock returns a new SCR ibdblock message that conforms to the
>>>>>>> 0e8ed786da4e31df71edebffb326da64b0a6c3b4
// Message interface. See MsgIBDBlock for details.
func NewMsgIBDBlock(msgBlock *MsgBlock) *MsgIBDBlock {
	return &MsgIBDBlock{MsgBlock: msgBlock}
}
