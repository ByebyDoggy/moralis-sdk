package main

import (
	"fmt"
	"log"
	"os"
	"time"

	moralisapi "github.com/ByebyDoggy/moralis-sdk/restapi"

	"github.com/joho/godotenv"
)

const (
	_host       = "https://deep-index.moralis.io/api/v2.2"
	_serverHost = "https://wamaxhbnkkbj.usemoralis.com:2053/server"
	// _serverHost = "https://mwq7lmluczgx.usemoralis.com:2053/server"

	_address = "0xcc7bcf633f6ce26ce3ed9e255b8eaa6f219a0956"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	apiKey := os.Getenv("API_KEY")
	if apiKey == "" {
		log.Fatal("API_KEY is empty")
	}

	c := moralisapi.NewClient(_host, apiKey, time.Second*5)

	//tx, err := c.GetTransactionByHash(&moralisapi.GetTransactionByHashInput{
	//	Hash: "0x25cd7f3bfc73aa0fa3ea1307b0271cb78d703594389462f596fc033ca5170c2b",
	//})
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//fmt.Printf("GetTransactionByHash: %+v", tx)

	// resp, err := c.GetTransactionsByAddress(&moralisapi.GetTransactionsByAddressInput{
	// 	Address: _address,
	// 	Chain:   moralisapi.ChainRopsten,
	// })
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Printf("GetTransactionsByAddress (%s): %+v\n", _address, resp)

	// respBalance, err := c.GetBalanceByAddress(&moralisapi.GetBalanceByAddressInput{
	// 	Address: _address,
	// 	Chain:   moralisapi.ChainEth,
	// })
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Printf("GetBalanceByAddress (%s): %+v\n", _address, respBalance)

	//respErc20Balance, err := c.GetERC20BalanceByAddress(&moralisapi.GetERC20BalanceByAddressInput{
	//	Address: _address,
	//	Chain:   moralisapi.ChainEth,
	//	// TokenAddresses: []string{"0xdAC17F958D2ee523a2206206994597C13D831ec7"},
	//})
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//for i, balance := range respErc20Balance {
	//	fmt.Printf("%d. GetERC20BalanceByAddress (%s): %+v\n", i+1, _address, balance)
	//}

	walletCreated, err := c.GetWalletActiveChains(&moralisapi.GetWalletActiveChainsInput{
		Address: _address,
	})
	if err != nil {
		log.Fatal(err)
	}
	for _, activeChain := range walletCreated.ActiveChains {
		fmt.Printf("Chain: %s, ChainID: %s, FirstTransaction: %+v, LastTransaction: %+v\n",
			activeChain.Chain, activeChain.ChainID, activeChain.FirstTransaction, activeChain.LastTransaction)
	}

	// respErc20Transfers, err := c.GetERC20TransfersByAddress(&moralisapi.GetERC20TransfersByAddressInput{
	// 	Address: _address,
	// })
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Printf("GetERC20TransfersByAddress (%s): %+v\n", _address, respErc20Transfers)

}
