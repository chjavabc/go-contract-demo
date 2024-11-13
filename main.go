package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"go-contract-demo/part5"
	"log"
	"math/big"
)

func main() {
	client, err := ethclient.Dial("http://localhost:7545")
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}
	store, err := part5.NewCalculator(common.HexToAddress("9A1B46d2932b439B4986498281f38C9A58b8f23B"), client)
	if err != nil {
		log.Fatalf("Failed to instantiate a Storage contract: %v", err)
	}

	// 准备合约方法调用的参数
	a := big.NewInt(190)
	b := big.NewInt(20)

	// 调用合约方法并获取返回值
	callOpts := &bind.CallOpts{}
	result, err := store.Add(callOpts, a, b)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("1合约方法返回值:", result)

	subtract, err := store.Subtract(callOpts, a, b)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("2合约方法返回值:{}", subtract)
}
