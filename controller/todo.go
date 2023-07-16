package controller

import (
	"fmt"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/izasoerya/RestAPI-Todo/config"
	"github.com/izasoerya/RestAPI-Todo/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetAllTodos(c *fiber.Ctx) error {
	todoCollection := config.MI.DB.Collection(os.Getenv("TODO_COLLECTION"))
	query := bson.D{{}}
	cursor, err := todoCollection.Find(c.Context(), query)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "someting wrong",
			"error":   err.Error(),
		})
	}

	var Todos []models.Todo = make([]models.Todo, 0)

	err = cursor.All(c.Context(), &Todos)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "someting wrong",
			"error":   err.Error(),
		})
	}
	
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": fiber.Map{
			"todo": Todos,
		},
	})
}

func SearchTodosGet(c *fiber.Ctx) error {
	todoCollection := config.MI.DB.Collection(os.Getenv("TODO_COLLECTION"))
	paramId := c.Params("id")
	SearchID, err := primitive.ObjectIDFromHex(paramId)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "cannot convert to int",
		})
	}

	todo := &models.Todo{}
	query := bson.D{{Key: "_id", Value: SearchID}}
	err = todoCollection.FindOne(c.Context(), query).Decode(todo)

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": "Todo Not found",
			"error":   err,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data": fiber.Map{
			"todo": todo,
		},
	})
}

func CreateTodos(c *fiber.Ctx) error {
	todoCollection := config.MI.DB.Collection(os.Getenv("TODO_COLLECTION"))
	
	type Request struct {
		Title string `json:"title"`
	}

	var body Request 
	err:= c.BodyParser(&body)

	if err != nil {
		fmt.Println(err)
	}

	data := new(models.Todo)

	completed := false
	data.ID = nil		
	data.Title = &body.Title
	data.Completed = &completed
	data.CreatedAt = time.Now()
	data.UpdatedAt = time.Now()

	fmt.Println(data.Title)
	todoCollection.InsertOne(c.Context(), data)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": body.Title,
		"action": c.RedirectBack("/app"),
	})
}

func DeleteTodos(c *fiber.Ctx) error {
	todoCollection := config.MI.DB.Collection(os.Getenv("TODO_COLLECTION"))
	paramID := c.Params("id")
	fmt.Println(paramID)

	SearchID, err := primitive.ObjectIDFromHex(paramID)
	if err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "bad request search todo",
		})
	}

	query := bson.D{{Key: "_id", Value: SearchID}}
	err = todoCollection.FindOneAndDelete(c.Context(), query).Err()

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"success": false,
				"message": "Todo Not found",
				"error":   err,
			})
		}

		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot delete todo",
			"error":   err,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
        "message": "todo deleted",
	})
}

func EditTodos(c *fiber.Ctx) error {
	todoCollection := config.MI.DB.Collection(os.Getenv("TODO_COLLECTION"))
	c.JSON(fiber.Map{
		"created": false,
	})
	paramID := c.Params("id")

	id, err := primitive.ObjectIDFromHex(paramID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse id",
			"error":   err,
		})
	}

	data := new(models.Todo)
	err = c.BodyParser(&data)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse JSON",
			"error":   err,
		})
	}

	query := bson.D{{Key: "_id", Value: id}}
	var dataToUpdate bson.D

	if data.Title != nil {
		dataToUpdate = append(dataToUpdate, bson.E{Key: "title", Value: data.Title})
	}

	if data.Completed != nil {
		dataToUpdate = append(dataToUpdate, bson.E{Key: "completed", Value: data.Completed})
	}

	dataToUpdate = append(dataToUpdate, bson.E{Key: "updatedAt", Value: time.Now()})

	update := bson.D{
		{Key: "$set", Value: dataToUpdate},
	}

	err = todoCollection.FindOneAndUpdate(c.Context(), query, update).Err()

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"success": false,
				"message": "Todo Not found",
				"error":   err,
			})
		}

		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot update todo",
			"error":   err,
		})
	}

	todo := &models.Todo{}

	todoCollection.FindOne(c.Context(), query).Decode(todo)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data": fiber.Map{
			"todo": todo,
		},
	})
}
