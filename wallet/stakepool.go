// Copyright (c) 2016 The commanderu developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package wallet

import (
	"errors"

	"github.com/commanderu/cdr/cdrutil"
	"github.com/commanderu/cdrwallet/wallet/udb"
	"github.com/commanderu/cdrwallet/walletdb"
)

// StakePoolUserInfo returns the stake pool user information for a user
// identified by their P2SH voting address.
func (w *Wallet) StakePoolUserInfo(userAddress cdrutil.Address) (*udb.StakePoolUser, error) {
	switch userAddress.(type) {
	case *cdrutil.AddressPubKeyHash: // ok
	case *cdrutil.AddressScriptHash: // ok
	default:
		return nil, errors.New("stake pool user address must be P2PKH or P2SH")
	}

	var user *udb.StakePoolUser
	err := walletdb.View(w.db, func(tx walletdb.ReadTx) error {
		stakemgrNs := tx.ReadBucket(wstakemgrNamespaceKey)
		var err error
		user, err = w.StakeMgr.StakePoolUserInfo(stakemgrNs, userAddress)
		return err
	})
	return user, err
}
