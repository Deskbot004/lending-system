package sql

import (
	"context"
	"fmt"
	"lending-system/ent"
	"log"
)

func AddGame(ctx context.Context, client *ent.Client, game ent.Game, user *ent.User) (*ent.Game, error) {
	g, err := client.Game.
		Create().
		SetName(game.Name).
		SetType(game.Type).
		SetOu(game.Ou).
		SetCu(game.Ou).
		SetNotes("").
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating game: %w", err)
	}
	err = AddGameToUser(ctx, user, g)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}
	log.Println("game was created: ", g)
	return g, nil
}

func GetAllGames(ctx context.Context, client *ent.Client) ([]*ent.Game, error) {
	games, err := client.Game.Query().All(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed finding games: %w", err)
	}
	return games, nil
}