package main

import (
	"log"
	"os"

	"github.com/OnlyLight/totolist-golang/helper"
	"github.com/OnlyLight/totolist-golang/route"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	app := fiber.New()

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	helper.Init()

	app.Get("/api/todos", route.GetTodos)
	app.Post("/api/todos", route.CreateTodo)
	app.Patch("/api/todos/:id", route.UpdateTodo)
	app.Delete("/api/todos/:id", route.DeleteTodo)

	app.Listen(":" + os.Getenv("PORT"))
}
