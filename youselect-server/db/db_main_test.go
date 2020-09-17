package db

import (
	"github.com/joho/godotenv"
	"github.com/youssefsiam38/youselect/framework"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	godotenv.Load("../.test.env")
	err := FlushDatabase(os.Getenv("DB_NAME"))
	if err != nil {
		framework.Log(err)
	}
	Setup()

	exitVal := m.Run()

	// err = FlushDatabase(os.Getenv("DB_NAME"))
	// if err != nil {
	// 	framework.Log(err)
	// }
	os.Exit(exitVal)
}
