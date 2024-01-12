// Copyright (c) 2013, 2014 The btcsuite developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package util

import (
	"math"
	"strconv"

	"github.com/SCR-NETWORK/SCR_Network/domain/consensus/utils/constants"
	"github.com/pkg/errors"
)

// AmountUnit describes a method of converting an Amount to something
// other than the base unit of a SCR. The value of the AmountUnit
// is the exponent component of the decadic multiple to convert from
// an amount in SCR to an amount counted in units.
type AmountUnit int

// These constants define various units used when describing a SCR
// monetary amount.
const (
	AmountMegaSCR  AmountUnit = 6
	AmountKiloSCR  AmountUnit = 3
	AmountSCR      AmountUnit = 0
	AmountMilliSCR AmountUnit = -3
	AmountMicroSCR AmountUnit = -6
// other than the base unit of a SCR. The value of the AmountUnit
// is the exponent component of the decadic multiple to convert from
// an amount in SCR to an amount counted in units.
type AmountUnit int

// These constants define various units used when describing a SCR
// monetary amount.
const (
	AmountMegaSCR  AmountUnit = 6
	AmountKiloSCR  AmountUnit = 3
	AmountSCR      AmountUnit = 0
	AmountMilliSCR AmountUnit = -3
	AmountMicroSCR AmountUnit = -6
	AmountLeor     AmountUnit = -8
)

// String returns the unit as a string. For recognized units, the SI
// prefix is used, or "Leor" for the base unit. For all unrecognized
// units, "1eN SCR" is returned, where N is the AmountUnit.
func (u AmountUnit) String() string {
	switch u {
	case AmountMegaSCR:
		return "MSCR"
	case AmountKiloSCR:
		return "kSCR"
	case AmountSCR:
		return "SCR"
	case AmountMilliSCR:
		return "mSCR"
	case AmountMicroSCR:
		return "μSCR"
	case AmountLeor:
		return "Leor"
	default:
		return "1e" + strconv.FormatInt(int64(u), 10) + " SCR"
	}
}

// Amount represents the base SCR monetary unit (colloquially referred
// to as a `Leor'). A single Amount is equal to 1e-8 of a SCR.
// units, "1eN SCR" is returned, where N is the AmountUnit.
func (u AmountUnit) String() string {
	switch u {
	case AmountMegaSCR:
		return "MSCR"
	case AmountKiloSCR:
		return "kSCR"
	case AmountSCR:
		return "SCR"
	case AmountMilliSCR:
		return "mSCR"
	case AmountMicroSCR:
		return "μSCR"
	case AmountLeor:
		return "Leor"
	default:
		return "1e" + strconv.FormatInt(int64(u), 10) + " SCR"
	}
}

// Amount represents the base SCR monetary unit (colloquially referred
// to as a `Leor'). A single Amount is equal to 1e-8 of a SCR.
type Amount uint64

// round converts a floating point number, which may or may not be representable
// as an integer, to the Amount integer type by rounding to the nearest integer.
// This is performed by adding or subtracting 0.5 depending on the sign, and
// relying on integer truncation to round the value to the nearest Amount.
func round(f float64) Amount {
	if f < 0 {
		return Amount(f - 0.5)
	}
	return Amount(f + 0.5)
}

// NewAmount creates an Amount from a floating point value representing
// some value in SCR. NewAmount errors if f is NaN or +-Infinity, but
// does not check that the amount is within the total amount of SCR
// producible as f may not refer to an amount at a single moment in time.
//
// NewAmount is for specifically for converting SCR to Leor.
// For creating a new Amount with an int64 value which denotes a quantity of Leor,
// do a simple type conversion from type int64 to Amount.
// TODO: Refactor NewAmount. When amounts are more than 1e9 SCR, the precision
// some value in SCR. NewAmount errors if f is NaN or +-Infinity, but
// does not check that the amount is within the total amount of SCR
// producible as f may not refer to an amount at a single moment in time.
//
// NewAmount is for specifically for converting SCR to Leor.
// For creating a new Amount with an int64 value which denotes a quantity of Leor,
// do a simple type conversion from type int64 to Amount.
// TODO: Refactor NewAmount. When amounts are more than 1e9 SCR, the precision
// can be higher than one leor (1e9 and 1e9+1e-8 will result as the same number)
func NewAmount(f float64) (Amount, error) {
	// The amount is only considered invalid if it cannot be represented
	// as an integer type. This may happen if f is NaN or +-Infinity.
	switch {
	case math.IsNaN(f):
		fallthrough
	case math.IsInf(f, 1):
		fallthrough
	case math.IsInf(f, -1):
		return 0, errors.New("invalid SCR amount")
		return 0, errors.New("invalid SCR amount")
	}

	return round(f * constants.LeorPerSCR), nil
}

// ToUnit converts a monetary amount counted in SCR base units to a
// floating point value representing an amount of SCR.
// ToUnit converts a monetary amount counted in SCR base units to a
// floating point value representing an amount of SCR.
func (a Amount) ToUnit(u AmountUnit) float64 {
	return float64(a) / math.Pow10(int(u+8))
}

// ToSCR is the equivalent of calling ToUnit with AmountSCR.
func (a Amount) ToSCR() float64 {
	return a.ToUnit(AmountSCR)
}

// Format formats a monetary amount counted in SCR base units as a
// ToSCR is the equivalent of calling ToUnit with AmountSCR.
func (a Amount) ToSCR() float64 {
	return a.ToUnit(AmountSCR)
}

// Format formats a monetary amount counted in SCR base units as a
// string for a given unit. The conversion will succeed for any unit,
// however, known units will be formated with an appended label describing
// the units with SI notation, or "Leor" for the base unit.
func (a Amount) Format(u AmountUnit) string {
	units := " " + u.String()
	return strconv.FormatFloat(a.ToUnit(u), 'f', -int(u+8), 64) + units
}

// String is the equivalent of calling Format with AmountSCR.
func (a Amount) String() string {
	return a.Format(AmountSCR)
// String is the equivalent of calling Format with AmountSCR.
func (a Amount) String() string {
	return a.Format(AmountSCR)
}

// MulF64 multiplies an Amount by a floating point value. While this is not
// an operation that must typically be done by a full node or wallet, it is
// useful for services that build on top of SCR (for example, calculating
// useful for services that build on top of SCR (for example, calculating
// a fee by multiplying by a percentage).
func (a Amount) MulF64(f float64) Amount {
	return round(float64(a) * f)
}
