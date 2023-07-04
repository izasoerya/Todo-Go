package controller

import (
	"github.com/gofiber/fiber/v2"
)

type todo struct {
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

var Todos = []*todo{
	{
		Id:        1,
		Title:     "Go tutorial",
		Completed: false,
	},
	{
		Id:        2,
		Title:     "JS tutorial",
		Completed: true,
	},
}

func GetTodos(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data": fiber.Map{
			"todos": Todos,
		},
	})
}
