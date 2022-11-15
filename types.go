// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package main

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/shopspring/decimal"
	"gorm.io/datatypes"
	"math/big"
	"sync"
	"time"
)

var (
	_ = decimal.Decimal{}
	_ = big.NewInt
	_ = datatypes.JSON{}
	_ = time.Time{}
)

func GetApprovalEventHash() string {
	return "0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925"
}

type ApprovalEvent struct {
	EventOwner   string
	EventSpender string
	EventValue   decimal.Decimal `gorm:"type:numeric"`

	ID              uint   `gorm:"primaryKey"`
	BlockNumber     uint64 `gorm:"uniqueIndex:57f4aeb2-68c2-4218-a92d-e186314d6ed1,unique;index"`
	TxHash          string
	TxIndex         uint `gorm:"uniqueIndex:57f4aeb2-68c2-4218-a92d-e186314d6ed1,unique"`
	BlockHash       string
	Gas             decimal.Decimal `gorm:"type:numeric"`
	GasPrice        decimal.Decimal `gorm:"type:numeric"`
	TxFrom          string          `gorm:"index"`
	TxTo            string          `gorm:"index"`
	TxValue         decimal.Decimal `gorm:"type:numeric"`
	Index           uint            `gorm:"uniqueIndex:57f4aeb2-68c2-4218-a92d-e186314d6ed1,unique"`
	BlockTime       time.Time       `gorm:"index"`
	ContractAddress string
	ChainID         string
}

func GetTransferEventHash() string {
	return "0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"
}

type TransferEvent struct {
	EventFrom  string
	EventTo    string
	EventValue decimal.Decimal `gorm:"type:numeric"`

	ID              uint   `gorm:"primaryKey"`
	BlockNumber     uint64 `gorm:"uniqueIndex:15388172-c42f-40d2-87cc-6a6c072c6478,unique;index"`
	TxHash          string
	TxIndex         uint `gorm:"uniqueIndex:15388172-c42f-40d2-87cc-6a6c072c6478,unique"`
	BlockHash       string
	Gas             decimal.Decimal `gorm:"type:numeric"`
	GasPrice        decimal.Decimal `gorm:"type:numeric"`
	TxFrom          string          `gorm:"index"`
	TxTo            string          `gorm:"index"`
	TxValue         decimal.Decimal `gorm:"type:numeric"`
	Index           uint            `gorm:"uniqueIndex:15388172-c42f-40d2-87cc-6a6c072c6478,unique"`
	BlockTime       time.Time       `gorm:"index"`
	ContractAddress string
	ChainID         string
}

func GetTransferMethodHash() string {
	return "a9059cbb"
}

type TransferMethod struct {
	MethodRecipient string
	MethodAmount    decimal.Decimal `gorm:"type:numeric"`

	ID              uint   `gorm:"primaryKey"`
	BlockNumber     uint64 `gorm:"uniqueIndex:f7855d0b-97c5-4ae9-a9e1-a1a4b4a54cb0,unique;index"`
	TxHash          string
	TxIndex         uint `gorm:"uniqueIndex:f7855d0b-97c5-4ae9-a9e1-a1a4b4a54cb0,unique"`
	BlockHash       string
	Gas             decimal.Decimal `gorm:"type:numeric"`
	GasPrice        decimal.Decimal `gorm:"type:numeric"`
	TxFrom          string          `gorm:"index"`
	TxTo            string          `gorm:"index"`
	TxValue         decimal.Decimal `gorm:"type:numeric"`
	BlockTime       time.Time       `gorm:"index"`
	ContractAddress string
	ChainID         string
}

func GetApproveMethodHash() string {
	return "095ea7b3"
}

