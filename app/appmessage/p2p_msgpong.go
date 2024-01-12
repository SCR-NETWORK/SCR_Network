// Copyright (c) 2013-2015 The btcsuite developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package appmessage

<<<<<<< HEAD
// MsgPong implements the Message interface and represents a SCR pong
// message which is used primarily to confirm that a connection is still valid
// in response to a SCR ping message (MsgPing).
=======
// MsgPong implements the Message interface and represents a SCR pong
// message which is used primarily to confirm that a connection is still valid
// in response to a SCR ping message (MsgPing).
>>>>>>> 0e8ed786da4e31df71edebffb326da64b0a6c3b4
//
// This message was not added until protocol versions AFTER BIP0031Version.
type MsgPong struct {
	baseMessage
	// Unique value associated with message that is used to identify
	// specific ping message.
	Nonce uint64
}

// Command returns the protocol command string for the message. This is part
// of the Message interface implementation.
func (msg *MsgPong) Command() MessageCommand {
	return CmdPong
}

<<<<<<< HEAD
// NewMsgPong returns a new SCR pong message that conforms to the Message
=======
// NewMsgPong returns a new SCR pong message that conforms to the Message
>>>>>>> 0e8ed786da4e31df71edebffb326da64b0a6c3b4
// interface. See MsgPong for details.
func NewMsgPong(nonce uint64) *MsgPong {
	return &MsgPong{
		Nonce: nonce,
	}
}
