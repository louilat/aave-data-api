package tps

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type MinioCreds struct {
	Endpoint        string
	AccessKeyID     string
	SecretAccessKey string
}

type ReserveData struct {
	BlockNumber             *big.Int `json:"blockNumber"`
	Name                    string   `json:"name"`
	UnderlyingAsset         string   `json:"underlyingAsset"`
	Decimals                uint8    `json:"decimals"`
	UnderlyingTokenPriceUSD *big.Int `json:"underlyingTokenPriceUSD"`
	ScaledTotalLiquidity    *big.Int `json:"scaledTotalLiquidity"`
	ScaledTotalVariableDebt *big.Int `json:"scaledTotalVariableDebt"`
	AvailableLiquidity      *big.Int `json:"availableLiquidity"`
	TreasuryAmount          *big.Int `json:"treasuryAmount"`

	Configuration               *big.Int       `json:"configuration"`
	LiquidityIndex              *big.Int       `json:"liquidityIndex"`
	CurrentLiquidityRate        *big.Int       `json:"currentLiquidityRate"`
	VariableBorrowIndex         *big.Int       `json:"variableBorrowIndex"`
	CurrentVariableBorrowRate   *big.Int       `json:"currentVariableBorrowRate"`
	CurrentStableBorrowRate     *big.Int       `json:"currentStableBorrowRate"`
	LastUpdateTimestamp         *big.Int       `json:"lastUpdateTimestamp"`
	Id                          uint16         `json:"id"`
	ATokenAddress               common.Address `json:"aTokenAddress"`
	StableDebtTokenAddress      common.Address `json:"stableDebtTokenAddress"`
	VariableDebtTokenAddress    common.Address `json:"variableDebtTokenAddress"`
	InterestRateStrategyAddress common.Address `json:"interestRateStrategyAddress"`
	AccruedToTreasury           *big.Int       `json:"accruedToTreasury"`
	Unbacked                    *big.Int       `json:"unbacked"`
	IsolationModeTotalDebt      *big.Int       `json:"isolationModeTotalDebt"`

	Ltv                    int64 `json:"ltv"`
	LiquidationThreshold   int64 `json:"liquidationThreshold"`
	LiquidationBonus       int64 `json:"liquidationBonus"`
	ReserveFactor          int64 `json:"reserveFactor"`
	BorrowCap              int64 `json:"borrowCap"`
	SupplyCap              int64 `json:"supplyCap"`
	LiquidationProtocolFee int64 `json:"liquidationProtocolFee"`
	EModeCategory          int64 `json:"eModeCategory"`
}

type UserBalanceData struct {
	User                    string  `json:"user"`
	BlockNumber             float64 `json:"blockNumber"`
	TokenName               string  `json:"tokenName"`
	UnderlyingAsset         string  `json:"underlyingAsset"`
	Decimals                float64 `json:"decimals"`
	ScaledATokenBalance     float64 `json:"scaledATokenBalance"`
	ScaledVariableDebt      float64 `json:"scaledVariableDebt"`
	UnderlyingTokenPriceUSD float64 `json:"underlyingTokenPriceUSD"`

	TotalCollateralBase         float64 `json:"totalCollateralBase"`
	TotalDebtBase               float64 `json:"totalDebtBase"`
	CurrentLiquidationThreshold float64 `json:"currentLiquidationThreshold"`
	Ltv                         float64 `json:"ltv"`
	HealthFactor                float64 `json:"healthFactor"`
	UserEMode                   float64 `json:"userEMode"`
}

// type PriceData struct {
// 	Block     string  `json:"BlockNumber"`
// 	Timestamp big.Int `json:"Timestamp"`
// 	Asset     string  `json:"UnderlyingToken"`
// 	Price     big.Int `json:"Price"`
// }

type LiquidationCallEvent struct {
	Block            big.Int `json:"blockNumber"`
	CollateralAsset  string  `json:"collateralAsset"`
	DebtAsset        string  `json:"debtAsset"`
	User             string  `json:"user"`
	DebtToCover      big.Int `json:"debtToCover"`
	CollateralAmount big.Int `json:"liquidatedCollateralAmount"`
	Liquidator       string  `json:"liquidator"`
	// ReceiveAToken    byte    `json:"receiveAToken"`
}

type SupplyEvent struct {
	Address     string   `json:"Address"`
	BlockNumber *big.Int `json:"BlockNumber"`
	TxHash      string   `json:"TxHash"`
	Reserve     string   `json:"Reserve"`
	OnBehalfOf  string   `json:"OnBehalfOf"`
	User        string   `json:"User"`
	Amount      big.Int  `json:"Amount"`
}

type BorrowEvent struct {
	Address          string   `json:"Address"`
	BlockNumber      *big.Int `json:"BlockNumber"`
	TxHash           string   `json:"TxHash"`
	Reserve          string   `json:"Reserve"`
	User             string   `json:"User"`
	OnBehalfOf       string   `json:"OnBehalfOf"`
	Amount           *big.Int `json:"Amount"`
	InterestRateMode *big.Int `json:"InterestRateMode"`
	BorrowRate       *big.Int `json:"BorrowRate"`
}

type WithdrawEvent struct {
	Address     string   `json:"Address"`
	BlockNumber *big.Int `json:"BlockNumber"`
	TxHash      string   `json:"TxHash"`
	Reserve     string   `json:"Reserve"`
	To          string   `json:"To"`
	User        string   `json:"User"`
	Amount      big.Int  `json:"Amount"`
}

type RepayEvent struct {
	Address     string   `json:"Address"`
	BlockNumber *big.Int `json:"BlockNumber"`
	TxHash      string   `json:"TxHash"`
	Reserve     string   `json:"Reserve"`
	User        string   `json:"User"`
	Repayer     string   `json:"Repayer"`
	Amount      big.Int  `json:"Amount"`
	UseATokens  big.Int  `json:"UseATokens"`
}
