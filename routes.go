package main

import (
	"positions/messages"

	"github.com/alexsasharegan/dotenv"
	"github.com/chenzhijie/go-web3"
	"github.com/gin-gonic/gin"
	"github.com/onrik/ethrpc"
)

func homePage(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello World",
	})
}

func getPositions(c *gin.Context, rpc_provider string, contractAddress string, currency string, abiData string){
	err := dotenv.Load()
	if err != nil {
		c.JSON(400, gin.H{
			"message": messages.EnvError(),
		})
		return
	}

	web3, err := web3.NewWeb3(rpc_provider)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Web3 error:" + err.Error(),
		})
		return
	}

	client := ethrpc.New(rpc_provider)

	//create a new contract instance
	contract, err := web3.Eth.NewContract(abiData)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Contract error:" + err.Error(),
		})
		return
	}

	stakedContractAddress := contractAddress

	stakeTokenName, err := contract.Call("name")
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Contract error:" + err.Error(),
		})
		return
	}

	stakeTokenSymbol, err := contract.Call("symbol")
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Contract error:" + err.Error(),
		})
		return
	}
	
	stakeTokenDecimals, err := contract.Call("decimals")
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Contract error:" + err.Error(),
		})
		return
	}

	rewardTokenName, err := contract.Call("rewardTokenName")
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Contract error:" + err.Error(),
		})
		return
	}

	rewardTokenSymbol, err := contract.Call("rewardTokenSymbol")
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Contract error:" + err.Error(),
		})
		return
	}

	rewardTokenDecimals, err := contract.Call("rewardTokenDecimals")
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Contract error:" + err.Error(),
		})
		return
	}

	
}
