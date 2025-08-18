package minio

import (
	"aave-data-api/tps"
	"encoding/json"
	"io"

	"github.com/minio/minio-go"
)

func ExtractUsersBalancesData(creds tps.MinioCreds, bucket string, key string) ([]tps.UserBalanceData, error) {
	useSSL := true
	minioClient, err := minio.New(creds.Endpoint, creds.AccessKeyID, creds.SecretAccessKey, useSSL)
	if err != nil {
		return make([]tps.UserBalanceData, 0), err
	}

	obj, err := minioClient.GetObject(bucket, key, minio.GetObjectOptions{})
	if err != nil {
		return make([]tps.UserBalanceData, 0), err
	}

	bytes, err := io.ReadAll(obj)
	if err != nil {
		return make([]tps.UserBalanceData, 0), err
	}

	var records []tps.UserBalanceData

	err = json.Unmarshal(bytes, &records)
	if err != nil {
		return make([]tps.UserBalanceData, 0), err
	}

	return records, nil
}
