package events

import (
	"aave-data-api/minio"
	"aave-data-api/tps"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

func GetSupply(c *gin.Context) {
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

	if dt.Weekday() > 7 {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "No content found"})
	} else {
		key := "aave-api-datasource/daily-decoded-events/decoded_events_snapshot_date=" + dt.String()[:10] + "/decoded_Supply.json"
		data, err := minio.ExtractSupplyData(creds, "projet-datalab-group-jprat", key)
		if err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err})
		}

		c.IndentedJSON(http.StatusOK, data)
	}
}

func GetBorrow(c *gin.Context) {
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

	if dt.Weekday() > 7 {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "No content found"})
	} else {
		key := "aave-api-datasource/daily-decoded-events/decoded_events_snapshot_date=" + dt.String()[:10] + "/decoded_Borrow.json"
		data, err := minio.ExtractBorrowData(creds, "projet-datalab-group-jprat", key)
		if err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err})
		}

		c.IndentedJSON(http.StatusOK, data)
	}
}

func GetWithdraw(c *gin.Context) {
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

	if dt.Weekday() > 7 {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "No content found"})
	} else {
		key := "aave-api-datasource/daily-decoded-events/decoded_events_snapshot_date=" + dt.String()[:10] + "/decoded_Withdraw.json"
		data, err := minio.ExtractWithdrawData(creds, "projet-datalab-group-jprat", key)
		if err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err})
		}

		c.IndentedJSON(http.StatusOK, data)
	}
}

func GetRepay(c *gin.Context) {
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

	if dt.Weekday() > 7 {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "No content found"})
	} else {
		key := "aave-api-datasource/daily-decoded-events/decoded_events_snapshot_date=" + dt.String()[:10] + "/decoded_Repay.json"
		data, err := minio.ExtractRepayData(creds, "projet-datalab-group-jprat", key)
		if err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err})
		}

		c.IndentedJSON(http.StatusOK, data)
	}
}

func GetLiquidationCall(c *gin.Context) {
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

	if dt.Weekday() > 7 {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "No content found"})
	} else {
		key := "aave-api-datasource/daily-decoded-events/decoded_events_snapshot_date=" + dt.String()[:10] + "/decoded_LiquidationCall.json"
		data, err := minio.ExtractLiquidationCallData(creds, "projet-datalab-group-jprat", key)
		if err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err})
		}

		c.IndentedJSON(http.StatusOK, data)
	}
}
