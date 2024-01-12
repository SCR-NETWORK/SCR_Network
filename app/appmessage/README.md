wire
====

[![ISC License](http://img.shields.io/badge/license-ISC-blue.svg)](https://choosealicense.com/licenses/isc/)
[![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg)](http://godoc.org/github.com/SCR-NETWORK/SCR_Network/wire)
=======

<<<<<<< HEAD
Package wire implements the SCR wire protocol.

## SCR Message Overview

The SCR protocol consists of exchanging messages between peers. Each message
is preceded by a header which identifies information about it such as which
SCR network it is a part of, its type, how big it is, and a checksum to
verify validity. All encoding and decoding of message headers is handled by this
package.

To accomplish this, there is a generic interface for SCR messages named
`Message` which allows messages of any type to be read, written, or passed
around through channels, functions, etc. In addition, concrete implementations
of most all SCR messages are provided. All of the details of marshalling and 
unmarshalling to and from the wire using SCR encoding are handled so the 
=======
Package wire implements the SCR wire protocol.

## SCR Message Overview

The SCR protocol consists of exchanging messages between peers. Each message
is preceded by a header which identifies information about it such as which
SCR network it is a part of, its type, how big it is, and a checksum to
verify validity. All encoding and decoding of message headers is handled by this
package.

To accomplish this, there is a generic interface for SCR messages named
`Message` which allows messages of any type to be read, written, or passed
around through channels, functions, etc. In addition, concrete implementations
of most all SCR messages are provided. All of the details of marshalling and 
unmarshalling to and from the wire using SCR encoding are handled so the 
>>>>>>> 0e8ed786da4e31df71edebffb326da64b0a6c3b4
caller doesn't have to concern themselves with the specifics.

## Reading Messages Example

<<<<<<< HEAD
In order to unmarshal SCR messages from the wire, use the `ReadMessage`
function. It accepts any `io.Reader`, but typically this will be a `net.Conn`
to a remote node running a SCR peer. Example syntax is:

```Go
	// Use the most recent protocol version supported by the package and the
	// main SCR network.
	pver := wire.ProtocolVersion
	SCRSCR := wire.Mainnet

	// Reads and validates the next SCR message from conn using the
	// protocol version pver and the SCR network SCRSCR. The returns
=======
In order to unmarshal SCR messages from the wire, use the `ReadMessage`
function. It accepts any `io.Reader`, but typically this will be a `net.Conn`
to a remote node running a SCR peer. Example syntax is:

```Go
	// Use the most recent protocol version supported by the package and the
	// main SCR network.
	pver := wire.ProtocolVersion
	SCRSCR := wire.Mainnet

	// Reads and validates the next SCR message from conn using the
	// protocol version pver and the SCR network SCRSCR. The returns
>>>>>>> 0e8ed786da4e31df71edebffb326da64b0a6c3b4
	// are a appmessage.Message, a []byte which contains the unmarshalled
	// raw payload, and a possible error.
	msg, rawPayload, err := wire.ReadMessage(conn, pver, SCRSCR)
	if err != nil {
		// Log and handle the error
	}
```

See the package documentation for details on determining the message type.

## Writing Messages Example

<<<<<<< HEAD
In order to marshal SCR messages to the wire, use the `WriteMessage`
function. It accepts any `io.Writer`, but typically this will be a `net.Conn`
to a remote node running a SCR peer. Example syntax to request addresses
=======
In order to marshal SCR messages to the wire, use the `WriteMessage`
function. It accepts any `io.Writer`, but typically this will be a `net.Conn`
to a remote node running a SCR peer. Example syntax to request addresses
>>>>>>> 0e8ed786da4e31df71edebffb326da64b0a6c3b4
from a remote peer is:

```Go
	// Use the most recent protocol version supported by the package and the
	// main bitcoin network.
	pver := wire.ProtocolVersion
	SCRSCR := wire.Mainnet

<<<<<<< HEAD
	// Create a new getaddr SCR message.
	msg := wire.NewMsgGetAddr()

	// Writes a SCR message msg to conn using the protocol version
	// pver, and the SCR network SCRSCR. The return is a possible
=======
	// Create a new getaddr SCR message.
	msg := wire.NewMsgGetAddr()

	// Writes a SCR message msg to conn using the protocol version
	// pver, and the SCR network SCRSCR. The return is a possible
>>>>>>> 0e8ed786da4e31df71edebffb326da64b0a6c3b4
	// error.
	err := wire.WriteMessage(conn, msg, pver, SCRSCR)
	if err != nil {
		// Log and handle the error
	}
```
