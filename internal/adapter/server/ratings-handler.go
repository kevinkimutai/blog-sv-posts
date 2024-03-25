package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kevinkimutai/metadata/internal/app/core/domain"
)

func (s *ServerAdapter) CreateMovieRating(c *fiber.Ctx) error {
	rating := domain.Rating{}

	//Bind To struct
	if err := c.BodyParser(&rating); err != nil {
		return c.Status(500).JSON(
			domain.ErrorResponse{
				StatusCode: 500,
				Message:    err.Error(),
			})
	}

	//Check Missing Inputs
	rating, err := domain.NewRatingsDomain(rating)
	if err != nil {
		return c.Status(400).JSON(
			domain.ErrorResponse{
				StatusCode: 400,
				Message:    err.Error(),
			})
	}

	//api
	rating, err = s.api.CreateNewRating(rating)
	if err != nil {
		return c.Status(500).JSON(
			domain.ErrorResponse{
				StatusCode: 500,
				Message:    err.Error(),
			})
	}

	//Map Response
	res := domain.DataResponse{
		StatusCode: 201,
		Message:    "success",
		Data:       rating,
	}

	return c.Status(fiber.StatusOK).JSON(res)
}
