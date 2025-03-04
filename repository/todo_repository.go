package repository

import (
	"context"
	"log"
	"todo_api/config"
	"todo_api/models"
)

func GetTodos() ([]models.Todo, error) {
	rows, err := config.DB.Query(context.Background(), "SELECT id, title, completed FROM todos")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Ensure slice is initialized
	todos := []models.Todo{}

	for rows.Next() {
		var todo models.Todo
		err := rows.Scan(&todo.ID, &todo.Title, &todo.Completed)
		if err != nil {
			log.Println("Error scanning row:", err)
			continue
		}
		todos = append(todos, todo)
	}

	return todos, nil
}

func GetTodoByID(id int) (models.Todo, error) {
	var todo models.Todo
	err := config.DB.QueryRow(context.Background(), "SELECT id, title, completed FROM todos WHERE id=$1", id).
		Scan(&todo.ID, &todo.Title, &todo.Completed)
	return todo, err
}

func CreateTodo(todo models.Todo) (int, error) {
	var id int
	err := config.DB.QueryRow(context.Background(), "INSERT INTO todos (title, completed) VALUES ($1, $2) RETURNING id",
		todo.Title, todo.Completed).Scan(&id)
	return id, err
}

func UpdateTodo(todo models.Todo) error {
	_, err := config.DB.Exec(context.Background(), "UPDATE todos SET title=$1, completed=$2 WHERE id=$3",
		todo.Title, todo.Completed, todo.ID)
	return err
}

func DeleteTodo(id int) error {
	_, err := config.DB.Exec(context.Background(), "DELETE FROM todos WHERE id=$1", id)
	return err
}
