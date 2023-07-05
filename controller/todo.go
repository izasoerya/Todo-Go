package controller

import (
	"fmt"
	"strconv"

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

func CreateTodos(c *fiber.Ctx) error {
	type Request struct {
		Title string `json:"title"`
	}

	var body Request
	err := c.BodyParser(&body)

	if err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "bad request create todo",
		})
	}

	todo := &todo{
		Id:        len(Todos) + 1,
		Title:     body.Title,
		Completed: false,
	}

	Todos = append(Todos, todo)

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": true,
		"message": "todo created!",
	})
}

func SearchTodos(c *fiber.Ctx) error {
	type Request struct {
		Id string `json:"id"`
	}

	var body Request
	err := c.BodyParser(&body)

	if err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "bad request search todo",
		})
	}
	SearchID, _ := strconv.Atoi((body.Id))
	for _, todo := range Todos {
		if todo.Id == SearchID {
			return c.Status(fiber.StatusOK).JSON(fiber.Map{
				"success": true,
				"data": fiber.Map{
					"todo": todo,
				},
			})
		}
	}
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"success": false,
		"message": "cant find id",
		"data": fiber.Map{
			"todo": Todos,
		},
	})
}


