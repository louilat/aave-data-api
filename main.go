package main

import (
	"aave-data-api/reserves"
	"aave-data-api/users"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/reserves", reserves.GetReserves)
	router.GET("/users-balances/:date", users.GetUsersBalances)
	router.Run("0.0.0.0:5000")
}
