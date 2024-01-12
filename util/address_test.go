// Copyright (c) 2013-2017 The btcsuite developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package util_test

import (
	"bytes"
	"fmt"
	"os"
	"reflect"
	"strings"
	"testing"

	"github.com/SCR-NETWORK/SCR_Network/util"
)

func TestAddresses(t *testing.T) {
	if os.Getenv("SKIP_ADDRESSES_RELATED_TESTS") != "" {
		t.Skip()
	}

	tests := []struct {
		name           string
		addr           string
		encoded        string
		valid          bool
		result         util.Address
		f              func() (util.Address, error)
		passedPrefix   util.Bech32Prefix
		expectedPrefix util.Bech32Prefix
	}{
		// Positive P2PK tests.
		{
			name:    "mainnet p2pk",
			addr:    "SCR:qq2u4dvcpvdcn0k03afx0lgqhlj06kk0ay8std3t87tuexlhq7m87e4n078d6",
			encoded: "SCR:qq2u4dvcpvdcn0k03afx0lgqhlj06kk0ay8std3t87tuexlhq7m87e4n078d6",
			addr:    "SCR:qq2u4dvcpvdcn0k03afx0lgqhlj06kk0ay8std3t87tuexlhq7m87e4n078d6",
			encoded: "SCR:qq2u4dvcpvdcn0k03afx0lgqhlj06kk0ay8std3t87tuexlhq7m87e4n078d6",
			valid:   true,
			result: util.TstAddressPubKey(
				util.Bech32PrefixSCR,
				[util.PublicKeySize]byte{
					0x2e, 0xfb, 0x76, 0xe3, 0x2f, 0x34, 0x67, 0x43, 0xc0, 0x2,
					0x5b, 0x59, 0x97, 0xb7, 0xcc, 0x3c, 0xaa, 0xf, 0x14, 0x72,
					0xe9, 0xe3, 0x60, 0x8b, 0x11, 0xdb, 0x9b, 0x98, 0x16, 0xac,
					0x6b, 0x3b,
				}),
			f: func() (util.Address, error) {
				publicKey := []byte{
					0x2e, 0xfb, 0x76, 0xe3, 0x2f, 0x34, 0x67, 0x43, 0xc0, 0x2,
					0x5b, 0x59, 0x97, 0xb7, 0xcc, 0x3c, 0xaa, 0xf, 0x14, 0x72,
					0xe9, 0xe3, 0x60, 0x8b, 0x11, 0xdb, 0x9b, 0x98, 0x16, 0xac,
					0x6b, 0x3b}
				return util.NewAddressPublicKey(publicKey, util.Bech32PrefixSCR)
			},
			passedPrefix:   util.Bech32PrefixUnknown,
			expectedPrefix: util.Bech32PrefixSCR,
		},
		{
			name:    "mainnet p2pk 2",
			addr:    "SCR:qq80qvqs0lfxuzmt7sz3909ze6camq9d4t35ennsep3hxfe7ln35cvfqgz3z8",
			encoded: "SCR:qq80qvqs0lfxuzmt7sz3909ze6camq9d4t35ennsep3hxfe7ln35cvfqgz3z8",
			addr:    "SCR:qq80qvqs0lfxuzmt7sz3909ze6camq9d4t35ennsep3hxfe7ln35cvfqgz3z8",
			encoded: "SCR:qq80qvqs0lfxuzmt7sz3909ze6camq9d4t35ennsep3hxfe7ln35cvfqgz3z8",
			valid:   true,
			result: util.TstAddressPubKey(
				util.Bech32PrefixSCR,
				[util.PublicKeySize]byte{
					0x0e, 0xf0, 0x30, 0x10, 0x7f, 0xd2, 0x6e, 0x0b, 0x6b, 0xf4,
					0x05, 0x12, 0xbc, 0xa2, 0xce, 0xb1, 0xdd, 0x80, 0xad, 0xaa,
					0xe3, 0x4c, 0xce, 0x70, 0xc8, 0x63, 0x73, 0x27, 0x3e, 0xfc,
					0xe3, 0x4c,
				}),
			f: func() (util.Address, error) {
				publicKey := []byte{
					0x0e, 0xf0, 0x30, 0x10, 0x7f, 0xd2, 0x6e, 0x0b, 0x6b, 0xf4,
					0x05, 0x12, 0xbc, 0xa2, 0xce, 0xb1, 0xdd, 0x80, 0xad, 0xaa,
					0xe3, 0x4c, 0xce, 0x70, 0xc8, 0x63, 0x73, 0x27, 0x3e, 0xfc,
					0xe3, 0x4c,
				}
				return util.NewAddressPublicKey(publicKey, util.Bech32PrefixSCR)
			},
			passedPrefix:   util.Bech32PrefixSCR,
			expectedPrefix: util.Bech32PrefixSCR,
		},
		{
			name:    "testnet p2pk",
			addr:    "SCRtest:qph45l0y247vu078ktt3g0r0qefeqqqxrushmz4hlv2fjg6vfk3njk7pp4mh0",
			encoded: "SCRtest:qph45l0y247vu078ktt3g0r0qefeqqqxrushmz4hlv2fjg6vfk3njk7pp4mh0",
			addr:    "SCRtest:qph45l0y247vu078ktt3g0r0qefeqqqxrushmz4hlv2fjg6vfk3njk7pp4mh0",
			encoded: "SCRtest:qph45l0y247vu078ktt3g0r0qefeqqqxrushmz4hlv2fjg6vfk3njk7pp4mh0",
			valid:   true,
			result: util.TstAddressPubKey(
				util.Bech32PrefixSCRTest,
				[util.PublicKeySize]byte{
					0x6f, 0x5a, 0x7d, 0xe4, 0x55, 0x7c, 0xce, 0x3f, 0xc7, 0xb2,
					0xd7, 0x14, 0x3c, 0x6f, 0x06, 0x53, 0x90, 0x00, 0x06, 0x1f,
					0x21, 0x7d, 0x8a, 0xb7, 0xfb, 0x14, 0x99, 0x23, 0x4c, 0x4d,
					0xa3, 0x39,
				}),
			f: func() (util.Address, error) {
				publicKey := []byte{
					0x6f, 0x5a, 0x7d, 0xe4, 0x55, 0x7c, 0xce, 0x3f, 0xc7, 0xb2,
					0xd7, 0x14, 0x3c, 0x6f, 0x06, 0x53, 0x90, 0x00, 0x06, 0x1f,
					0x21, 0x7d, 0x8a, 0xb7, 0xfb, 0x14, 0x99, 0x23, 0x4c, 0x4d,
					0xa3, 0x39,
				}
				return util.NewAddressPublicKey(publicKey, util.Bech32PrefixSCRTest)
			},
			passedPrefix:   util.Bech32PrefixSCRTest,
			expectedPrefix: util.Bech32PrefixSCRTest,
		},

		// ECDSA P2PK tests.
		{
			name:    "mainnet ecdsa p2pk",
			addr:    "SCR:q835ennsep3hxfe7lnz5ee7j5jgmkjswsn35ennsep3hxfe7ln35e2sm7yrlr4w",
			encoded: "SCR:q835ennsep3hxfe7lnz5ee7j5jgmkjswsn35ennsep3hxfe7ln35e2sm7yrlr4w",
			addr:    "SCR:q835ennsep3hxfe7lnz5ee7j5jgmkjswsn35ennsep3hxfe7ln35e2sm7yrlr4w",
			encoded: "SCR:q835ennsep3hxfe7lnz5ee7j5jgmkjswsn35ennsep3hxfe7ln35e2sm7yrlr4w",
			valid:   true,
			result: util.TstAddressPubKeyECDSA(
				util.Bech32PrefixSCR,
				[util.PublicKeySizeECDSA]byte{
					0xe3, 0x4c, 0xce, 0x70, 0xc8, 0x63, 0x73, 0x27, 0x3e, 0xfc,
					0xc5, 0x4c, 0xe7, 0xd2, 0xa4, 0x91, 0xbb, 0x4a, 0x0e, 0x84,
					0xe3, 0x4c, 0xce, 0x70, 0xc8, 0x63, 0x73, 0x27, 0x3e, 0xfc,
					0xe3, 0x4c, 0xaa,
				}),
			f: func() (util.Address, error) {
				publicKey := []byte{
					0xe3, 0x4c, 0xce, 0x70, 0xc8, 0x63, 0x73, 0x27, 0x3e, 0xfc,
					0xc5, 0x4c, 0xe7, 0xd2, 0xa4, 0x91, 0xbb, 0x4a, 0x0e, 0x84,
					0xe3, 0x4c, 0xce, 0x70, 0xc8, 0x63, 0x73, 0x27, 0x3e, 0xfc,
					0xe3, 0x4c, 0xaa}
				return util.NewAddressPublicKeyECDSA(publicKey, util.Bech32PrefixSCR)
			},
			passedPrefix:   util.Bech32PrefixUnknown,
			expectedPrefix: util.Bech32PrefixSCR,
		},

		// Negative P2PK tests.
		{
			name:  "p2pk wrong public key length",
			addr:  "",
			valid: false,
			f: func() (util.Address, error) {
				publicKey := []byte{
					0x00, 0x0e, 0xf0, 0x30, 0x10, 0x7f, 0xd2, 0x6e, 0x0b, 0x6b,
					0xf4, 0x05, 0x12, 0xbc, 0xa2, 0xce, 0xb1, 0xdd, 0x80, 0xad,
					0xaa}
				return util.NewAddressPublicKey(publicKey, util.Bech32PrefixSCR)
			},
			passedPrefix:   util.Bech32PrefixSCR,
			expectedPrefix: util.Bech32PrefixSCR,
		},
		{
			name:           "p2pk bad checksum",
			addr:           "SCR:qr35ennsep3hxfe7lnz5ee7j5jgmkjswss74as46gx",
			addr:           "SCR:qr35ennsep3hxfe7lnz5ee7j5jgmkjswss74as46gx",
			valid:          false,
			passedPrefix:   util.Bech32PrefixSCR,
			expectedPrefix: util.Bech32PrefixSCR,
		},

		// Positive P2SH tests.
		{
			name:    "mainnet p2sh",
			addr:    "SCR:prq20q4qd9ulr044cauyy9wtpeupqpjv67pn2vyc6acly7xqkrjdzmh8rj9f4",
			encoded: "SCR:prq20q4qd9ulr044cauyy9wtpeupqpjv67pn2vyc6acly7xqkrjdzmh8rj9f4",
			addr:    "SCR:prq20q4qd9ulr044cauyy9wtpeupqpjv67pn2vyc6acly7xqkrjdzmh8rj9f4",
			encoded: "SCR:prq20q4qd9ulr044cauyy9wtpeupqpjv67pn2vyc6acly7xqkrjdzmh8rj9f4",
			valid:   true,
			result: util.TstAddressScriptHash(
				util.Bech32PrefixSCR,
				[32]byte{
					0xc0, 0xa7, 0x82, 0xa0, 0x69, 0x79, 0xf1, 0xbe,
					0xb5, 0xc7, 0x78, 0x42, 0x15, 0xcb, 0x0e, 0x78,
					0x10, 0x06, 0x4c, 0xd7, 0x83, 0x35, 0x30, 0x98,
					0xd7, 0x71, 0xf2, 0x78, 0xc0, 0xb0, 0xe4, 0xd1,
				}),
			f: func() (util.Address, error) {
				script := []byte{
					0x52, 0x41, 0x04, 0x91, 0xbb, 0xa2, 0x51, 0x09, 0x12, 0xa5,
					0xbd, 0x37, 0xda, 0x1f, 0xb5, 0xb1, 0x67, 0x30, 0x10, 0xe4,
					0x3d, 0x2c, 0x6d, 0x81, 0x2c, 0x51, 0x4e, 0x91, 0xbf, 0xa9,
					0xf2, 0xeb, 0x12, 0x9e, 0x1c, 0x18, 0x33, 0x29, 0xdb, 0x55,
					0xbd, 0x86, 0x8e, 0x20, 0x9a, 0xac, 0x2f, 0xbc, 0x02, 0xcb,
					0x33, 0xd9, 0x8f, 0xe7, 0x4b, 0xf2, 0x3f, 0x0c, 0x23, 0x5d,
					0x61, 0x26, 0xb1, 0xd8, 0x33, 0x4f, 0x86, 0x41, 0x04, 0x86,
					0x5c, 0x40, 0x29, 0x3a, 0x68, 0x0c, 0xb9, 0xc0, 0x20, 0xe7,
					0xb1, 0xe1, 0x06, 0xd8, 0xc1, 0x91, 0x6d, 0x3c, 0xef, 0x99,
					0xaa, 0x43, 0x1a, 0x56, 0xd2, 0x53, 0xe6, 0x92, 0x56, 0xda,
					0xc0, 0x9e, 0xf1, 0x22, 0xb1, 0xa9, 0x86, 0x81, 0x8a, 0x7c,
					0xb6, 0x24, 0x53, 0x2f, 0x06, 0x2c, 0x1d, 0x1f, 0x87, 0x22,
					0x08, 0x48, 0x61, 0xc5, 0xc3, 0x29, 0x1c, 0xcf, 0xfe, 0xf4,
					0xec, 0x68, 0x74, 0x41, 0x04, 0x8d, 0x24, 0x55, 0xd2, 0x40,
					0x3e, 0x08, 0x70, 0x8f, 0xc1, 0xf5, 0x56, 0x00, 0x2f, 0x1b,
					0x6c, 0xd8, 0x3f, 0x99, 0x2d, 0x08, 0x50, 0x97, 0xf9, 0x97,
					0x4a, 0xb0, 0x8a, 0x28, 0x83, 0x8f, 0x07, 0x89, 0x6f, 0xba,
					0xb0, 0x8f, 0x39, 0x49, 0x5e, 0x15, 0xfa, 0x6f, 0xad, 0x6e,
					0xdb, 0xfb, 0x1e, 0x75, 0x4e, 0x35, 0xfa, 0x1c, 0x78, 0x44,
					0xc4, 0x1f, 0x32, 0x2a, 0x18, 0x63, 0xd4, 0x62, 0x13, 0x53,
					0xae}
				return util.NewAddressScriptHash(script, util.Bech32PrefixSCR)
			},
			passedPrefix:   util.Bech32PrefixSCR,
			expectedPrefix: util.Bech32PrefixSCR,
		},
		{
			name:    "mainnet p2sh 2",
			addr:    "SCR:pr5vxqxg0xrwl2zvxlq9rxffqx00sm44kn5vxqxg0xrwl2zvxl5vxyhvsake2",
			encoded: "SCR:pr5vxqxg0xrwl2zvxlq9rxffqx00sm44kn5vxqxg0xrwl2zvxl5vxyhvsake2",
			addr:    "SCR:pr5vxqxg0xrwl2zvxlq9rxffqx00sm44kn5vxqxg0xrwl2zvxl5vxyhvsake2",
			encoded: "SCR:pr5vxqxg0xrwl2zvxlq9rxffqx00sm44kn5vxqxg0xrwl2zvxl5vxyhvsake2",
			valid:   true,
			result: util.TstAddressScriptHash(
				util.Bech32PrefixSCR,
				[32]byte{
					0xe8, 0xc3, 0x00, 0xc8, 0x79, 0x86, 0xef, 0xa8, 0x4c, 0x37,
					0xc0, 0x51, 0x99, 0x29, 0x01, 0x9e, 0xf8, 0x6e, 0xb5, 0xb4,
					0xe8, 0xc3, 0x00, 0xc8, 0x79, 0x86, 0xef, 0xa8, 0x4c, 0x37,
					0xe8, 0xc3,
				}),
			f: func() (util.Address, error) {
				hash := []byte{
					0xe8, 0xc3, 0x00, 0xc8, 0x79, 0x86, 0xef, 0xa8, 0x4c, 0x37,
					0xc0, 0x51, 0x99, 0x29, 0x01, 0x9e, 0xf8, 0x6e, 0xb5, 0xb4,
					0xe8, 0xc3, 0x00, 0xc8, 0x79, 0x86, 0xef, 0xa8, 0x4c, 0x37,
					0xe8, 0xc3,
				}
				return util.NewAddressScriptHashFromHash(hash, util.Bech32PrefixSCR)
			},
			passedPrefix:   util.Bech32PrefixSCR,
			expectedPrefix: util.Bech32PrefixSCR,
		},
		{
			name:    "testnet p2sh",
			addr:    "SCRtest:qpcs7zh3ufpv20eekyyy0f7u6fgz0mm9mf9s778qqx2zfvfmnxx86gjmfwhnh",
			encoded: "SCRtest:qpcs7zh3ufpv20eekyyy0f7u6fgz0mm9mf9s778qqx2zfvfmnxx86gjmfwhnh",
			addr:    "SCRtest:qpcs7zh3ufpv20eekyyy0f7u6fgz0mm9mf9s778qqx2zfvfmnxx86gjmfwhnh",
			encoded: "SCRtest:qpcs7zh3ufpv20eekyyy0f7u6fgz0mm9mf9s778qqx2zfvfmnxx86gjmfwhnh",
			valid:   true,
			result: util.TstAddressScriptHash(
				util.Bech32PrefixSCRTest,
				[32]byte{
					0xc5, 0x79, 0x34, 0x2c, 0x2c, 0x4c, 0x92, 0x20, 0x20, 0x5e,
					0x2c, 0xdc, 0x28, 0x56, 0x17, 0x04, 0x0c, 0x92, 0x4a, 0x0a,
					0xe8, 0xc3, 0x00, 0xc8, 0x79, 0x86, 0xef, 0xa8, 0x4c, 0x37,
					0xe8, 0xc3,
				}),
			f: func() (util.Address, error) {
				hash := []byte{
					0xc5, 0x79, 0x34, 0x2c, 0x2c, 0x4c, 0x92, 0x20, 0x20, 0x5e,
					0x2c, 0xdc, 0x28, 0x56, 0x17, 0x04, 0x0c, 0x92, 0x4a, 0x0a,
					0xe8, 0xc3, 0x00, 0xc8, 0x79, 0x86, 0xef, 0xa8, 0x4c, 0x37,
					0xe8, 0xc3,
				}
				return util.NewAddressScriptHashFromHash(hash, util.Bech32PrefixSCRTest)
			},
			passedPrefix:   util.Bech32PrefixSCRTest,
			expectedPrefix: util.Bech32PrefixSCRTest,
		},

		// Negative P2SH tests.
		{
			name:  "p2sh wrong hash length",
			addr:  "",
			valid: false,
			f: func() (util.Address, error) {
				hash := []byte{
					0x00, 0xf8, 0x15, 0xb0, 0x36, 0xd9, 0xbb, 0xbc, 0xe5, 0xe9,
					0xf2, 0xa0, 0x0a, 0xbd, 0x1b, 0xf3, 0xdc, 0x91, 0xe9, 0x55,
					0x10}
				return util.NewAddressScriptHashFromHash(hash, util.Bech32PrefixSCR)
			},
			passedPrefix:   util.Bech32PrefixSCR,
			expectedPrefix: util.Bech32PrefixSCR,
		},
	}

	for _, test := range tests {
		// Decode addr and compare error against valid.
		decoded, err := util.DecodeAddress(test.addr, test.passedPrefix)
		if (err == nil) != test.valid {
			t.Errorf("%v: decoding test failed: %v", test.name, err)
			return
		}

		if err == nil {
			// Ensure the stringer returns the same address as the
			// original.
			if decodedStringer, ok := decoded.(fmt.Stringer); ok {
				addr := test.addr

				if addr != decodedStringer.String() {
					t.Errorf("%v: String on decoded value does not match expected value: %v != %v",
						test.name, test.addr, decodedStringer.String())
					return
				}
			}

			// Encode again and compare against the original.
			encoded := decoded.EncodeAddress()
			if test.encoded != encoded {
				t.Errorf("%v: decoding and encoding produced different addressess: %v != %v",
					test.name, test.encoded, encoded)
				return
			}

			// Perform type-specific calculations.
			var saddr []byte
			switch decoded.(type) {
			case *util.AddressPublicKey:
				saddr = util.TstAddressSAddrP2PK(encoded)

			case *util.AddressPublicKeyECDSA:
				saddr = util.TstAddressSAddrP2PKECDSA(encoded)

			case *util.AddressScriptHash:
				saddr = util.TstAddressSAddrP2SH(encoded)
			}

			// Check script address, as well as the HashBlake2b method for P2SH addresses.
			if !bytes.Equal(saddr, decoded.ScriptAddress()) {
				t.Errorf("%v: script addresses do not match:\n%x != \n%x",
					test.name, saddr, decoded.ScriptAddress())
				return
			}
			switch a := decoded.(type) {
			case *util.AddressPublicKey:
				if h := a.ScriptAddress()[:]; !bytes.Equal(saddr, h) {
					t.Errorf("%v: hashes do not match:\n%x != \n%x",
						test.name, saddr, h)
					return
				}

			case *util.AddressScriptHash:
				if h := a.HashBlake3()[:]; !bytes.Equal(saddr, h) {
					t.Errorf("%v: hashes do not match:\n%x != \n%x",
						test.name, saddr, h)
					return
				}
			}

			// Ensure the address is for the expected network.
			if !decoded.IsForPrefix(test.expectedPrefix) {
				t.Errorf("%v: calculated network does not match expected",
					test.name)
				return
			}
		}

		if !test.valid {
			// If address is invalid, but a creation function exists,
			// verify that it returns a nil addr and non-nil error.
			if test.f != nil {
				_, err := test.f()
				if err == nil {
					t.Errorf("%v: address is invalid but creating new address succeeded",
						test.name)
					return
				}
			}
			continue
		}

		// Valid test, compare address created with f against expected result.
		addr, err := test.f()
		if err != nil {
			t.Errorf("%v: address is valid but creating new address failed with error %v",
				test.name, err)
			return
		}

		if !reflect.DeepEqual(addr, test.result) {
			t.Errorf("%v: created address does not match expected result",
				test.name)
			return
		}

		if !reflect.DeepEqual(addr, decoded) {
			t.Errorf("%v: created address does not match the decoded address",
				test.name)
			return
		}

		if !reflect.DeepEqual(addr, decoded) {
			t.Errorf("%v: created address does not match the decoded address",
				test.name)
			return
		}
	}
}

