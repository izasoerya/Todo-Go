package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/izasoerya/RestAPI-Todo/controller"
)

func TodoRoute(route fiber.Router) {
	route.Get("/:id", controller.SearchTodosGet)
	route.Get("", controller.GetAllTodos)
	route.Delete("/:id", controller.DeleteTodos)
	route.Put("/:id", controller.EditTodos)

	route.Post("/:CreateTodoPage", controller.CreateTodoPage)
}
