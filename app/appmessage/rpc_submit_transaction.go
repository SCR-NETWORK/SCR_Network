package appmessage

// SubmitTransactionRequestMessage is an appmessage corresponding to
// its respective RPC message
type SubmitTransactionRequestMessage struct {
	baseMessage
	Transaction *RPCTransaction
	AllowOrphan bool
}

// Command returns the protocol command string for the message
func (msg *SubmitTransactionRequestMessage) Command() MessageCommand {
	return CmdSubmitTransactionRequestMessage
}

// NewSubmitTransactionRequestMessage returns a instance of the message
func NewSubmitTransactionRequestMessage(transaction *RPCTransaction, allowOrphan bool) *SubmitTransactionRequestMessage {
	return &SubmitTransactionRequestMessage{
		Transaction: transaction,
		AllowOrphan: allowOrphan,
	}
}

// SubmitTransactionResponseMessage is an appmessage corresponding to
// its respective RPC message
type SubmitTransactionResponseMessage struct {
	baseMessage
	TransactionID string

	Error *RPCError
}

// Command returns the protocol command string for the message
func (msg *SubmitTransactionResponseMessage) Command() MessageCommand {
	return CmdSubmitTransactionResponseMessage
}

// NewSubmitTransactionResponseMessage returns a instance of the message
func NewSubmitTransactionResponseMessage(transactionID string) *SubmitTransactionResponseMessage {
	return &SubmitTransactionResponseMessage{
		TransactionID: transactionID,
	}
}

<<<<<<< HEAD
// RPCTransaction is a SCR-NETWORK transaction representation meant to be
=======
// RPCTransaction is a SCR_Network transaction representation meant to be
>>>>>>> 0e8ed786da4e31df71edebffb326da64b0a6c3b4
// used over RPC
type RPCTransaction struct {
	Version      uint16
	Inputs       []*RPCTransactionInput
	Outputs      []*RPCTransactionOutput
	LockTime     uint64
	SubnetworkID string
	Gas          uint64
	Payload      string
	VerboseData  *RPCTransactionVerboseData
}

<<<<<<< HEAD
// RPCTransactionInput is a SCR-NETWORK transaction input representation
=======
// RPCTransactionInput is a SCR_Network transaction input representation
>>>>>>> 0e8ed786da4e31df71edebffb326da64b0a6c3b4
// meant to be used over RPC
type RPCTransactionInput struct {
	PreviousOutpoint *RPCOutpoint
	SignatureScript  string
	Sequence         uint64
	SigOpCount       byte
	VerboseData      *RPCTransactionInputVerboseData
}

<<<<<<< HEAD
// RPCScriptPublicKey is a SCR-NETWORK ScriptPublicKey representation
=======
// RPCScriptPublicKey is a SCR_Network ScriptPublicKey representation
>>>>>>> 0e8ed786da4e31df71edebffb326da64b0a6c3b4
type RPCScriptPublicKey struct {
	Version uint16
	Script  string
}

<<<<<<< HEAD
// RPCTransactionOutput is a SCR-NETWORK transaction output representation
=======
// RPCTransactionOutput is a SCR_Network transaction output representation
>>>>>>> 0e8ed786da4e31df71edebffb326da64b0a6c3b4
// meant to be used over RPC
type RPCTransactionOutput struct {
	Amount          uint64
	ScriptPublicKey *RPCScriptPublicKey
	VerboseData     *RPCTransactionOutputVerboseData
}

<<<<<<< HEAD
// RPCOutpoint is a SCR-NETWORK outpoint representation meant to be used
=======
// RPCOutpoint is a SCR_Network outpoint representation meant to be used
>>>>>>> 0e8ed786da4e31df71edebffb326da64b0a6c3b4
// over RPC
type RPCOutpoint struct {
	TransactionID string
	Index         uint32
}

<<<<<<< HEAD
// RPCUTXOEntry is a SCR-NETWORK utxo entry representation meant to be used
=======
// RPCUTXOEntry is a SCR_Network utxo entry representation meant to be used
>>>>>>> 0e8ed786da4e31df71edebffb326da64b0a6c3b4
// over RPC
type RPCUTXOEntry struct {
	Amount          uint64
	ScriptPublicKey *RPCScriptPublicKey
	BlockDAAScore   uint64
	IsCoinbase      bool
}

// RPCTransactionVerboseData holds verbose data about a transaction
type RPCTransactionVerboseData struct {
	TransactionID string
	Hash          string
	Mass          uint64
	BlockHash     string
	BlockTime     uint64
}

// RPCTransactionInputVerboseData holds data about a transaction input
type RPCTransactionInputVerboseData struct {
}

// RPCTransactionOutputVerboseData holds data about a transaction output
type RPCTransactionOutputVerboseData struct {
	ScriptPublicKeyType    string
	ScriptPublicKeyAddress string
}
