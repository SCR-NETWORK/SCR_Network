// Copyright (c) 2013-2015 The btcsuite developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package appmessage

import (
	"github.com/SCR-NETWORK/SCR_Network/domain/consensus/model/externalapi"
)

<<<<<<< HEAD
// MsgRequestAddresses implements the Message interface and represents a SCR
=======
// MsgRequestAddresses implements the Message interface and represents a SCR
>>>>>>> 0e8ed786da4e31df71edebffb326da64b0a6c3b4
// RequestAddresses message. It is used to request a list of known active peers on the
// network from a peer to help identify potential nodes. The list is returned
// via one or more addr messages (MsgAddresses).
//
// This message has no payload.
type MsgRequestAddresses struct {
	baseMessage
	IncludeAllSubnetworks bool
	SubnetworkID          *externalapi.DomainSubnetworkID
}

// Command returns the protocol command string for the message. This is part
// of the Message interface implementation.
func (msg *MsgRequestAddresses) Command() MessageCommand {
	return CmdRequestAddresses
}

<<<<<<< HEAD
// NewMsgRequestAddresses returns a new SCR RequestAddresses message that conforms to the
=======
// NewMsgRequestAddresses returns a new SCR RequestAddresses message that conforms to the
>>>>>>> 0e8ed786da4e31df71edebffb326da64b0a6c3b4
// Message interface. See MsgRequestAddresses for details.
func NewMsgRequestAddresses(includeAllSubnetworks bool, subnetworkID *externalapi.DomainSubnetworkID) *MsgRequestAddresses {
	return &MsgRequestAddresses{
		IncludeAllSubnetworks: includeAllSubnetworks,
		SubnetworkID:          subnetworkID,
	}
}
