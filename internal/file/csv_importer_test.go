package file_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"ynabconverter/internal/file"
	utils_test "ynabconverter/tests/utils"
)

func TestCsvImporter(t *testing.T) {
	t.Parallel()

	t.Run("should return error when file does not exist", func(t *testing.T) {
		// Given
		filePath := "examples/does-not-exist.csv"
		csvReader := file.NewCsvImporterFromFileSytem(utils_test.ExampleFilesFS)
		// When
		_, err := csvReader.GetRecordsFrom(filePath)
		// Then
		assert.EqualError(
			t,
			err,
			"fail to open file: open examples/does-not-exist.csv: "+
				"file does not exist",
		)
	})

	t.Run("should return error when file exists but is not a csv", func(t *testing.T) {
		// Given
		filePath := "examples/not-a-csv.txt"
		csvReader := file.NewCsvImporterFromFileSytem(utils_test.ExampleFilesFS)
		// When
		output, err := csvReader.GetRecordsFrom(filePath)
		// Then
		assert.Nil(t, output)
		assert.ErrorContains(t, err, "fail to read csv file:")
	})

	t.Run("should not return error when file exists", func(t *testing.T) {
		// Given
		filePath := "examples/cash_app_report.csv"
		csvReader := file.NewCsvImporterFromFileSytem(utils_test.ExampleFilesFS)
		// When
		_, err := csvReader.GetRecordsFrom(filePath)
		// Then
		require.NoError(t, err)
	})

	t.Run("should not return error when file exists and is empty", func(t *testing.T) {
		// Given
		filePath := "examples/empty.csv"
		csvReader := file.NewCsvImporterFromFileSytem(utils_test.ExampleFilesFS)
		// When
		records, err := csvReader.GetRecordsFrom(filePath)
		// Then
		require.NoError(t, err)
		assert.Empty(t, records)
	})

	t.Run("should ignore title from records", func(t *testing.T) {
		// Given
		filePath := "examples/just_title.csv"
		csvReader := file.NewCsvImporterFromFileSytem(utils_test.ExampleFilesFS)
		// When
		records, err := csvReader.GetRecordsFrom(filePath)
		// Then
		require.NoError(t, err)
		assert.Empty(t, records)
	})

	t.Run("should return matrix of strings when file exists and is not empty", func(t *testing.T) {
		// Given
		filePath := "examples/cash_app_report.csv"
		csvReader := file.NewCsvImporterFromFileSytem(utils_test.ExampleFilesFS)
		// When
		records, err := csvReader.GetRecordsFrom(filePath)
		// Then
		require.NoError(t, err)
		assert.NotEmpty(t, records)
	})
}
