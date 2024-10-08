package ynab

import (
	"fmt"

	"ynabconverter/pkg/reports"
)

type TransactionToRecordTransformer struct {
	header []string
}

func (ynab TransactionToRecordTransformer) GetHeader() []string {
	return ynab.header
}

func (
	ynab TransactionToRecordTransformer,
) GetRecords(transactions []reports.Transactioner) ([][]string, error) {
	records := [][]string{}
	for _, transaction := range transactions {
		ynabTransaction, err := NewYnabTransaction(transaction)
		if err != nil {
			return nil, fmt.Errorf("error creating ynab transaction: %w", err)
		}

		records = append(records, []string{
			ynabTransaction.Date,
			ynabTransaction.Payee,
			ynabTransaction.Memo,
			ynabTransaction.Amount,
		})
	}

	return records, nil
}

func (
	ynab TransactionToRecordTransformer,
) GetRecordsWithHeader(transactions []reports.Transactioner) ([][]string, error) {
	records, err := ynab.GetRecords(transactions)
	if err != nil {
		return nil, fmt.Errorf("error getting records: %w", err)
	}

	return append([][]string{ynab.GetHeader()}, records...), nil
}

func NewYnabRecordTransformer() TransactionToRecordTransformer {
	return TransactionToRecordTransformer{
		header: []string{"Date", "Payee", "Memo", "Amount"},
	}
}
