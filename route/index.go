package route

import (
	"context"
	"log"

	"github.com/OnlyLight/totolist-golang/helper"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type Todo struct {
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Completed bool               `json:"completed"`
	Body      string             `json:"body"`
}

func GetTodos(c *fiber.Ctx) error {
	var todos []Todo

	cursor, err := helper.Collection.Find(context.Background(), bson.M{})

	if err != nil {
		return err
	}

	// defer will run after the func GetTodos completed
	defer func() {
		if err = cursor.Close(context.Background()); err != nil {
			panic(err)
		}
	}()

	for cursor.Next(context.Background()) {
		todo := &Todo{}

		if err := cursor.Decode(todo); err != nil {
			return err
		}
		todos = append(todos, *todo)
	}

	return c.JSON(todos)
}

func CreateTodo(c *fiber.Ctx) error {
	// todo := new(Todo)
	// var todo *Todo = &Todo{}
	todo := &Todo{
		ID: primitive.NewObjectID(),
	}

	if err := c.BodyParser(todo); err != nil {
		return err
	}

	if todo.Body == "" {
		return c.Status(400).JSON(fiber.Map{"error": "Todo body is required"})
	}

	_, err := helper.Collection.InsertOne(context.Background(), todo)
	if err != nil {
		return err
	}

	return c.Status(201).JSON(fiber.Map{"data": *todo})
}

func GetTodoById(objectID primitive.ObjectID) Todo {
	todo := &Todo{}
	err := helper.Collection.FindOne(context.Background(), bson.M{"_id": objectID}).Decode(&todo)
	if err != nil {
		log.Fatal(err)
	}

	return *todo
}

func UpdateTodo(c *fiber.Ctx) error {
	id := c.Params("id")

	objectID, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid Todo ID"})
	}

	filter := bson.M{"_id": objectID}
	update := bson.M{"$set": bson.M{"completed": true}}

	_, err = helper.Collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return err
	}

	todo := GetTodoById(objectID)

	return c.Status(200).JSON(fiber.Map{"data": todo})
}

func DeleteTodo(c *fiber.Ctx) error {
	id := c.Params("id")
	objectID, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid todo ID"})
	}

	filter := bson.M{"_id": objectID}
	_, err = helper.Collection.DeleteOne(context.Background(), filter)

	if err != nil {
		return err
	}

	return c.Status(200).JSON(fiber.Map{"data": "Deleted " + id})
}
