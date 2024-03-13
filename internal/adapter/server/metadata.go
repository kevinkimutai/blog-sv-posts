package server

import (
	"github.com/gofiber/fiber/v2"
)

func (s *ServerAdapter) MetadataRouter(api fiber.Router) {
	api.Post("/", s.CreateMovie)
	api.Get("/:movieID", s.GetMovieMetadataByID)

}
