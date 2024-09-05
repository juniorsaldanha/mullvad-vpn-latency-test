package main

import (
	"github.com/juniorsaldanha/mullvad-vpn-latency-test/internal/server"
)

func main() {
	app := server.New()
	app.RegisterFiberRoutes()

	err := app.Listen(":8080")
	if err != nil {
		panic(err)
	}
}
