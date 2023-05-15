package test01

import (
	"context"
	"log"

	"ginent/ent"
)

func GetTest01(ctx *context.Context, client *ent.Client) ([]*ent.Test01, error) {
	tasks, err := client.Test01.Query().All(*ctx)
	if err != nil {
		log.Fatalf("error : %v", err)
		return nil, err
	}

	return tasks, err
}

func CreateTest01(ctx *context.Context, client *ent.Client, input *ent.CreateTest01Input) (*ent.Test01, error) {
	task, err := client.Test01.Create().SetInput(*input).Save(*ctx)
	if err != nil {
		log.Fatalf("error : %v", err)
		return nil, err
	}

	return task, err
}
