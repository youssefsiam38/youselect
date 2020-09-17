package db

import (
	"testing"
)

func TestSetup(t *testing.T) {
	Setup()
	db, err := Connect()
	defer db.Close()

	if err != nil {
		t.Error(err)
	}
}
