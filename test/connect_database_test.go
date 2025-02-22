package test

import (
	"harmonee-pos/database"
	"testing"
)

func TestConnectDatabase(t *testing.T) {
	database.ConnectDatabase()
}
