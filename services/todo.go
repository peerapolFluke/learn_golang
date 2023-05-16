package services

import (
	"context"
	"fmt"
	"log"

	"ginent/ent"
)

func GetTodo(ctx *context.Context, client *ent.Client) ([]*ent.Todo, error) {
	tasks, err := client.Todo.Query().WithUser().All(*ctx)
	fmt.Println(tasks[0].Edges.User)

	if err != nil {
		log.Fatalf("error : %v", err)
		return nil, err
	}

	return tasks, err
}

func CreateTodo(ctx *context.Context, client *ent.Client, input *ent.CreateTodoInput) (*ent.Todo, error) {
	task, err := client.Todo.Create().SetInput(*input).Save(*ctx)
	print("Create todo ==================")
	if err != nil {
		print("Create todo In If ==================")
		log.Fatalf("error : %v", err)
		return nil, err
	}

	return task, err
}
