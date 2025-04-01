package minio

import (
	"aave-data-api/tps"
	"encoding/json"
	"io"

	"github.com/minio/minio-go"
)

func ExtractPricesData(creds tps.MinioCreds, bucket string, key string) ([]tps.PriceData, error) {
	useSSL := false
	minioClient, err := minio.New(creds.Endpoint, creds.AccessKeyID, creds.SecretAccessKey, useSSL)
	if err != nil {
		return make([]tps.PriceData, 0), err
	}

	obj, err := minioClient.GetObject(bucket, key, minio.GetObjectOptions{})
	if err != nil {
		return make([]tps.PriceData, 0), err
	}

	bytes, err := io.ReadAll(obj)
	if err != nil {
		return make([]tps.PriceData, 0), err
	}

	var records []tps.PriceData

	er := json.Unmarshal(bytes, &records)
	if er != nil {
		return make([]tps.PriceData, 0), er
	}

	return records, nil
}
