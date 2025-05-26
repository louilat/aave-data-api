package users

import (
	"aave-data-api/minio"
	"aave-data-api/tps"
	"crypto/sha256"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

func GetUsersBalances(c *gin.Context) {
	creds := tps.MinioCreds{
		Endpoint:        "minio-simple.lab.groupe-genes.fr",
		AccessKeyID:     os.Getenv("ACCESS_KEY_ID"),
		SecretAccessKey: os.Getenv("SECRET_ACCESS_KEY"),
	}

	rdt := c.Query("date")
	const layout = "2006-Jan-02"
	dt, err := time.Parse(layout, rdt)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err})
	}

	if dt.Weekday() != 1 || dt.Year() > 2024 {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "No content found"})
	} else {
		key := "aave-api-datasource/daily-users-balances/users_balances_snapshot_date=" + dt.String()[:10] + "/users_balances.json"
		data, err := minio.ExtractUsersBalancesData(creds, "projet-datalab-group-jprat", key)
		if err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err})
		}
		for i := range data {
			data[i].User = fmt.Sprintf("%x", sha256.Sum256([]byte(data[i].User)))
		}
		c.IndentedJSON(http.StatusOK, data)
	}
}

func GetUsersSelectionBalances(c *gin.Context) {
	creds := tps.MinioCreds{
		Endpoint:        "minio-simple.lab.groupe-genes.fr",
		AccessKeyID:     os.Getenv("ACCESS_KEY_ID"),
		SecretAccessKey: os.Getenv("SECRET_ACCESS_KEY"),
	}

	rdt := c.Query("date")
	user := c.Query("user")

	const layout = "2006-Jan-02"
	dt, err := time.Parse(layout, rdt)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err})
	}

	if dt.Weekday() > 7 {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "No content found"})
	} else {
		key := "aave-api-datasource/daily-users-balances/users_balances_snapshot_date=" + dt.String()[:10] + "/users_balances.json"
		data, err := minio.ExtractUsersBalancesData(creds, "projet-datalab-group-jprat", key)
		if err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err})
		}

		var selection []tps.UserBalanceData
		for _, rec := range data {
			if rec.User == user {
				selection = append(selection, rec)
			}
		}

		c.IndentedJSON(http.StatusOK, selection)
	}
}
