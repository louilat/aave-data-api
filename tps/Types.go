package tps

import "math/big"

type MinioCreds struct {
	Endpoint        string
	AccessKeyID     string
	SecretAccessKey string
}

type ReserveData struct {
	Name                 string  `json:"name"`
	LastUpdateTimestamp  int64   `json:"lastUpdateTimestamp"`
	Decimals             int64   `json:"decimals"`
	Asset                string  `json:"underlyingAsset"`
	PriceUSD             float64 `json:"underlyingTokenPriceUSD"`
	LiquidityRate        string  `json:"liquidityRate"`
	VariableBorrowRate   string  `json:"variableBorrowRate"`
	LiquidityIndex       string  `json:"liquidityIndex"`
	VariableBorrowIndex  string  `json:"variableBorrowIndex"`
	LiquidationThreshold int64   `json:"reserveLiquidationThreshold"`
	AvailableLiquidity   string  `json:"availableLiquidity"`
	TotalVariableDebt    string  `json:"totalScaledVariableDebt"`
}

type UserBalanceData struct {
	Block               big.Int `json:"snapshot_block"`
	User                string  `json:"user_address"`
	Name                string  `json:"name"`
	Asset               string  `json:"underlyingAsset"`
	Decimals            big.Int `json:"decimals"`
	ScaledATokenBalance big.Int `json:"scaledATokenBalance"`
	ScaledVariableDebt  big.Int `json:"scaledVariableDebt"`
}
