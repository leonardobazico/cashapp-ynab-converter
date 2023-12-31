package reports

import (
	"time"
)

type ReportImporter interface {
	ParseFileRecords(filePath string) error
	GetTransactions() []Transactioner
}

type TransactionToRecordTransformer interface {
	GetHeader() []string
	GetRecords(transactions []Transactioner) ([][]string, error)
	GetRecordsWithHeader(transactions []Transactioner) ([][]string, error)
}

type Transactioner interface {
	GetCounterparty() string
	GetDescription() string
	GetAmount() (float32, error)
	GetDatetime() (*time.Time, error)
}

type Transaction struct {
	Counterparty string
	Description  string
	Amount       float32
	Datetime     time.Time
}

func (transaction Transaction) GetCounterparty() string {
	return transaction.Counterparty
}

func (transaction Transaction) GetDescription() string {
	return transaction.Description
}

func (transaction Transaction) GetAmount() (float32, error) {
	return transaction.Amount, nil
}

func (transaction Transaction) GetDatetime() (*time.Time, error) {
	return &transaction.Datetime, nil
}
