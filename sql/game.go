package sql

import (
	"context"
	"fmt"
	"lending-system/ent"
	"lending-system/ent/game"
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

func UpdateGame(ctx context.Context, game *ent.Game, gamenew ent.Game) error {
	_, err := game.Update().
		SetName(gamenew.Name).
		SetType(gamenew.Type).
		SetCu(gamenew.Cu).
		SetNotes(gamenew.Notes).
		Save(ctx)
	if err != nil {
		return fmt.Errorf("failed updating user: %w", err)
	}
	return nil
}

func GetGameByID(ctx context.Context, client *ent.Client, ID int) (*ent.Game, error) {
	u, err := client.Game.
		Query().
		Where(game.ID(ID)).
		Only(ctx)
	if ent.IsNotFound(err) {
		return nil, nil
	} else if err != nil {
		return nil, fmt.Errorf("failed querying game: %w", err)
	}
	return u, nil
}

func DeleteGame(ctx context.Context, client *ent.Client, game *ent.Game) error {
	err := client.Game.
		DeleteOne(game).
		Exec(ctx)
	if err != nil {
		return fmt.Errorf("failed deleting game: %w", err)
	}
	return nil
}
