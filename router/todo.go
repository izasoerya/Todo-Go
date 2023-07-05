package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/izasoerya/RestAPI-Todo/controller"
)

func TodoRoute(route fiber.Router) {
	route.Get("/:id", controller.SearchTodosGet)
	route.Get("", controller.GetAllTodos)
	route.Post("", controller.CreateTodos)
	route.Delete("", controller.DeleteTodosRequest)

}
