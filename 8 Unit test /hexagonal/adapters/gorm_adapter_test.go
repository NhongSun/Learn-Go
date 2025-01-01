package adapters

import (
	"nhongsun/core"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestGormOrderRepository_Save(t *testing.T) {
	// Mock database
	sqlDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("An error '%s' was not expected when opening a stub database connection", err)
	}
	defer sqlDB.Close()

	// Expectation for the sqlite version check
	mock.ExpectQuery("select sqlite_version()").WillReturnRows(sqlmock.NewRows([]string{"version"}).AddRow("3.31.1"))

	dialector := sqlite.Dialector{Conn: sqlDB}
	gormDB, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		t.Fatalf("Failed to open gorm database: %v", err)
	}

	repo := NewGormOrderRepository(gormDB)

	// Success case
	t.Run("success", func(t *testing.T) {
		// Setup expectations
		mock.ExpectBegin()
		mock.ExpectExec("INSERT INTO").WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()

		err := repo.Save(core.Order{Total: 100})
		assert.NoError(t, err)

		// Ensure all expectations were met
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	// Failure case
	t.Run("failure", func(t *testing.T) {
		// Setup expectations
		mock.ExpectBegin()
		mock.ExpectExec("INSERT INTO").WillReturnError(gorm.ErrInvalidData)
		mock.ExpectRollback()

		err := repo.Save(core.Order{Total: 100})
		assert.Error(t, err)

		// Ensure all expectations were met
		assert.NoError(t, mock.ExpectationsWereMet())
	})
}
