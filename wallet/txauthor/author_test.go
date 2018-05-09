// Copyright (c) 2016 The btcsuite developers
// Copyright (c) 2016 The commanderu developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package txauthor_test

import (
	"testing"

	"github.com/commanderu/cdr/cdrutil"
	"github.com/commanderu/cdr/wire"
	. "github.com/commanderu/cdrwallet/wallet/txauthor"
	"github.com/commanderu/cdrwallet/wallet/txrules"

	"github.com/commanderu/cdrwallet/wallet/internal/txsizes"
)

func p2pkhOutputs(amounts ...cdrutil.Amount) []*wire.TxOut {
	v := make([]*wire.TxOut, 0, len(amounts))
	for _, a := range amounts {
		outScript := make([]byte, txsizes.P2PKHOutputSize)
		v = append(v, wire.NewTxOut(int64(a), outScript))
	}
	return v
}

func makeInputSource(unspents []*wire.TxOut) InputSource {
	// Return outputs in order.
	currentTotal := cdrutil.Amount(0)
	currentInputs := make([]*wire.TxIn, 0, len(unspents))
	f := func(target cdrutil.Amount) (cdrutil.Amount, []*wire.TxIn, [][]byte, error) {
		for currentTotal < target && len(unspents) != 0 {
			u := unspents[0]
			unspents = unspents[1:]
			nextInput := wire.NewTxIn(&wire.OutPoint{}, nil)
			currentTotal += cdrutil.Amount(u.Value)
			currentInputs = append(currentInputs, nextInput)
		}
		return currentTotal, currentInputs, make([][]byte, len(currentInputs)), nil
	}
	return InputSource(f)
}

