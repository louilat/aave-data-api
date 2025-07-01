package reserves

import (
	"aave-data-api/minio"
	"aave-data-api/tps"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

func GetReserves(c *gin.Context) {
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
		key := "aavev3-raw-datasource /daily-users-balances /users_balances_snapshot_date=" + dt.String()[:10] + "/reserves_data.json"
		data, err := minio.ExtractReservesData(creds, "projet-datalab-group-jprat", key)
		if err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err})
		}

		c.IndentedJSON(http.StatusOK, data)
	}
}
