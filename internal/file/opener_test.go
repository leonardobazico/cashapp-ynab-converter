package file_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"cash2ynab/internal/file"
	utils_test "cash2ynab/tests/utils"
)

func TestOpener(t *testing.T) {
	t.Parallel()

	t.Run("should open file in a given file system", func(t *testing.T) {
		t.Parallel()

		// Given
		filePath := "examples/cash_app_report.csv"
		opener := file.NewFileSytemOpener(utils_test.ExampleFilesFS)
		// When
		fileOpened, err := opener.Open(filePath)
		// Then
		assert.NoError(t, err)
		err = fileOpened.Close()
		assert.NoError(t, err)
	})

	t.Run("should open file using os.OpenFile when no file system is given", func(t *testing.T) {
		t.Parallel()

		// Given
		filePath := "opener_test.go"
		opener := file.NewOsOpener()
		// When
		fileOpened, err := opener.Open(filePath)
		// Then
		assert.NoError(t, err)
		err = fileOpened.Close()
		assert.NoError(t, err)
	})

	t.Run("should open file given an absolute path", func(t *testing.T) {
		t.Parallel()

		// Given
		testFolder, _ := os.Getwd()
		filePath := testFolder + "/opener_test.go"
		opener := file.NewOsOpener()
		// When
		fileOpened, err := opener.Open(filePath)
		// Then
		assert.NoError(t, err)
		err = fileOpened.Close()
		assert.NoError(t, err)
	})

	t.Run("should fail when file does not exist", func(t *testing.T) {
		t.Parallel()

		// Given
		filePath := "does-not-exist.csv"
		opener := file.NewOsOpener()
		// When
		fileOpened, err := opener.Open(filePath)
		// Then
		assert.Error(t, err)
		assert.Nil(t, fileOpened)
	})
}
