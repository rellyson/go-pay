package main

import (
	"os"

	"github.com/rellyson/go-pay/infra/http"
	"github.com/rellyson/go-pay/infra/persistence"
)

func main() {
	// dbDsn := "root:foobar@tcp(0.0.0.0:3306)/go-pay_dev?charset=utf8mb4&parseTime=True&loc=Local"
	dbDsn := os.Getenv("DB_DSN")
	addr := os.Getenv("APP_ADDR")

	persistence.CreateDbConnection(dbDsn)
	s := http.CreateHttpServer()

	s.Listen(addr)
}
