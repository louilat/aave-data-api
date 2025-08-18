package minio

import (
	"aave-data-api/tps"
	"encoding/json"
	"fmt"
	"io"

	"github.com/minio/minio-go"
)

func ExtractReservesData(creds tps.MinioCreds, bucket string, key string) ([]tps.ReserveData, error) {
	useSSL := true
	minioClient, err := minio.New(creds.Endpoint, creds.AccessKeyID, creds.SecretAccessKey, useSSL)
	if err != nil {
		fmt.Println(err.Error())
		return make([]tps.ReserveData, 0), err
	}

	obj, err := minioClient.GetObject(bucket, key, minio.GetObjectOptions{})
	if err != nil {
		fmt.Println(err.Error())
		return make([]tps.ReserveData, 0), err
	}

	bytes, err := io.ReadAll(obj)
	if err != nil {
		fmt.Println(err.Error())
		return make([]tps.ReserveData, 0), err
	}

	var records []tps.ReserveData

	err = json.Unmarshal(bytes, &records)
	if err != nil {
		fmt.Println(err.Error())
		return make([]tps.ReserveData, 0), err
	}

	return records, nil
}
