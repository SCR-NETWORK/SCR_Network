package util_test

import (
	"fmt"
	"math"
	"math/big"

	"github.com/SCR-NETWORK/SCR_Network/util/difficulty"

	"github.com/SCR-NETWORK/SCR_Network/util"
)

func ExampleAmount() {

	a := util.Amount(0)
	fmt.Println("Zero Leor:", a)

	a = util.Amount(1e8)
	fmt.Println("100,000,000 Leor:", a)

	a = util.Amount(1e5)
	fmt.Println("100,000 Leor:", a)
	// Output:
	// Zero Leor: 0 SCR
	// 100,000,000 Leor: 1 SCR
	// 100,000 Leor: 0.001 SCR
	// Zero Leor: 0 SCR
	// 100,000,000 Leor: 1 SCR
	// 100,000 Leor: 0.001 SCR
}

func ExampleNewAmount() {
	amountOne, err := util.NewAmount(1)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(amountOne) //Output 1

	amountFraction, err := util.NewAmount(0.01234567)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(amountFraction) //Output 2

	amountZero, err := util.NewAmount(0)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(amountZero) //Output 3

	amountNaN, err := util.NewAmount(math.NaN())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(amountNaN) //Output 4

	// Output: 1 SCR
	// 0.01234567 SCR
	// 0 SCR
	// invalid SCR amount
	// Output: 1 SCR
	// 0.01234567 SCR
	// 0 SCR
	// invalid SCR amount
}

func ExampleAmount_unitConversions() {
	amount := util.Amount(44433322211100)

	fmt.Println("Leor to kSCR:", amount.Format(util.AmountKiloSCR))
	fmt.Println("Leor to SCR:", amount)
	fmt.Println("Leor to MilliSCR:", amount.Format(util.AmountMilliSCR))
	fmt.Println("Leor to MicroSCR:", amount.Format(util.AmountMicroSCR))
	fmt.Println("Leor to Leor:", amount.Format(util.AmountLeor))

	// Output:
	// Leor to kSCR: 444.333222111 kSCR
	// Leor to SCR: 444333.222111 SCR
	// Leor to MilliSCR: 444333222.111 mSCR
	// Leor to MicroSCR: 444333222111 μSCR
	fmt.Println("Leor to kSCR:", amount.Format(util.AmountKiloSCR))
	fmt.Println("Leor to SCR:", amount)
	fmt.Println("Leor to MilliSCR:", amount.Format(util.AmountMilliSCR))
	fmt.Println("Leor to MicroSCR:", amount.Format(util.AmountMicroSCR))
	fmt.Println("Leor to Leor:", amount.Format(util.AmountLeor))

	// Output:
	// Leor to kSCR: 444.333222111 kSCR
	// Leor to SCR: 444333.222111 SCR
	// Leor to MilliSCR: 444333222.111 mSCR
	// Leor to MicroSCR: 444333222111 μSCR
	// Leor to Leor: 44433322211100 Leor
}

// This example demonstrates how to convert the compact "bits" in a block header
// which represent the target difficulty to a big integer and display it using
// the typical hex notation.
func ExampleCompactToBig() {
	bits := uint32(419465580)
	targetDifficulty := difficulty.CompactToBig(bits)

	// Display it in hex.
	fmt.Printf("%064x\n", targetDifficulty.Bytes())

	// Output:
	// 0000000000000000896c00000000000000000000000000000000000000000000
}

// This example demonstrates how to convert a target difficulty into the compact
// "bits" in a block header which represent that target difficulty .
func ExampleBigToCompact() {
	// Convert the target difficulty from block 300000 in the bitcoin
	// main chain to compact form.
	t := "0000000000000000896c00000000000000000000000000000000000000000000"
	targetDifficulty, success := new(big.Int).SetString(t, 16)
	if !success {
		fmt.Println("invalid target difficulty")
		return
	}
	bits := difficulty.BigToCompact(targetDifficulty)

	fmt.Println(bits)

	// Output:
	// 419465580
}
