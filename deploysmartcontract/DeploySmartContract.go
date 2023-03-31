package deploysmartcontract

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	ERC20 "test/ERC20"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func DeploySmartContract() {
	client, err := ethclient.Dial("https://data-seed-prebsc-1-s1.binance.org:8545/")
	if err != nil {
		log.Fatal(err)
	}

	privateKey, err := crypto.HexToECDSA("649bd57ed5bc788f9bc709dd00658a1a095ef972199d00f3a0fe74afd1bcfd89")
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)      // in wei
	auth.GasLimit = uint64(3000000) // in units
	auth.GasPrice = gasPrice
	total := new(big.Int)
	total.SetString("1000000000000000000000000", 10)
	//input := "1.0"
	address, tx, instance, err := ERC20.DeployERC20(auth, client, "binh", "BNB", total)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(address.Hex())   // 0xbab6d3E8b9cC5CcAF8dcce8E6b28E792cF1ac175
	fmt.Println(tx.Hash().Hex()) // 0xf61754e3c421c8bb026db13e11d17252350947fb7274ab17b301a2a2bb31685b

	_ = instance
}
