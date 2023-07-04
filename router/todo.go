package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/izasoerya/RestAPI-Todo/controller"
)

func TodoRoute(route fiber.Router) {
	route.Get("", controller.GetTodos)
}
