package ynab_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"ynabconverter/pkg/reports"
	"ynabconverter/pkg/reports/cashapp"
	"ynabconverter/pkg/reports/ynab"
)

func TestYnabTransaction(t *testing.T) {
	t.Parallel()

	// 	The YNAB csv output file looks like this:

	// ```csv
	// Date,Payee,Memo,Amount
	// 10/07/2023,Cash Reward,CARD REFUNDED,1.00
	// 10/06/2023,MTA*NYCT PAYGO,CARD CHARGED,-2.90
	// 10/06/2023,MTA*NYCT PAYGO,CARD CHARGED,-2.90
	// 06/13/2023,Some business name,PAYMENT SENT,-10.00

	// ```

	t.Run("should create a ynab.Transaction from a Transaction", func(t *testing.T) {
		// Given
		transaction := reports.Transaction{
			Counterparty: "MTA*NYCT PAYGO",
			Description:  "CARD CHARGED",
			Amount:       -2.9,
			Datetime:     time.Date(2023, 10, 6, 23, 59, 59, 0, time.UTC),
		}
		expectedYnabTransaction := ynab.Transaction{
			Date:   "10/06/2023",
			Payee:  "MTA*NYCT PAYGO",
			Memo:   "CARD CHARGED",
			Amount: "-2.90",
		}
		// When
		ynabTransaction, err := ynab.NewYnabTransaction(transaction)
		// Then
		require.NoError(t, err)
		assert.Equal(t, expectedYnabTransaction, *ynabTransaction)
	})

	t.Run("should create a YnabTransaction from a cashapp.Transaction", func(t *testing.T) {
		// Given
		cashAppTransaction := cashapp.Transaction{
			TransactionID:        "rmgsrz",
			Date:                 "2023-10-06 23:59:59 EDT",
			TransactionType:      "Cash Card Debit",
			Currency:             "USD",
			Amount:               "-$2.90",
			Fee:                  "$0",
			NetAmount:            "-$2.90",
			AssetType:            "",
			AssetPrice:           "",
			AssetAmount:          "",
			Status:               "CARD CHARGED",
			Notes:                "MTA*NYCT PAYGO",
			NameOfSenderReceiver: "",
			Account:              "Visa Debit 0987",
		}
		expectedYnabTransaction := ynab.Transaction{
			Date:   "10/06/2023",
			Payee:  "MTA*NYCT PAYGO",
			Memo:   "CARD CHARGED",
			Amount: "-2.90",
		}
		// When
		ynabTransaction, err := ynab.NewYnabTransaction(&cashAppTransaction)
		// Then
		require.NoError(t, err)
		assert.Equal(t, expectedYnabTransaction, *ynabTransaction)
	})

	t.Run("should fail to create YnabTransaction if transaction.GetDatetime() fails", func(t *testing.T) {
		// Given
		cashAppNotValidDateTransaction := cashapp.Transaction{
			Date: "not a valid date",
		}
		// When
		_, err := ynab.NewYnabTransaction(&cashAppNotValidDateTransaction)
		// Then
		assert.ErrorContains(t, err, "error getting datetime")
	})

	t.Run("should fail to create YnabTransaction if transaction.GetAmount() fails", func(t *testing.T) {
		// Given
		cashAppNotValidAmountTransaction := cashapp.Transaction{
			Date:   "2023-10-06 23:59:59 EDT",
			Amount: "not a valid amount",
		}
		// When
		_, err := ynab.NewYnabTransaction(&cashAppNotValidAmountTransaction)
		// Then
		assert.ErrorContains(t, err, "error getting amount")
	})
}
