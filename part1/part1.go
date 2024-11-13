package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
)

const (
	privateKey      = "49b6167bfcbcadca22d6a4da24a2762062b749963dc6fee0f065aba8bab35a0d"
	contractAddress = "4449C18fF263765663b566B2D576c5f6Cc60d96D"
	toAddress       = "2a92ec1b06559f1d24008665cb08b3713ad8f054c744dab28c3bb2f6441d237f"
)

func transfer(client *ethclient.Client, privateKey, toAddress, contract string) (string, error) {

	//从私钥推导出 公钥
	privateKeyECDSA, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		fmt.Println("crypto.HexToECDSA error ,", err)
		return "", err
	}
	publicKey := privateKeyECDSA.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		fmt.Println("publicKeyECDSA error ,", err)
		return "", err
	}
	//从公钥推导出钱包地址
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	fmt.Println("钱包地址：", fromAddress.Hex())
	//构造请求参数
	var data []byte
	methodName := crypto.Keccak256([]byte("add(uint256 a, uint256 b)"))[:4]
	paddedToAddress := common.LeftPadBytes(common.HexToAddress(toAddress).Bytes(), 32)
	amount, _ := new(big.Int).SetString("100000000000000000000", 10)
	paddedAmount := common.LeftPadBytes(amount.Bytes(), 32)
	data = append(data, methodName...)
	data = append(data, paddedToAddress...)
	data = append(data, paddedAmount...)

	//获取nonce
	nonce, err := client.NonceAt(context.Background(), fromAddress, nil)
	if err != nil {
		return "", err
	}
	//获取小费
	gasTipCap, _ := client.SuggestGasTipCap(context.Background())
	//transfer 默认是 使用 21000 gas
	gas := uint64(100000)
	//最大gas fee
	gasFeeCap := big.NewInt(38694000460)

	contractAddress := common.HexToAddress(contract)
	//创建交易
	tx := types.NewTx(&types.DynamicFeeTx{
		Nonce:     nonce,
		GasTipCap: gasTipCap,
		GasFeeCap: gasFeeCap,
		Gas:       gas,
		To:        &contractAddress,
		Value:     big.NewInt(0),
		Data:      data,
	})
	// 获取当前区块链的ChainID
	chainID, err := client.ChainID(context.Background())
	if err != nil {
		fmt.Println("获取ChainID失败:", err)
		return "", err
	}

	fmt.Println("当前区块链的ChainID:", chainID)
	//创建签名者
	signer := types.NewLondonSigner(chainID)
	//对交易进行签名
	signTx, err := types.SignTx(tx, signer, privateKeyECDSA)
	if err != nil {
		return "", err
	}
	//发送交易
	err = client.SendTransaction(context.Background(), signTx)
	if err != nil {
		return "", err
	}
	//返回交易哈希
	return signTx.Hash().Hex(), err

}

//func main() {
//	client, err := ethclient.Dial("HTTP://127.0.0.1:7545")
//	if err != nil {
//		fmt.Println("ethclient.Dial error : ", err)
//		os.Exit(0)
//	}
//	tx, err := transfer(client, privateKey, toAddress, contractAddress)
//	if err != nil {
//		fmt.Println("transfer error : ", err)
//		os.Exit(0)
//	}
//
//	fmt.Println("transfer tx : ", tx)
//
//}
