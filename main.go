package main

import (
	"aave-data-api/reserves"
	"aave-data-api/users"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/reserves/:date", reserves.GetReserves)
	router.GET("/users-balances/:date", users.GetUsersBalances)
	router.Run("localhost:8080")
}
