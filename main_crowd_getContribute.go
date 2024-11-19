package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"go-contract-demo/part5"
	"golang.org/x/net/context"
	"log"
)

func main() {
	client, err := ethclient.Dial("http://localhost:7545")
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}
	store, err := part5.NewCrowdFundCaller(common.HexToAddress("aa565f3d8CB7F5eAAaF73aA05Ef25f04BCF5734a"), client)
	if err != nil {
		log.Fatalf("Failed to instantiate a Storage contract: %v", err)
	}

	// 准备合约方法调用的参数, 账户1
	contributorAddress := common.HexToAddress("9Cdd40e9487295936F5DC5768675A4888B563E2")

	//方式1 调用合约的getContributionAmount方法
	amount, err := store.GetContributionAmount(nil, contributorAddress)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("1合约方法返回值:", amount.String())

	//方式2  调用合约的getContributionAmount方法
	contributionAmount, err := store.GetContributionAmount(&bind.CallOpts{Context: context.Background()}, contributorAddress)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("2合约方法返回值:", contributionAmount.String())

	// 方式3 写法

	// 调用合约方法并获取返回值
	callOpts := &bind.CallOpts{}
	contributionAmount1, err := store.GetContributionAmount(callOpts, contributorAddress)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("2合约方法返回值:", contributionAmount1.String())

}
