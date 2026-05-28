package main

import (
	"flag"
	"todo-cli/internal/repository"
	"todo-cli/internal/services"
)

func main() {
	repository.LoadTodos()

	todoCreate := flag.String("c", "", "criar um novo todo")
	todoList := flag.Bool("l", false, "listar todos os todos")
	todoComplete := flag.Int("u", 0, "marcar todo como completo")
	todoDelete := flag.Int("d", 0, "deletar todo")
	flag.Parse()

	if *todoCreate != "" {
		services.CreateTodo(*todoCreate)
	}
	if *todoList == true {
		services.ListTodos()
	}
	if *todoComplete != 0 {
		services.MarkCompleteTodo(*todoComplete)
	}
	if *todoDelete != 0 {
		services.DeleteTodo(*todoDelete)
	}
}
