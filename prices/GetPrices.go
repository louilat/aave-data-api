package prices

// func GetPrices(c *gin.Context) {
// 	creds := tps.MinioCreds{
// 		Endpoint:        "minio-simple.lab.groupe-genes.fr",
// 		AccessKeyID:     os.Getenv("ACCESS_KEY_ID"),
// 		SecretAccessKey: os.Getenv("SECRET_ACCESS_KEY"),
// 	}

// 	rdt := c.Query("date")
// 	const layout = "2006-Jan-02"
// 	dt, err := time.Parse(layout, rdt)
// 	if err != nil {
// 		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err})
// 	}

// 	if dt.Weekday() > 7 {
// 		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "No content found"})
// 	} else {
// 		key := "aave-raw-datasource/hourly-prices/hourly_prices_snapshot_date=" + dt.String()[:10] + "/hourly_prices.json"
// 		data, err := minio.ExtractPricesData(creds, "projet-datalab-group-jprat", key)
// 		if err != nil {
// 			c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err})
// 		}

// 		c.IndentedJSON(http.StatusOK, data)
// 	}
// }
