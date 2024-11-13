package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/params"
	"golang.org/x/net/context"
	"log"
	"math/big"
)

const (
	privateKeyHex = "49b6167bfcbcadca22d6a4da24a2762062b749963dc6fee0f065aba8bab35a0d"
	Address       = "5cF5144683dE498915587a0F6A236c78d89288e0"
)

func main() {
	//这个能成功转账～～～～～
	//转账

	// 连接到本地区块链节点
	client, err := ethclient.Dial("http://localhost:7545")
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}

	// 替换为你的私钥
	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		log.Fatalf("Failed to parse private key: %v", err)
	}

	// 获取账户地址
	fromAddress := crypto.PubkeyToAddress(privateKey.PublicKey)

	// 创建转账交易
	toAddress := common.HexToAddress(Address)
	value := big.NewInt(1000000000000000000) // 1 ETH in wei

	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatalf("Failed to get nonce: %v", err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatalf("Failed to suggest gas price: %v", err)
	}

	tx := types.NewTransaction(nonce, toAddress, value, params.TxGas, gasPrice, nil)

	// 签名交易
	chainID, err := client.ChainID(context.Background())
	if err != nil {
		log.Fatalf("Failed to get network ID: %v", err)
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Fatalf("Failed to sign transaction: %v", err)
	}

	// 发送交易
	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatalf("Failed to send transaction: %v", err)
	}

	fmt.Printf("Transaction sent: %s\n", signedTx.Hash().Hex())
}
