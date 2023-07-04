package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/izasoerya/RestAPI-Todo/controller"
)

func todoRoute (route fiber.Router) {
	route.Get("", controller.GetTodos)
}

