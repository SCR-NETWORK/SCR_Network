// Copyright (c) 2013-2016 The btcsuite developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package appmessage

import (
	"github.com/pkg/errors"
)

<<<<<<< HEAD
// MaxInvPerMsg is the maximum number of inventory vectors that can be in any type of SCR inv message.
=======
// MaxInvPerMsg is the maximum number of inventory vectors that can be in any type of SCR inv message.
>>>>>>> 0e8ed786da4e31df71edebffb326da64b0a6c3b4
const MaxInvPerMsg = 1 << 17

// errNonCanonicalVarInt is the common format string used for non-canonically
// encoded variable length integer errors.
var errNonCanonicalVarInt = "non-canonical varint %x - discriminant %x must " +
	"encode a value greater than %x"

// errNoEncodingForType signifies that there's no encoding for the given type.
var errNoEncodingForType = errors.New("there's no encoding for this type")