type ApproveMethod struct {
	MethodSpender string
	MethodAmount  decimal.Decimal `gorm:"type:numeric"`

	ID              uint   `gorm:"primaryKey"`
	BlockNumber     uint64 `gorm:"uniqueIndex:b7b25620-8b2c-47ad-a78b-3b65f81d3dc4,unique;index"`
	TxHash          string
	TxIndex         uint `gorm:"uniqueIndex:b7b25620-8b2c-47ad-a78b-3b65f81d3dc4,unique"`
	BlockHash       string
	Gas             decimal.Decimal `gorm:"type:numeric"`
	GasPrice        decimal.Decimal `gorm:"type:numeric"`
	TxFrom          string          `gorm:"index"`
	TxTo            string          `gorm:"index"`
	TxValue         decimal.Decimal `gorm:"type:numeric"`
	BlockTime       time.Time       `gorm:"index"`
	ContractAddress string
	ChainID         string
}

type LastSyncedBlock struct {
	Contract    string `gorm:"primaryKey"`
	ChainID     string `gorm:"primaryKey"`
	SyncType    string `gorm:"primaryKey"`
	BlockNumber uint64
}

// Plugin Models
type TokenDetails struct {
	ID      int
	Address string `gorm:"uniqueIndex:address_and_chain"`
	Symbol  string
	ChainID string `gorm:"uniqueIndex:address_and_chain"`
	Decimal int
}

var tokenCache = sync.Map{}

// Config
type PostgresConfig struct {
	ConnectionString string `mapstructure:"connection_string"`
	TablePrefix      string `mapstructure:"table_prefix"`
	CreateBatchSize  int    `mapstructure:"create_batch_size"`
}

type IndexerConfig struct {
	EthEndpoint       string `mapstructure:"eth_endpoint"`
	ContractAddress   string `mapstructure:"contract_address"`
	StartBlock        int    `mapstructure:"start_block"`
	ApiKey            string `mapstructure:"api_key"`
	PostgresConfig    `mapstructure:"postgres_config"`
	LagToHighestBlock int `mapstructure:"lag_to_highest_block"`
	StepSize          int `mapstructure:"step_size"`
	ParallelCalls     int `mapstructure:"parallel_calls_for_logs"`
	PrometheusConfig  `mapstructure:"prometheus_config"`
	Metrics           *Metrics
}

type PrometheusConfig struct {
	PrometheusNameSpace string //PrometheusNameSpace The prometheus namespace metrics will be reported on. Default -> custom_indexer
	PrometheusSubsystem string //PrometheusSubsystem The prometheus subsystem metrics will be reported on. Default -> parser
	MetricsPort         int    `mapstructure:"metrics_port"`    //MetricsPort port metrics is scraped on. Default -> 9091
	PrometheusHost      string `mapstructure:"prometheus_host"` //PrometheusHost Host for the metrics server. Default -> 0.0.0.0
}

type Metrics struct {
	//List of metrics
	NodeRestarts         *prometheus.CounterVec   //NodeRestarts node restarts specific to this parser
	NodeLatency          *prometheus.HistogramVec //NodeLatency Response latency for node
	UnmarshalAPIRestarts *prometheus.CounterVec   //UnmarshalAPIRestarts api gateway restarts specific to this parser
	UnmarshalAPILatency  *prometheus.HistogramVec //UnmarshalAPILatency Response Latency for unmarshal API calls

}

func (i *IndexerConfig) AssignDefaults() {
	if i.PostgresConfig.CreateBatchSize == 0 {
		i.PostgresConfig.CreateBatchSize = 100
	}
	if i.StepSize == 0 {
		i.StepSize = 50
	}
	if i.LagToHighestBlock == 0 {
		i.LagToHighestBlock = 10
	}
	if i.ParallelCalls == 0 {
		i.ParallelCalls = 1
	}
	if i.MetricsPort == 0 {
		i.MetricsPort = 9091
	}
	if i.PrometheusHost == "" {
		i.PrometheusHost = "0.0.0.0"
	}

	i.PrometheusNameSpace = "custom_indexer"
	i.PrometheusSubsystem = "parser"
}
