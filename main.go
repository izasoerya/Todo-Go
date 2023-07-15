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

func setupRoutes(app *fiber.App) {									//* Create home page
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
	engine := html.New("./static", ".html")		
	app := fiber.New(fiber.Config{
		Views: engine,
	})
	app.Static("/static", "./static")

	app.Use(logger.New())											//* Logger
	err := godotenv.Load()											//* Init .env
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	
	config.ConnectDB()

	setupRoutes(app)

	app.Get("/app", func(c *fiber.Ctx) error {						//* set index.html in /app
		return c.Render("index", fiber.Map{})
	})
	app.Get("/app/edit/:id", func(c *fiber.Ctx) error {				//* set editTodo.html in /app/edit
		return c.Render("editTodo", fiber.Map{})
	})

	err = app.Listen(":3000")
	if err != nil {
		panic(err)
	}

}
