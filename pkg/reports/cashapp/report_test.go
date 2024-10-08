package cashapp_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"ynabconverter/internal/file"
	"ynabconverter/pkg/reports"
	"ynabconverter/pkg/reports/cashapp"
	utils_test "ynabconverter/tests/utils"
)

func TestCashAppReportImporter(t *testing.T) {
	t.Parallel()

	t.Run("should return error when file does not exist", func(t *testing.T) {
		// Given
		cashAppReport := cashapp.NewCashAppReport(file.NewCsvImporterFromFileSytem(utils_test.ExampleFilesFS))
		// When
		err := cashAppReport.ParseFileRecords("examples/does-not-exist.csv")
		// Then
		require.Error(t, err)
	})

	t.Run("should parse file and get an array of cashapp.Transaction", func(t *testing.T) {
		// Given
		cashAppReport := cashapp.NewCashAppReport(file.NewCsvImporterFromFileSytem(utils_test.ExampleFilesFS))
		// When
		err := cashAppReport.ParseFileRecords("examples/cash_app_report_sample.csv")
		transactions := cashAppReport.GetTransactions()
		// Then
		require.NoError(t, err)
		expectedCashAppTransaction := cashapp.Transaction{
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
		assert.Equal(t, &expectedCashAppTransaction, transactions[0])
	})

	t.Run("should implement report.ReportImporter interface", func(t *testing.T) {
		// Given
		cashAppReport := cashapp.NewCashAppReport(file.NewCsvImporterFromFileSytem(utils_test.ExampleFilesFS))
		// When
		_, implementsInterface := interface{}(&cashAppReport).(reports.ReportImporter)
		// Then
		assert.True(t, implementsInterface, "CashAppReport does not implement ReportImporter interface")
	})
}
