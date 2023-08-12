package sql

import (
	"context"
	"fmt"
	"lending-system/ent"
	"log"
)

func AddUser(ctx context.Context, client *ent.Client, user ent.User) (*ent.User, error) {
	g, err := client.User.
		Create().
		SetName(user.Name).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating user: %w", err)
	}
	log.Println("user was created: ", g)
	return g, nil
}

func GetAllUsers(ctx context.Context, client *ent.Client) ([]*ent.User, error) {
	users, err := client.User.Query().All(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed finding users: %w", err)
	}
	return users, nil
}
