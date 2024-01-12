// Copyright (c) 2013-2016 The btcsuite developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package appmessage

import (
	"github.com/SCR-NETWORK/SCR_Network/domain/consensus/model/externalapi"
)

<<<<<<< HEAD
// MsgRequestHeaders implements the Message interface and represents a SCR
=======
// MsgRequestHeaders implements the Message interface and represents a SCR
>>>>>>> 0e8ed786da4e31df71edebffb326da64b0a6c3b4
// RequestHeaders message. It is used to request a list of blocks starting after the
// low hash and until the high hash.
type MsgRequestHeaders struct {
	baseMessage
	LowHash  *externalapi.DomainHash
	HighHash *externalapi.DomainHash
}

// Command returns the protocol command string for the message. This is part
// of the Message interface implementation.
func (msg *MsgRequestHeaders) Command() MessageCommand {
	return CmdRequestHeaders
}

<<<<<<< HEAD
// NewMsgRequstHeaders returns a new SCR RequestHeaders message that conforms to the
=======
// NewMsgRequstHeaders returns a new SCR RequestHeaders message that conforms to the
>>>>>>> 0e8ed786da4e31df71edebffb326da64b0a6c3b4
// Message interface using the passed parameters and defaults for the remaining
// fields.
func NewMsgRequstHeaders(lowHash, highHash *externalapi.DomainHash) *MsgRequestHeaders {
	return &MsgRequestHeaders{
		LowHash:  lowHash,
		HighHash: highHash,
	}
}