func TestDecodeAddressErrorConditions(t *testing.T) {
	tests := []struct {
		address      string
		prefix       util.Bech32Prefix
		errorMessage string
	}{
		{
			"bitcoincash:qpzry9x8gf2tvdw0s3jn54khce6mua7lcw20ayyn",
			util.Bech32PrefixUnknown,
			"decoded address's prefix could not be parsed",
		},
		//{
		//	"SCRsim:qz830led3jg2wpym5nv8wfg6g3nz8cwt92zhc4azpzgjr6rz95nezuqmzj40a",
		//	"SCRsim:qz830led3jg2wpym5nv8wfg6g3nz8cwt92zhc4azpzgjr6rz95nezuqmzj40a",
		//	util.Bech32PrefixSCRSim,
		//	"unknown address type",
		//},
		//{
		//	"SCRsim:raskzcg58mth0an",
		//	"SCRsim:raskzcg58mth0an",
		//	util.Bech32PrefixSCRSim,
		//	"unknown address type",
		//},
		{
			"SCRtest:qpcs7zh3ufpv20eekyyy0f7u6fgz0mm9mf9s778qqx2zfvfmnxx86gjmfwhnh",
			"SCRtest:qpcs7zh3ufpv20eekyyy0f7u6fgz0mm9mf9s778qqx2zfvfmnxx86gjmfwhnh",
			util.Bech32PrefixSCR,
			"decoded address is of wrong network",
		},
	}

	for _, test := range tests {
		_, err := util.DecodeAddress(test.address, test.prefix)
		if err == nil {
			t.Errorf("decodeAddress unexpectedly succeeded")
		} else if !strings.Contains(err.Error(), test.errorMessage) {
			t.Errorf("received mismatched error. Expected '%s' but got '%s'",
				test.errorMessage, err)
		}
	}
}

