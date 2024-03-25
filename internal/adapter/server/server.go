package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/kevinkimutai/metadata/internal/ports"
)

type ServerAdapter struct {
	port string
	api  ports.APIPort
}

func New(port string, api ports.APIPort) *ServerAdapter {
	return &ServerAdapter{port: port, api: api}
}

func (s *ServerAdapter) Run() {
	//Initialize Fiber
	app := fiber.New()

	//Logger Middleware
	app.Use(logger.New())

	// Define routes
	app.Route("/api/v1/rating", s.RatingsRouter)
	app.Route("/api/v1/metadata", s.MetadataRouter)

	app.Listen(":" + s.port)
}
