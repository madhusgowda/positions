package main

import (
	"positions/messages"

	"github.com/gin-gonic/gin"
)

func PositionFields(c *gin.Context) {

	chainName := c.PostForm("chainName")
	if chainName == "" {
		c.JSON(400, gin.H{
			"message": messages.ChainNameError(),
		})
		return
	}

	protocol := c.PostForm("protocol")
	if protocol == "" {
		c.JSON(400, gin.H{
			"message": messages.ProtocolError(),
		})
		return
	}

	fromAddress := c.PostForm("fromAddress")
	if fromAddress == "" {
		c.JSON(400, gin.H{
			"message": messages.FromAddressError(),
		})
		return
	}

	currency := c.PostForm("currency")
	if currency == "" {
		c.JSON(400, gin.H{
			"message": messages.CurrencyError(),
		})
		return
	}


}
