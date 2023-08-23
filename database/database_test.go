package database_test

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestRetrieveUserDetails(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectBegin()
	mock.ExpectExec("SELECT  username, password ").WillReturnResult(sqlmock.NewResult(1, 1))

}
