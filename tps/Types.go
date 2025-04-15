package tps

import "math/big"

type MinioCreds struct {
	Endpoint        string
	AccessKeyID     string
	SecretAccessKey string
}

// type AaveV3Data interface {
// 	DataName()
// }

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

type PriceData struct {
	Block     string  `json:"BlockNumber"`
	Timestamp big.Int `json:"Timestamp"`
	Asset     string  `json:"UnderlyingToken"`
	Price     big.Int `json:"Price"`
}

type LiquidationCallEvent struct {
	Block            big.Int `json:"blockNumber"`
	CollateralAsset  string  `json:"collateralAsset"`
	DebtAsset        string  `json:"debtAsset"`
	User             string  `json:"user"`
	DebtToCover      big.Int `json:"debtToCover"`
	CollateralAmount big.Int `json:"liquidatedCollateralAmount"`
	Liquidator       string  `json:"liquidator"`
	ReceiveAToken    byte    `json:"receiveAToken"`
}

type SupplyEvent struct {
	Block      big.Int `json:"blockNumber"`
	Reserve    string  `json:"reserve"`
	OnBehalfOf string  `json:"onBehalfOf"`
	User       string  `json:"user"`
	Amount     big.Int `json:"amount"`
}

type BorrowEvent struct {
	Block      big.Int `json:"blockNumber"`
	Reserve    string  `json:"reserve"`
	OnBehalfOf string  `json:"onBehalfOf"`
	User       string  `json:"user"`
	Amount     big.Int `json:"amount"`
	BorrowRate big.Int `json:"borrowRate"`
}

type WithdrawEvent struct {
	Block   big.Int `json:"blockNumber"`
	Reserve string  `json:"reserve"`
	To      string  `json:"to"`
	User    string  `json:"user"`
	Amount  big.Int `json:"amount"`
}

type RepayEvent struct {
	Block      big.Int `json:"blockNumber"`
	Reserve    string  `json:"reserve"`
	User       string  `json:"user"`
	Repayer    string  `json:"repayer"`
	Amount     big.Int `json:"amount"`
	UseATokens big.Int `json:"useATokens"`
}

type ReserveDataUpadedEvent struct {
	Block               big.Int `json:"blockNumber"`
	Reserve             string  `json:"reserve"`
	LiquidityRate       string  `json:"liquidityRate"`
	VariableBorrowRate  string  `json:"variableBorrowRate"`
	LiquidityIndex      string  `json:"liquidityIndex"`
	VariableBorrowIndex string  `json:"variableBorrowIndex"`
}
