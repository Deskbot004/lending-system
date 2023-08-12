package sql

import (
	"context"
	"fmt"
	"lending-system/ent"
	"lending-system/ent/user"
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

func GetUserByName(ctx context.Context, client *ent.Client, name string) (*ent.User, error) {
	u, err := client.User.
		Query().
		Where(user.Name(name)).
		Only(ctx)
	if ent.IsNotFound(err) {
		return nil, nil
	} else if err != nil {
		return nil, fmt.Errorf("failed querying user: %w", err)
	}
	return u, nil
}

func DeleteUser(ctx context.Context, client *ent.Client, user *ent.User) error {
	err := client.User.
		DeleteOne(user).
		Exec(ctx)
	if err != nil {
		return fmt.Errorf("failed deleting user: %w", err)
	}
	return nil
}

func AddGameToUser(ctx context.Context, user *ent.User, game *ent.Game) error {
	_, err := user.Update().
		AddGames(game).
		Save(ctx)
	if err != nil {
		return fmt.Errorf("failed adding edge to user:  %w", err)
	}
	return nil
}

func GetUserGames(ctx context.Context, user *ent.User) ([]*ent.Game, error) {
	games, err := user.QueryGames().All(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed querying user games: %w", err)
	}
	return games, nil
}
