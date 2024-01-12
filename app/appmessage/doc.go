/*
<<<<<<< HEAD
Package appmessage implements the SCR appmessage protocol.

At a high level, this package provides support for marshalling and unmarshalling
supported SCR messages to and from the appmessage. This package does not deal
=======
Package appmessage implements the SCR appmessage protocol.

At a high level, this package provides support for marshalling and unmarshalling
supported SCR messages to and from the appmessage. This package does not deal
>>>>>>> 0e8ed786da4e31df71edebffb326da64b0a6c3b4
with the specifics of message handling such as what to do when a message is
received. This provides the caller with a high level of flexibility.

# SCR Message Overview

<<<<<<< HEAD
The SCR protocol consists of exchanging messages between peers. Each
message is preceded by a header which identifies information about it such as
which SCR network it is a part of, its type, how big it is, and a checksum
to verify validity. All encoding and decoding of message headers is handled by
this package.

To accomplish this, there is a generic interface for SCR messages named
Message which allows messages of any type to be read, written, or passed around
through channels, functions, etc. In addition, concrete implementations of most
of the currently supported SCR messages are provided. For these supported
messages, all of the details of marshalling and unmarshalling to and from the
appmessage using SCR encoding are handled so the caller doesn't have to concern
=======
The SCR protocol consists of exchanging messages between peers. Each
message is preceded by a header which identifies information about it such as
which SCR network it is a part of, its type, how big it is, and a checksum
to verify validity. All encoding and decoding of message headers is handled by
this package.

To accomplish this, there is a generic interface for SCR messages named
Message which allows messages of any type to be read, written, or passed around
through channels, functions, etc. In addition, concrete implementations of most
of the currently supported SCR messages are provided. For these supported
messages, all of the details of marshalling and unmarshalling to and from the
appmessage using SCR encoding are handled so the caller doesn't have to concern
>>>>>>> 0e8ed786da4e31df71edebffb326da64b0a6c3b4
themselves with the specifics.

# Message Interaction

<<<<<<< HEAD
The following provides a quick summary of how the SCR messages are intended
=======
The following provides a quick summary of how the SCR messages are intended
>>>>>>> 0e8ed786da4e31df71edebffb326da64b0a6c3b4
to interact with one another. As stated above, these interactions are not
directly handled by this package.

The initial handshake consists of two peers sending each other a version message
(MsgVersion) followed by responding with a verack message (MsgVerAck). Both
peers use the information in the version message (MsgVersion) to negotiate
things such as protocol version and supported services with each other. Once
the initial handshake is complete, the following chart indicates message
interactions in no particular order.

	Peer A Sends                            Peer B Responds
	----------------------------------------------------------------------------
	getaddr message (MsgRequestAddresses)       addr message (MsgAddresses)
	getblockinvs message (MsgGetBlockInvs)  inv message (MsgInv)
	inv message (MsgInv)                    getdata message (MsgGetData)
	getdata message (MsgGetData)            block message (MsgBlock) -or-
	                                        tx message (MsgTx) -or-
	                                        notfound message (MsgNotFound)
	ping message (MsgPing)                  pong message (MsgPong)

# Common Parameters

There are several common parameters that arise when using this package to read
<<<<<<< HEAD
and write SCR messages. The following sections provide a quick overview of
=======
and write SCR messages. The following sections provide a quick overview of
>>>>>>> 0e8ed786da4e31df71edebffb326da64b0a6c3b4
these parameters so the next sections can build on them.

# Protocol Version

The protocol version should be negotiated with the remote peer at a higher
level than this package via the version (MsgVersion) message exchange, however,
this package provides the appmessage.ProtocolVersion constant which indicates the
latest protocol version this package supports and is typically the value to use
for all outbound connections before a potentially lower protocol version is
negotiated.

# SCR Network

<<<<<<< HEAD
The SCR network is a magic number which is used to identify the start of a
message and which SCR network the message applies to. This package provides
=======
The SCR network is a magic number which is used to identify the start of a
message and which SCR network the message applies to. This package provides
>>>>>>> 0e8ed786da4e31df71edebffb326da64b0a6c3b4
the following constants:

	appmessage.Mainnet
	appmessage.Testnet (Test network)
	appmessage.Simnet  (Simulation test network)
	appmessage.Devnet  (Development network)

# Determining Message Type

<<<<<<< HEAD
As discussed in the SCR message overview section, this package reads
and writes SCR messages using a generic interface named Message. In
=======
As discussed in the SCR message overview section, this package reads
and writes SCR messages using a generic interface named Message. In
>>>>>>> 0e8ed786da4e31df71edebffb326da64b0a6c3b4
order to determine the actual concrete type of the message, use a type
switch or type assertion. An example of a type switch follows:

	// Assumes msg is already a valid concrete message such as one created
	// via NewMsgVersion or read via ReadMessage.
	switch msg := msg.(type) {
	case *appmessage.MsgVersion:
		// The message is a pointer to a MsgVersion struct.
		fmt.Printf("Protocol version: %d", msg.ProtocolVersion)
	case *appmessage.MsgBlock:
		// The message is a pointer to a MsgBlock struct.
		fmt.Printf("Number of tx in block: %d", msg.Header.TxnCount)
	}

# Reading Messages

<<<<<<< HEAD
In order to unmarshall SCR messages from the appmessage, use the ReadMessage
function. It accepts any io.Reader, but typically this will be a net.Conn to
a remote node running a SCR peer. Example syntax is:

	// Reads and validates the next SCR message from conn using the
	// protocol version pver and the SCR network SCRSCR. The returns
=======
In order to unmarshall SCR messages from the appmessage, use the ReadMessage
function. It accepts any io.Reader, but typically this will be a net.Conn to
a remote node running a SCR peer. Example syntax is:

	// Reads and validates the next SCR message from conn using the
	// protocol version pver and the SCR network SCRSCR. The returns
>>>>>>> 0e8ed786da4e31df71edebffb326da64b0a6c3b4
	// are a appmessage.Message, a []byte which contains the unmarshalled
	// raw payload, and a possible error.
	msg, rawPayload, err := appmessage.ReadMessage(conn, pver, SCRSCR)
	if err != nil {
		// Log and handle the error
	}

# Writing Messages

<<<<<<< HEAD
In order to marshall SCR messages to the appmessage, use the WriteMessage
function. It accepts any io.Writer, but typically this will be a net.Conn to
a remote node running a SCR peer. Example syntax to request addresses
from a remote peer is:

	// Create a new getaddr SCR message.
	msg := appmessage.NewMsgRequestAddresses()

	// Writes a SCR message msg to conn using the protocol version
	// pver, and the SCR network SCRSCR. The return is a possible
=======
In order to marshall SCR messages to the appmessage, use the WriteMessage
function. It accepts any io.Writer, but typically this will be a net.Conn to
a remote node running a SCR peer. Example syntax to request addresses
from a remote peer is:

	// Create a new getaddr SCR message.
	msg := appmessage.NewMsgRequestAddresses()

	// Writes a SCR message msg to conn using the protocol version
	// pver, and the SCR network SCRSCR. The return is a possible
>>>>>>> 0e8ed786da4e31df71edebffb326da64b0a6c3b4
	// error.
	err := appmessage.WriteMessage(conn, msg, pver, SCRSCR)
	if err != nil {
		// Log and handle the error
	}

# Errors

Errors returned by this package are either the raw errors provided by underlying
calls to read/write from streams such as io.EOF, io.ErrUnexpectedEOF, and
io.ErrShortWrite, or of type appmessage.MessageError. This allows the caller to
differentiate between general IO errors and malformed messages through type
assertions.
*/
package appmessage
