package web

import (
	"context"
	"fmt"
	"lending-system/ent"
	"lending-system/sql"

	"github.com/gin-gonic/gin"
)

func getGamesAndUsers(ctx context.Context, client *ent.Client) ([]*ent.User, []*ent.Game, error) {
	users, err := sql.GetAllUsers(ctx, client)
	if err != nil {
		return nil, nil, fmt.Errorf("getting all users failed: %v", err)
	}
	games, err := sql.GetAllGames(ctx, client)
	if err != nil {
		return nil, nil, fmt.Errorf("getting all games failed: %v", err)
	}
	return users, games, err
}

func checkElseredirect(c *gin.Context) {

}
