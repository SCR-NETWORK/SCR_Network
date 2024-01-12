package appmessage

<<<<<<< HEAD
// MsgTrustedData represents a SCR TrustedData message
=======
// MsgTrustedData represents a SCR TrustedData message
>>>>>>> 0e8ed786da4e31df71edebffb326da64b0a6c3b4
type MsgTrustedData struct {
	baseMessage

	DAAWindow    []*TrustedDataDAAHeader
	GHOSTDAGData []*BlockGHOSTDAGDataHashPair
}

// Command returns the protocol command string for the message
func (msg *MsgTrustedData) Command() MessageCommand {
	return CmdTrustedData
}

// NewMsgTrustedData returns a new MsgTrustedData.
func NewMsgTrustedData() *MsgTrustedData {
	return &MsgTrustedData{}
}

// TrustedDataDAAHeader is an appmessage representation of externalapi.TrustedDataDataDAAHeader
type TrustedDataDAAHeader struct {
	Header       *MsgBlockHeader
	GHOSTDAGData *BlockGHOSTDAGData
}
