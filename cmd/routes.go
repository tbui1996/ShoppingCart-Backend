package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tbui1996/backend-practice/handlers"
)

func setupRoutes(app *fiber.App) {
	api := app.Group("/api")

	//books
	api.Get("/books/:title/:author/:genre", handlers.ListBooks)
	api.Post("/book", handlers.CreateBook)
	api.Delete("/books", handlers.DeleteAllBooks)
	//cart
	api.Get("/cart", handlers.ListCart)
	api.Put("/cart/:id", handlers.AddToCart)
	api.Delete("/cart/:id", handlers.RemoveFromCart)
	api.Post("/NewCart", handlers.CreateEmptyCart) // New endpoint to create an empty cart
	api.Get("/temporary/reset", handlers.TemporaryResetHandler)
}
