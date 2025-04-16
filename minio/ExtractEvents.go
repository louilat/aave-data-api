package minio

import (
	"aave-data-api/tps"
	"encoding/json"
	"io"

	"github.com/minio/minio-go"
)

func ExtractSupplyData(creds tps.MinioCreds, bucket string, key string) ([]tps.SupplyEvent, error) {
	useSSL := false
	minioClient, err := minio.New(creds.Endpoint, creds.AccessKeyID, creds.SecretAccessKey, useSSL)
	if err != nil {
		return make([]tps.SupplyEvent, 0), err
	}

	obj, err := minioClient.GetObject(bucket, key, minio.GetObjectOptions{})
	if err != nil {
		return make([]tps.SupplyEvent, 0), err
	}

	bytes, err := io.ReadAll(obj)
	if err != nil {
		return make([]tps.SupplyEvent, 0), err
	}

	var records []tps.SupplyEvent

	er := json.Unmarshal(bytes, &records)
	if er != nil {
		return make([]tps.SupplyEvent, 0), er
	}

	return records, nil
}

func ExtractBorrowData(creds tps.MinioCreds, bucket string, key string) ([]tps.BorrowEvent, error) {
	useSSL := false
	minioClient, err := minio.New(creds.Endpoint, creds.AccessKeyID, creds.SecretAccessKey, useSSL)
	if err != nil {
		return make([]tps.BorrowEvent, 0), err
	}

	obj, err := minioClient.GetObject(bucket, key, minio.GetObjectOptions{})
	if err != nil {
		return make([]tps.BorrowEvent, 0), err
	}

	bytes, err := io.ReadAll(obj)
	if err != nil {
		return make([]tps.BorrowEvent, 0), err
	}

	var records []tps.BorrowEvent

	er := json.Unmarshal(bytes, &records)
	if er != nil {
		return make([]tps.BorrowEvent, 0), er
	}

	return records, nil
}

func ExtractWithdrawData(creds tps.MinioCreds, bucket string, key string) ([]tps.WithdrawEvent, error) {
	useSSL := false
	minioClient, err := minio.New(creds.Endpoint, creds.AccessKeyID, creds.SecretAccessKey, useSSL)
	if err != nil {
		return make([]tps.WithdrawEvent, 0), err
	}

	obj, err := minioClient.GetObject(bucket, key, minio.GetObjectOptions{})
	if err != nil {
		return make([]tps.WithdrawEvent, 0), err
	}

	bytes, err := io.ReadAll(obj)
	if err != nil {
		return make([]tps.WithdrawEvent, 0), err
	}

	var records []tps.WithdrawEvent

	er := json.Unmarshal(bytes, &records)
	if er != nil {
		return make([]tps.WithdrawEvent, 0), er
	}

	return records, nil
}

func ExtractRepayData(creds tps.MinioCreds, bucket string, key string) ([]tps.RepayEvent, error) {
	useSSL := false
	minioClient, err := minio.New(creds.Endpoint, creds.AccessKeyID, creds.SecretAccessKey, useSSL)
	if err != nil {
		return make([]tps.RepayEvent, 0), err
	}

	obj, err := minioClient.GetObject(bucket, key, minio.GetObjectOptions{})
	if err != nil {
		return make([]tps.RepayEvent, 0), err
	}

	bytes, err := io.ReadAll(obj)
	if err != nil {
		return make([]tps.RepayEvent, 0), err
	}

	var records []tps.RepayEvent

	er := json.Unmarshal(bytes, &records)
	if er != nil {
		return make([]tps.RepayEvent, 0), er
	}

	return records, nil
}

func ExtractLiquidationCallData(creds tps.MinioCreds, bucket string, key string) ([]tps.LiquidationCallEvent, error) {
	useSSL := false
	minioClient, err := minio.New(creds.Endpoint, creds.AccessKeyID, creds.SecretAccessKey, useSSL)
	if err != nil {
		return make([]tps.LiquidationCallEvent, 0), err
	}

	obj, err := minioClient.GetObject(bucket, key, minio.GetObjectOptions{})
	if err != nil {
		return make([]tps.LiquidationCallEvent, 0), err
	}

	bytes, err := io.ReadAll(obj)
	if err != nil {
		return make([]tps.LiquidationCallEvent, 0), err
	}

	var records []tps.LiquidationCallEvent

	er := json.Unmarshal(bytes, &records)
	if er != nil {
		return make([]tps.LiquidationCallEvent, 0), er
	}

	return records, nil
}

func ExtractReserveDataUpdatedData(creds tps.MinioCreds, bucket string, key string) ([]tps.ReserveDataUpadedEvent, error) {
	useSSL := false
	minioClient, err := minio.New(creds.Endpoint, creds.AccessKeyID, creds.SecretAccessKey, useSSL)
	if err != nil {
		return make([]tps.ReserveDataUpadedEvent, 0), err
	}

	obj, err := minioClient.GetObject(bucket, key, minio.GetObjectOptions{})
	if err != nil {
		return make([]tps.ReserveDataUpadedEvent, 0), err
	}

	bytes, err := io.ReadAll(obj)
	if err != nil {
		return make([]tps.ReserveDataUpadedEvent, 0), err
	}

	var records []tps.ReserveDataUpadedEvent

	er := json.Unmarshal(bytes, &records)
	if er != nil {
		return make([]tps.ReserveDataUpadedEvent, 0), er
	}

	return records, nil
}

func ExtractBalanceTransferData(creds tps.MinioCreds, bucket string, key string) ([]tps.BalanceTransferEvent, error) {
	useSSL := false
	minioClient, err := minio.New(creds.Endpoint, creds.AccessKeyID, creds.SecretAccessKey, useSSL)
	if err != nil {
		return make([]tps.BalanceTransferEvent, 0), err
	}

	obj, err := minioClient.GetObject(bucket, key, minio.GetObjectOptions{})
	if err != nil {
		return make([]tps.BalanceTransferEvent, 0), err
	}

	bytes, err := io.ReadAll(obj)
	if err != nil {
		return make([]tps.BalanceTransferEvent, 0), err
	}

	var records []tps.BalanceTransferEvent

	er := json.Unmarshal(bytes, &records)
	if er != nil {
		return make([]tps.BalanceTransferEvent, 0), er
	}

	return records, nil
}
