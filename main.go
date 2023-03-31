package main

import (
	account "test/accounts"

	//import deploysmartcontract "test/deploysmartcontract"

	transaction "test/transactions"
)

// đọc số dư ERC20 và coinbase
// Tạo transaction chuyển ERC20 tự deploy
// 0x0310A60e312189441afc1b12CeaA645a496D8dFc
// 0xa2F0540C71691A2E05E9124DC87Ec485c50ad19E
func main() {
	// account.AccountBalancesCoinbase()
	transaction.TransferringETH()
	account.AccountTokenBalances()

	//deploysmartcontract.DeploySmartContract()
}