func TestParsePrefix(t *testing.T) {
	tests := []struct {
		prefixStr      string
		expectedPrefix util.Bech32Prefix
		expectedError  bool
	}{
		{"SCR", util.Bech32PrefixSCR, false},
		{"SCRtest", util.Bech32PrefixSCRTest, false},
		{"SCRsim", util.Bech32PrefixSCRSim, false},
		{"SCR", util.Bech32PrefixSCR, false},
		{"SCRtest", util.Bech32PrefixSCRTest, false},
		{"SCRsim", util.Bech32PrefixSCRSim, false},
		{"blabla", util.Bech32PrefixUnknown, true},
		{"unknown", util.Bech32PrefixUnknown, true},
		{"", util.Bech32PrefixUnknown, true},
	}

	for _, test := range tests {
		result, err := util.ParsePrefix(test.prefixStr)
		if (err != nil) != test.expectedError {
			t.Errorf("TestParsePrefix: %s: expected error status: %t, but got %t",
				test.prefixStr, test.expectedError, err != nil)
		}

		if result != test.expectedPrefix {
			t.Errorf("TestParsePrefix: %s: expected prefix: %d, but got %d",
				test.prefixStr, test.expectedPrefix, result)
		}
	}
}

func TestPrefixToString(t *testing.T) {
	tests := []struct {
		prefix            util.Bech32Prefix
		expectedPrefixStr string
	}{
		{util.Bech32PrefixSCR, "SCR"},
		{util.Bech32PrefixSCRTest, "SCRtest"},
		{util.Bech32PrefixSCRSim, "SCRsim"},
		{util.Bech32PrefixSCR, "SCR"},
		{util.Bech32PrefixSCRTest, "SCRtest"},
		{util.Bech32PrefixSCRSim, "SCRsim"},
		{util.Bech32PrefixUnknown, ""},
	}

	for _, test := range tests {
		result := test.prefix.String()

		if result != test.expectedPrefixStr {
			t.Errorf("TestPrefixToString: %s: expected string: %s, but got %s",
				test.prefix, test.expectedPrefixStr, result)
		}
	}
}