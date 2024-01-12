// Copyright (c) 2013-2016 The btcsuite developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package appmessage

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	// DefaultServices describes the default services that are supported by
	// the server.
	DefaultServices = SFNodeNetwork | SFNodeBloom | SFNodeCF
)

<<<<<<< HEAD
// ServiceFlag identifies services supported by a SCR peer.
=======
// ServiceFlag identifies services supported by a SCR peer.
>>>>>>> 0e8ed786da4e31df71edebffb326da64b0a6c3b4
type ServiceFlag uint64

const (
	// SFNodeNetwork is a flag used to indicate a peer is a full node.
	SFNodeNetwork ServiceFlag = 1 << iota

	// SFNodeGetUTXO is a flag used to indicate a peer supports the
	// getutxos and utxos commands (BIP0064).
	SFNodeGetUTXO

	// SFNodeBloom is a flag used to indicate a peer supports bloom
	// filtering.
	SFNodeBloom

	// SFNodeXthin is a flag used to indicate a peer supports xthin blocks.
	SFNodeXthin

	// SFNodeBit5 is a flag used to indicate a peer supports a service
	// defined by bit 5.
	SFNodeBit5

	// SFNodeCF is a flag used to indicate a peer supports committed
	// filters (CFs).
	SFNodeCF
)

// Map of service flags back to their constant names for pretty printing.
var sfStrings = map[ServiceFlag]string{
	SFNodeNetwork: "SFNodeNetwork",
	SFNodeGetUTXO: "SFNodeGetUTXO",
	SFNodeBloom:   "SFNodeBloom",
	SFNodeXthin:   "SFNodeXthin",
	SFNodeBit5:    "SFNodeBit5",
	SFNodeCF:      "SFNodeCF",
}

// orderedSFStrings is an ordered list of service flags from highest to
// lowest.
var orderedSFStrings = []ServiceFlag{
	SFNodeNetwork,
	SFNodeGetUTXO,
	SFNodeBloom,
	SFNodeXthin,
	SFNodeBit5,
	SFNodeCF,
}

// String returns the ServiceFlag in human-readable form.
func (f ServiceFlag) String() string {
	// No flags are set.
	if f == 0 {
		return "0x0"
	}

	// Add individual bit flags.
	s := ""
	for _, flag := range orderedSFStrings {
		if f&flag == flag {
			s += sfStrings[flag] + "|"
			f -= flag
		}
	}

	// Add any remaining flags which aren't accounted for as hex.
	s = strings.TrimRight(s, "|")
	if f != 0 {
		s += "|0x" + strconv.FormatUint(uint64(f), 16)
	}
	s = strings.TrimLeft(s, "|")
	return s
}

<<<<<<< HEAD
// SCRSCRNet represents which SCR network a message belongs to.
type SCRSCRNet uint32

// Constants used to indicate the message SCR network. They can also be
=======
// SCRSCRNet represents which SCR network a message belongs to.
type SCRSCRNet uint32

// Constants used to indicate the message SCR network. They can also be
>>>>>>> 0e8ed786da4e31df71edebffb326da64b0a6c3b4
// used to seek to the next message when a stream's state is unknown, but
// this package does not provide that functionality since it's generally a
// better idea to simply disconnect clients that are misbehaving over TCP.
const (
<<<<<<< HEAD
	// Mainnet represents the main SCR network.
=======
	// Mainnet represents the main SCR network.
>>>>>>> 0e8ed786da4e31df71edebffb326da64b0a6c3b4
	Mainnet SCRSCRNet = 0x3ddcf71d

	// Testnet represents the test network.
	Testnet SCRSCRNet = 0xddb8af8f

	// Simnet represents the simulation test network.
	Simnet SCRSCRNet = 0x374dcf1c

	// Devnet represents the development test network.
	Devnet SCRSCRNet = 0x732d87e1
)

<<<<<<< HEAD
// bnStrings is a map of SCR networks back to their constant names for
=======
// bnStrings is a map of SCR networks back to their constant names for
>>>>>>> 0e8ed786da4e31df71edebffb326da64b0a6c3b4
// pretty printing.
var bnStrings = map[SCRSCRNet]string{
	Mainnet: "Mainnet",
	Testnet: "Testnet",
	Simnet:  "Simnet",
	Devnet:  "Devnet",
}

// String returns the SCRSCRNet in human-readable form.
func (n SCRSCRNet) String() string {
	if s, ok := bnStrings[n]; ok {
		return s
	}

	return fmt.Sprintf("Unknown SCRSCRNet (%d)", uint32(n))
}
