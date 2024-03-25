package server

import (
	"github.com/gofiber/fiber/v2"
)

func (s *ServerAdapter) MetadataRouter(api fiber.Router) {
	api.Post("/", s.CreateMovie)
	api.Get("/", s.GetAllMovies)
	api.Get("/:movieID", s.GetMovieMetadataByID)

}
