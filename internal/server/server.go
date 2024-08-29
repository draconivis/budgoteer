package server

import (
	"github.com/gofiber/fiber/v2"

	"budgoteer/internal/database"
)

type FiberServer struct {
	*fiber.App

	db database.Service
}

func New() *FiberServer {
	server := &FiberServer{
		App: fiber.New(fiber.Config{
			ServerHeader: "new",
			AppName:      "new",
		}),

		db: database.New(),
	}

	return server
}
