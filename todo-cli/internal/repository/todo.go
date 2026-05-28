package repository

import (
	"encoding/json"
	"fmt"
	"os"
	"todo-cli/internal/models"
)

var Todos []models.Todo

func LoadTodos() {
	file, err := os.Open("dados/todos.json")
	if err != nil {
		fmt.Println("error file:", err)
		return
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&Todos); err != nil {
		fmt.Println("error decoding JSON:", err)
	}
}

func SaveTodo() {
	file, err := os.Create("dados/todos.json")
	if err != nil {
		fmt.Println("error file:", err)
		return
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	if err := encoder.Encode(Todos); err != nil {
		fmt.Println("error encoding JSON:", err)
	}
}
