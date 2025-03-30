package users

import (
	"aave-data-api/minio"
	"aave-data-api/tps"
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

	rdt := c.Param("date")
	const layout = "2006-Jan-02"
	dt, err := time.Parse(layout, rdt)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err})
	}

	if dt.Weekday() != 1 {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "No content found"})
	} else {
		key := "aave-api-datasource/daily-users-balances/users_balances_snapshot_date=" + dt.String()[:10] + "/users_balances.json"
		data, err := minio.ExtractUsersBalancesData(creds, "projet-datalab-group-jprat", key)
		if err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err})
		}

		c.IndentedJSON(http.StatusOK, data)
	}
}
