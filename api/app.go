package main

import (
	"github.com/robertocamp/gd5gofcd/api/routes"
	"github.com/robertocamp/gd5gofcd/pkg/book"
	// "github.com/robertocamp/gd5gofcd/pkg/stack"
	"context"
	"fmt"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	db, cancel, err := databaseConnection()
	if err != nil {
		log.Fatal("Database Connection Error $s", err)
	}
	// services are defined in the pkg "book" or "stack"
	fmt.Println("Database connection success!")
	bookCollection := db.Collection("books")
	bookRepo := book.NewRepo(bookCollection)
	bookService := book.NewService(bookRepo)
	stackService := stack.NewService(stackAPI)

	app := fiber.New()
	app.Use(cors.New())
	app.Use(logger.New())
	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.Send([]byte("Welcome to the github.com/robertocamp/ mongo book shop!"))
		// return ctx.SendString("Hello, World 👋!")
	})
	api := app.Group("/api")
	routes.BookRouter(api, bookService)
	routes.StackRouter(api, stackService)

	defer cancel()
	log.Fatal(app.Listen(":3000"))
}

func databaseConnection() (*mongo.Database, context.CancelFunc, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(
		"mongodb://127.0.0.1:27017/test").SetServerSelectionTimeout(5*time.
		Second))
	if err != nil {
		cancel()
		return nil, nil, err
	}
	db := client.Database("books")
	return db, cancel, nil
}