func TestNewUnsignedTransaction(t *testing.T) {
	tests := []struct {
		UnspentOutputs   []*wire.TxOut
		Outputs          []*wire.TxOut
		RelayFee         cdrutil.Amount
		ChangeAmount     cdrutil.Amount
		InputSourceError bool
		InputCount       int
	}{
		0: {
			UnspentOutputs:   p2pkhOutputs(1e8),
			Outputs:          p2pkhOutputs(1e8),
			RelayFee:         1e3,
			InputSourceError: true,
		},
		1: {
			UnspentOutputs: p2pkhOutputs(1e8),
			Outputs:        p2pkhOutputs(1e6),
			RelayFee:       1e3,
			ChangeAmount: 1e8 - 1e6 - txrules.FeeForSerializeSize(1e3,
				txsizes.EstimateSerializeSize([]txsizes.ScriptSizer{txsizes.P2PKHScriptSize}, p2pkhOutputs(1e6), true)),
			InputCount: 1,
		},
		2: {
			UnspentOutputs: p2pkhOutputs(1e8),
			Outputs:        p2pkhOutputs(1e6),
			RelayFee:       1e4,
			ChangeAmount: 1e8 - 1e6 - txrules.FeeForSerializeSize(1e4,
				txsizes.EstimateSerializeSize([]txsizes.ScriptSizer{txsizes.P2PKHScriptSize}, p2pkhOutputs(1e6), true)),
			InputCount: 1,
		},
		3: {
			UnspentOutputs: p2pkhOutputs(1e8),
			Outputs:        p2pkhOutputs(1e6, 1e6, 1e6),
			RelayFee:       1e4,
			ChangeAmount: 1e8 - 3e6 - txrules.FeeForSerializeSize(1e4,
				txsizes.EstimateSerializeSize([]txsizes.ScriptSizer{txsizes.P2PKHScriptSize}, p2pkhOutputs(1e6, 1e6, 1e6), true)),
			InputCount: 1,
		},
		4: {
			UnspentOutputs: p2pkhOutputs(1e8),
			Outputs:        p2pkhOutputs(1e6, 1e6, 1e6),
			RelayFee:       2.55e3,
			ChangeAmount: 1e8 - 3e6 - txrules.FeeForSerializeSize(2.55e3,
				txsizes.EstimateSerializeSize([]txsizes.ScriptSizer{txsizes.P2PKHScriptSize}, p2pkhOutputs(1e6, 1e6, 1e6), true)),
			InputCount: 1,
		},

		// Test dust thresholds (603 for a 1e3 relay fee).
		5: {
			UnspentOutputs: p2pkhOutputs(1e8),
			Outputs: p2pkhOutputs(1e8 - 602 - txrules.FeeForSerializeSize(1e3,
				txsizes.EstimateSerializeSize([]txsizes.ScriptSizer{txsizes.P2PKHScriptSize}, p2pkhOutputs(0), true))),
			RelayFee:     1e3,
			ChangeAmount: 0,
			InputCount:   1,
		},
		6: {
			UnspentOutputs: p2pkhOutputs(1e8),
			Outputs: p2pkhOutputs(1e8 - 603 - txrules.FeeForSerializeSize(1e3,
				txsizes.EstimateSerializeSize([]txsizes.ScriptSizer{txsizes.P2PKHScriptSize}, p2pkhOutputs(0), true))),
			RelayFee:     1e3,
			ChangeAmount: 603,
			InputCount:   1,
		},

		// Test dust thresholds (1537.65 for a 2.55e3 relay fee).
		7: {
			UnspentOutputs: p2pkhOutputs(1e8),
			Outputs: p2pkhOutputs(1e8 - 1537 - txrules.FeeForSerializeSize(2.55e3,
				txsizes.EstimateSerializeSize([]txsizes.ScriptSizer{txsizes.P2PKHScriptSize}, p2pkhOutputs(0), true))),
			RelayFee:     2.55e3,
			ChangeAmount: 0,
			InputCount:   1,
		},
		8: {
			UnspentOutputs: p2pkhOutputs(1e8),
			Outputs: p2pkhOutputs(1e8 - 1538 - txrules.FeeForSerializeSize(2.55e3,
				txsizes.EstimateSerializeSize([]txsizes.ScriptSizer{txsizes.P2PKHScriptSize}, p2pkhOutputs(0), true))),
			RelayFee:     2.55e3,
			ChangeAmount: 1538,
			InputCount:   1,
		},

		// Test two unspent outputs available but only one needed
		// (tested fee only includes one input rather than using a
		// serialize size for each).
		9: {
			UnspentOutputs: p2pkhOutputs(1e8, 1e8),
			Outputs: p2pkhOutputs(1e8 - 603 - txrules.FeeForSerializeSize(1e3,
				txsizes.EstimateSerializeSize([]txsizes.ScriptSizer{txsizes.P2PKHScriptSize}, p2pkhOutputs(0), true))),
			RelayFee:     1e3,
			ChangeAmount: 603,
			InputCount:   1,
		},

		// Test that second output is not included to make the change
		// output not dust and be included in the transaction.
		//
		// It's debatable whether or not this is a good idea, but it's
		// how the function was written, so test it anyways.
		10: {
			UnspentOutputs: p2pkhOutputs(1e8, 1e8),
			Outputs: p2pkhOutputs(1e8 - 545 - txrules.FeeForSerializeSize(1e3,
				txsizes.EstimateSerializeSize([]txsizes.ScriptSizer{txsizes.P2PKHScriptSize}, p2pkhOutputs(0), true))),
			RelayFee:     1e3,
			ChangeAmount: 0,
			InputCount:   1,
		},

		// Test two unspent outputs available where both are needed.
		11: {
			UnspentOutputs: p2pkhOutputs(1e8, 1e8),
			Outputs:        p2pkhOutputs(1e8),
			RelayFee:       1e3,
			ChangeAmount: 1e8 - txrules.FeeForSerializeSize(1e3,
				txsizes.EstimateSerializeSize([]txsizes.ScriptSizer{txsizes.P2PKHScriptSize, txsizes.P2PKHScriptSize}, p2pkhOutputs(1e8), true)),
			InputCount: 2,
		},

		// Test that zero change outputs are not included
		// (ChangeAmount=0 means don't include any change output).
		12: {
			UnspentOutputs: p2pkhOutputs(1e8),
			Outputs:        p2pkhOutputs(1e8),
			RelayFee:       0,
			ChangeAmount:   0,
			InputCount:     1,
		},
	}

	changeSource := func() ([]byte, uint16, error) {
		// Only length matters for these tests.
		return make([]byte, txsizes.P2PKHPkScriptSize), 0, nil
	}

	for i, test := range tests {
		inputSource := makeInputSource(test.UnspentOutputs)
		tx, err := NewUnsignedTransaction(test.Outputs, test.RelayFee, inputSource, changeSource)
		switch e := err.(type) {
		case nil:
		case InputSourceError:
			if !test.InputSourceError {
				t.Errorf("Test %d: Returned InputSourceError but expected "+
					"change output with amount %v", i, test.ChangeAmount)
			}
			continue
		default:
			t.Errorf("Test %d: Unexpected error: %v", i, e)
			continue
		}
		if tx.ChangeIndex < 0 {
			if test.ChangeAmount != 0 {
				t.Errorf("Test %d: No change output added but expected output with amount %v",
					i, test.ChangeAmount)
				continue
			}
		} else {
			changeAmount := cdrutil.Amount(tx.Tx.TxOut[tx.ChangeIndex].Value)
			if test.ChangeAmount == 0 {
				t.Errorf("Test %d: Included change output with value %v but expected no change",
					i, changeAmount)
				continue
			}
			if changeAmount != test.ChangeAmount {
				t.Errorf("Test %d: Got change amount %v, Expected %v",
					i, changeAmount, test.ChangeAmount)
				continue
			}
		}
		if len(tx.Tx.TxIn) != test.InputCount {
			t.Errorf("Test %d: Used %d outputs from input source, Expected %d",
				i, len(tx.Tx.TxIn), test.InputCount)
		}
	}
}