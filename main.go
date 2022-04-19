package main

import (
	"log"
	"os"

	"github.com/rellyson/go-pay/infra/http"
)

func main() {
	addr, exists := os.LookupEnv("APP_ADDR")
	if !exists {
		addr = ":3000"
		log.Printf("missing APP_ADDR environment variable! Listens default to %v", addr)
	}

	s := http.CreateHttpServer()
	s.Listen(addr)
}
