package main

import (
	"log"
	"net/http"
	"os"
	"todo_api/config"
	"todo_api/handlers"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	_ = godotenv.Load()

	// Connect to DB
	config.ConnectDB()
	defer config.CloseDB()

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	// Routes
	r.Get("/todos", handlers.GetTodosHandler)
	r.Get("/todos/{id}", handlers.GetTodoByIDHandler)
	r.Post("/todos", handlers.CreateTodoHandler)
	r.Put("/todos", handlers.UpdateTodoHandler)
	r.Delete("/todos/{id}", handlers.DeleteTodoHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Println("Server running on port", port)
	http.ListenAndServe(":"+port, r)
}
