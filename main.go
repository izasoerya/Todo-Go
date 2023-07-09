package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/template/html/v2"
	"github.com/izasoerya/RestAPI-Todo/config"
	"github.com/izasoerya/RestAPI-Todo/router"

	"log"

	"github.com/joho/godotenv"
)

func setupRoutes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"success": true,
			"message": "get to end point",
		})
	})

	api := app.Group("/api")
	api.Get("", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"success": true,
			"message": "api endpoint",
		})
	})

	router.TodoRoute(api.Group("/todos"))
}

func main() {
	engine := html.New("./views", ".html")		
	app := fiber.New(fiber.Config{
		Views: engine,
	})
	app.Use(logger.New())

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	config.ConnectDB()

	setupRoutes(app)

	app.Get("/app", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{})
	})

	err = app.Listen(":3000")
	if err != nil {
		panic(err)
	}

}
