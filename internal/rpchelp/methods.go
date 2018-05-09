// Copyright (c) 2015 The btcsuite developers
// Copyright (c) 2015-2017 The commanderu developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

//+build !generate

package rpchelp

import "github.com/commanderu/cdr/cdrjson"

// Common return types.
var (
	returnsBool        = []interface{}{(*bool)(nil)}
	returnsNumber      = []interface{}{(*float64)(nil)}
	returnsString      = []interface{}{(*string)(nil)}
	returnsStringArray = []interface{}{(*[]string)(nil)}
	returnsLTRArray    = []interface{}{(*[]cdrjson.ListTransactionsResult)(nil)}
)

// Methods contains all methods and result types that help is generated for,
// for every locale.
var Methods = []struct {
	Method      string
	ResultTypes []interface{}
}{
	{"accountaddressindex", []interface{}{(*int)(nil)}},
	{"accountsyncaddressindex", nil},
	{"addmultisigaddress", returnsString},
	{"consolidate", returnsString},
	{"createmultisig", []interface{}{(*cdrjson.CreateMultiSigResult)(nil)}},
	{"dumpprivkey", returnsString},
	{"getaccount", returnsString},
	{"getaccountaddress", returnsString},
	{"getaddressesbyaccount", returnsStringArray},
	{"getbalance", []interface{}{(*cdrjson.GetBalanceResult)(nil)}},
	{"getbestblockhash", returnsString},
	{"getblockcount", returnsNumber},
	{"getinfo", []interface{}{(*cdrjson.InfoWalletResult)(nil)}},
	{"getmasterpubkey", []interface{}{(*string)(nil)}},
	{"getmultisigoutinfo", []interface{}{(*cdrjson.GetMultisigOutInfoResult)(nil)}},
	{"getnewaddress", returnsString},
	{"getrawchangeaddress", returnsString},
	{"getreceivedbyaccount", returnsNumber},
	{"getreceivedbyaddress", returnsNumber},
	{"gettickets", []interface{}{(*cdrjson.GetTicketsResult)(nil)}},
	{"gettransaction", []interface{}{(*cdrjson.GetTransactionResult)(nil)}},
	{"getvotechoices", []interface{}{(*cdrjson.GetVoteChoicesResult)(nil)}},
	{"help", append(returnsString, returnsString[0])},
	{"importprivkey", nil},
	{"importscript", nil},
	{"keypoolrefill", nil},
	{"listaccounts", []interface{}{(*map[string]float64)(nil)}},
	{"listlockunspent", []interface{}{(*[]cdrjson.TransactionInput)(nil)}},
	{"listreceivedbyaccount", []interface{}{(*[]cdrjson.ListReceivedByAccountResult)(nil)}},
	{"listreceivedbyaddress", []interface{}{(*[]cdrjson.ListReceivedByAddressResult)(nil)}},
	{"listsinceblock", []interface{}{(*cdrjson.ListSinceBlockResult)(nil)}},
	{"listtransactions", returnsLTRArray},
	{"listunspent", []interface{}{(*cdrjson.ListUnspentResult)(nil)}},
	{"lockunspent", returnsBool},
	{"redeemmultisigout", []interface{}{(*cdrjson.RedeemMultiSigOutResult)(nil)}},
	{"redeemmultisigouts", []interface{}{(*cdrjson.RedeemMultiSigOutResult)(nil)}},
	{"rescanwallet", nil},
	{"revoketickets", nil},
	{"sendfrom", returnsString},
	{"sendmany", returnsString},
	{"sendtoaddress", returnsString},
	{"sendtomultisig", returnsString},
	{"settxfee", returnsBool},
	{"setvotechoice", nil},
	{"signmessage", returnsString},
	{"signrawtransaction", []interface{}{(*cdrjson.SignRawTransactionResult)(nil)}},
	{"signrawtransactions", []interface{}{(*cdrjson.SignRawTransactionsResult)(nil)}},
	{"startautobuyer", nil},
	{"stopautobuyer", nil},
	{"sweepaccount", []interface{}{(*cdrjson.SweepAccountResult)(nil)}},
	{"validateaddress", []interface{}{(*cdrjson.ValidateAddressWalletResult)(nil)}},
	{"verifymessage", returnsBool},
	{"version", []interface{}{(*map[string]cdrjson.VersionResult)(nil)}},
	{"walletlock", nil},
	{"walletpassphrase", nil},
	{"walletpassphrasechange", nil},
	{"createnewaccount", nil},
	{"exportwatchingwallet", returnsString},
	{"getbestblock", []interface{}{(*cdrjson.GetBestBlockResult)(nil)}},
	{"getunconfirmedbalance", returnsNumber},
	{"listaddresstransactions", returnsLTRArray},
	{"listalltransactions", returnsLTRArray},
	{"renameaccount", nil},
	{"walletislocked", returnsBool},
	{"walletinfo", []interface{}{(*cdrjson.WalletInfoResult)(nil)}},

	// TODO Alphabetize
	{"purchaseticket", returnsString},
	{"generatevote", []interface{}{(*cdrjson.GenerateVoteResult)(nil)}},
	{"getstakeinfo", []interface{}{(*cdrjson.GetStakeInfoResult)(nil)}},
	{"getticketfee", returnsNumber},
	{"setticketfee", returnsBool},
	{"getwalletfee", returnsNumber},
	{"addticket", nil},
	{"listscripts", []interface{}{(*cdrjson.ListScriptsResult)(nil)}},
	{"stakepooluserinfo", []interface{}{(*cdrjson.StakePoolUserInfoResult)(nil)}},
	{"ticketsforaddress", returnsBool},
}

// HelpDescs contains the locale-specific help strings along with the locale.
var HelpDescs = []struct {
	Locale   string // Actual locale, e.g. en_US
	GoLocale string // Locale used in Go names, e.g. EnUS
	Descs    map[string]string
}{
	{"en_US", "EnUS", helpDescsEnUS}, // helpdescs_en_US.go
}
