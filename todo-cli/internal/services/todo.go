package services

import (
	"fmt"
	"time"
	"todo-cli/internal/models"
	"todo-cli/internal/repository"
)

func CreateTodo(msg string) {
	var newTodo models.Todo

	newTodo.ID = len(repository.Todos) + 1
	newTodo.Complete = false
	newTodo.CreatedAt = time.Now()
	newTodo.Description = msg

	repository.Todos = append(repository.Todos, newTodo)
	repository.SaveTodo()
}

func ListTodos() {
	for _, t := range repository.Todos {
		fmt.Printf("ID: %d | Descrição: %s | Completo: %v | Criado em [%v]\n", t.ID, t.Description, t.Complete, t.CreatedAt.Format("02/01/2006-15:04:05"))
	}
}

func MarkCompleteTodo(id int) {
	for i, t := range repository.Todos {
		if t.ID == id {
			repository.Todos[i].Complete = true
			fmt.Println("ToDo atualizado com sucesso.")
			repository.SaveTodo()
			return
		}
	}
	fmt.Println("ToDo não encontrado")
}

func DeleteTodo(id int) {
	for i, t := range repository.Todos {
		if t.ID == id {
			repository.Todos = append(repository.Todos[:i], repository.Todos[i+1:]...)
			fmt.Println("toDo deletado com sucesso.")
			repository.SaveTodo()
			return
		}
	}
	fmt.Println("ToDo não encontrado.")
}
