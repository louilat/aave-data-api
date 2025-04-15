package main

import (
	"aave-data-api/events"
	"aave-data-api/prices"
	"aave-data-api/reserves"
	"aave-data-api/users"

	"github.com/gin-gonic/gin"
)

func HomePage(c *gin.Context) {
	c.JSON(200, gin.H{"message": "This API provides Aave data"})
}

func main() {
	router := gin.Default()
	router.GET("/", HomePage)
	router.GET("/reserves", reserves.GetReserves)
	router.GET("/users-balances", users.GetUsersBalances)
	router.GET("/user-selec-balances", users.GetUsersSelectionBalances)
	router.GET("/prices", prices.GetPrices)
	router.GET("/events/supply", events.GetSupply)
	router.GET("/events/borrow", events.GetBorrow)
	router.GET("/events/withdraw", events.GetWithdraw)
	router.GET("/events/repay", events.GetRepay)
	router.GET("/events/liquidation", events.GetLiquidationCall)
	router.GET("/events/reservedataupdated", events.GetReserveDataUpdated)
	router.Run("0.0.0.0:5000")
}
