package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	"strings"
)

const (
	privateKeyHex1 = "49b6167bfcbcadca22d6a4da24a2762062b749963dc6fee0f065aba8bab35a0d"
	Address1       = "9A1B46d2932b439B4986498281f38C9A58b8f23B"
)

// 假设这是你的智能合约ABI的JSON字符串
// 使用abi方式进行调用
const contractABIJSON1 = `[
	{
		"inputs": [
			{
				"internalType": "uint256",
				"name": "a",
				"type": "uint256"
			},
			{
				"internalType": "uint256",
				"name": "b",
				"type": "uint256"
			}
		],
		"name": "add",
		"outputs": [
			{
				"internalType": "uint256",
				"name": "",
				"type": "uint256"
			}
		],
		"stateMutability": "pure",
		"type": "function"
	},
	{
		"inputs": [
			{
				"internalType": "uint256",
				"name": "a",
				"type": "uint256"
			},
			{
				"internalType": "uint256",
				"name": "b",
				"type": "uint256"
			}
		],
		"name": "subtract",
		"outputs": [
			{
				"internalType": "uint256",
				"name": "",
				"type": "uint256"
			}
		],
		"stateMutability": "pure",
		"type": "function"
	}
]`

func main() {
	// 连接到以太坊节点
	client, err := ethclient.Dial("http://localhost:7545")
	if err != nil {
		fmt.Println("无法连接到以太坊节点:", err)
		return
	}

	// 加载私钥并创建授权对象
	privateKey, err := crypto.HexToECDSA(privateKeyHex1)
	if err != nil {
		fmt.Println("无法解析私钥:", err)
		return
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		fmt.Println("无法转换为公钥对象")
		return
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		fmt.Println("获取交易序号失败:", err)
		return
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		fmt.Println("获取Gas价格失败:", err)
		return
	}

	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(300000)
	auth.GasPrice = gasPrice

	// 智能合约已部署的地址
	contractAddress := common.HexToAddress(Address1)

	// 解析合约ABI
	contractABI, err := abi.JSON(strings.NewReader(contractABIJSON1))
	if err != nil {
		fmt.Println("解析合约ABI失败:", err)
		return
	}

	// 准备合约方法调用的参数
	a := big.NewInt(30)
	b := big.NewInt(20)
	methodArgs := []interface{}{a, b}

	// 准备调用合约方法的数据
	data, err := contractABI.Pack("add", methodArgs...)
	if err != nil {
		fmt.Println("打包合约方法调用数据失败:", err)
		return
	}

	// 创建调用消息
	callMsg := ethereum.CallMsg{
		To:   &contractAddress,
		Data: data,
	}

	// 调用合约方法并获取返回值
	result, err := client.CallContract(context.Background(), callMsg, nil)
	if err != nil {
		fmt.Println("调用合约方法失败:", err)
		return
	}

	// 解析返回值
	var returnValue *big.Int
	err = contractABI.UnpackIntoInterface(&returnValue, "add", result)
	if err != nil {
		fmt.Println("解析合约返回值失败:", err)
		return
	}

	fmt.Println("合约方法返回值:", returnValue)
}
