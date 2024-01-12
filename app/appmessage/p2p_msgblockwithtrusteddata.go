package appmessage

import (
	"math/big"

	"github.com/SCR-NETWORK/SCR_Network/domain/consensus/model/externalapi"
)

<<<<<<< HEAD
// MsgBlockWithTrustedData represents a SCR BlockWithTrustedData message
=======
// MsgBlockWithTrustedData represents a SCR BlockWithTrustedData message
>>>>>>> 0e8ed786da4e31df71edebffb326da64b0a6c3b4
type MsgBlockWithTrustedData struct {
	baseMessage

	Block        *MsgBlock
	DAAScore     uint64
	DAAWindow    []*TrustedDataDataDAABlock
	GHOSTDAGData []*BlockGHOSTDAGDataHashPair
}

// Command returns the protocol command string for the message
func (msg *MsgBlockWithTrustedData) Command() MessageCommand {
	return CmdBlockWithTrustedData
}

// NewMsgBlockWithTrustedData returns a new MsgBlockWithTrustedData.
func NewMsgBlockWithTrustedData() *MsgBlockWithTrustedData {
	return &MsgBlockWithTrustedData{}
}

// TrustedDataDataDAABlock is an appmessage representation of externalapi.TrustedDataDataDAABlock
type TrustedDataDataDAABlock struct {
	Block        *MsgBlock
	GHOSTDAGData *BlockGHOSTDAGData
}

// BlockGHOSTDAGData is an appmessage representation of externalapi.BlockGHOSTDAGData
type BlockGHOSTDAGData struct {
	BlueScore          uint64
	BlueWork           *big.Int
	SelectedParent     *externalapi.DomainHash
	MergeSetBlues      []*externalapi.DomainHash
	MergeSetReds       []*externalapi.DomainHash
	BluesAnticoneSizes []*BluesAnticoneSizes
}

// BluesAnticoneSizes is an appmessage representation of the BluesAnticoneSizes part of GHOSTDAG data.
type BluesAnticoneSizes struct {
	BlueHash     *externalapi.DomainHash
	AnticoneSize externalapi.KType
}

// BlockGHOSTDAGDataHashPair is an appmessage representation of externalapi.BlockGHOSTDAGDataHashPair
type BlockGHOSTDAGDataHashPair struct {
	Hash         *externalapi.DomainHash
	GHOSTDAGData *BlockGHOSTDAGData
}
