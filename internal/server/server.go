package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type FiberServer struct {
	*fiber.App
}

func New() *FiberServer {
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST",
	}))

	// Serve static files (frontend)
	app.Static("/", "dist/")

	server := &FiberServer{
		App: app,
	}

	return server
}
