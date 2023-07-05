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
	{
		Id:        3,
		Title:     "C# tutorial",
		Completed: true,
	},
}

func GetAllTodos(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": fiber.Map{
			"todo": Todos,
		},
	})
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

func SearchTodosRequest(c *fiber.Ctx) error {
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

func SearchTodosGet(c *fiber.Ctx) error {
	paramId := c.Params("id")

	SearchID, err := strconv.Atoi(paramId)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "cannot convert to int",
		})
	}

	for _, todo := range Todos {
		if todo.Id == SearchID {
			return c.Status(fiber.StatusOK).JSON(fiber.Map{
				"success": true,
				"message": "found the id",
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

func DeleteTodos(c *fiber.Ctx) error {
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

	SearchID, _ := strconv.Atoi(body.Id)
	for i, todo := range Todos {
		if SearchID == todo.Id {
			Todos = append(Todos[:i], Todos[i+1:]...)
			return c.Status(fiber.StatusOK).JSON(fiber.Map{
				"success": true,
				"message": "Successfully deleted",
			})
		}
	}

	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"success": false,
		"message": "failed delete",
	})
}

func EditTodos(c *fiber.Ctx) error {
	paramID := c.Params("id")
	type Request struct {
		Id        int    `json:"id"`
		Title     string `json:"title"`
		Completed bool   `json:"completed"`
	}

	var body Request
	err := c.BodyParser(&body)
	if err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "bad request edit todo",
		})
	}

	EditID, _ := strconv.Atoi(paramID)

	for i, todo := range Todos {
		if EditID == todo.Id {
			Todos[i].Id			= body.Id
			Todos[i].Title 		= body.Title
			Todos[i].Completed 	= body.Completed
			return c.Status(fiber.StatusOK).JSON(fiber.Map{
				"data": fiber.Map{
					"todo": Todos,
				},
			})
		}
	}

	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"success": false,
		"message": "failed edit",
	})
}
