# 本工程主要演示用go调用智能合约三种方法

## part1：使用函数选择器方式进行调用
## part2：使用abi方式进行调用
## part3：使用abigen生成合约go文件进行调用




/Users/chenhao/Downloads/web3/geth-alltools/abigen --abi Calculator.abi -type Calculator -pkg main -out calculator.go



/Users/chenhao/Downloads/web3/geth-alltools/abigen --abi CrowdFunding.abi -type CrowdFund -pkg main -out CrowdFunding.go