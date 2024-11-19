package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"go-contract-demo/part5"
	"log"
	"math/big"
)

func main() {
	// 连接到以太坊节点
	client, err := ethclient.Dial("http://127.0.0.1:7545") // 使用你的以太坊节点地址
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}

	// 智能合约的地址
	contractAddress := common.HexToAddress("15c7004D9B00825df1778b20BbC6Cf9586B9BF98")

	// 初始化CrowdFunding合约
	crowdFunding, err := part5.NewCrowdFund(contractAddress, client)
	if err != nil {
		log.Fatalf("Failed to initialize contract: %v", err)
	}

	// 调用合约方法
	// 例如，调用contribute方法
	privateKey := "49b6167bfcbcadca22d6a4da24a2762062b749963dc6fee0f065aba8bab35a0d"

	// 替换为你的私钥
	privateKey1, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		log.Fatalf("Failed to parse private key: %v", err)
	}

	auth := bind.NewKeyedTransactor(privateKey1) // 使用你的私钥

	if err != nil {
		log.Fatalf("Failed to create transactor: %v", err)
	}

	// 确保合约没有关闭
	closed, err := crowdFunding.Closed(&bind.CallOpts{})
	if err != nil {
		log.Fatalf("Failed to call Closed: %v", err)
	}
	if closed {
		log.Fatal("Crowdfunding is already closed")
	}

	//一定需要设置gas和上下文
	auth.Value = big.NewInt(1000000000000000000)
	auth.GasLimit = uint64(300000) // 设置合理的 Gas 限制
	auth.GasPrice, err = client.SuggestGasPrice(context.Background())

	// 贡献资金
	tx, err := crowdFunding.Contribute(auth)
	if err != nil {
		log.Fatalf("Failed to contribute: %v", err)
	}
	fmt.Printf("Contribution transaction hash: %v", tx.Hash())

}
