package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/relentlessworks/fakerkit/internal/config"
	"github.com/relentlessworks/fakerkit/internal/server"
)

func main() {
	cfg := config.FromFlags()
	srv := server.New()

	addr := cfg.Addr()
	fmt.Printf("fakerkit listening on %s\n", addr)
	fmt.Printf("GET /help for usage\n")

	if err := http.ListenAndServe(addr, srv.Routes()); err != nil {
		log.Fatal(err)
	}
}
