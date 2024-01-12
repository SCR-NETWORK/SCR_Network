// Copyright (c) 2013-2015 The btcsuite developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package appmessage

<<<<<<< HEAD
// MsgVerAck defines a SCR verack message which is used for a peer to
=======
// MsgVerAck defines a SCR verack message which is used for a peer to
>>>>>>> 0e8ed786da4e31df71edebffb326da64b0a6c3b4
// acknowledge a version message (MsgVersion) after it has used the information
// to negotiate parameters. It implements the Message interface.
//
// This message has no payload.
type MsgVerAck struct {
	baseMessage
}

// Command returns the protocol command string for the message. This is part
// of the Message interface implementation.
func (msg *MsgVerAck) Command() MessageCommand {
	return CmdVerAck
}

<<<<<<< HEAD
// NewMsgVerAck returns a new SCR verack message that conforms to the
=======
// NewMsgVerAck returns a new SCR verack message that conforms to the
>>>>>>> 0e8ed786da4e31df71edebffb326da64b0a6c3b4
// Message interface.
func NewMsgVerAck() *MsgVerAck {
	return &MsgVerAck{}
}
