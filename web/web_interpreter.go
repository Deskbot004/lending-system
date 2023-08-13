package web

import (
	"context"
	"fmt"
	"lending-system/ent"
	"lending-system/sql"
	"net/http"

	"github.com/gin-contrib/sessions"
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

func checkElseredirect(c *gin.Context) bool {
	session := sessions.Default(c)
	if session.Get("login") != "true" {
		c.Redirect(http.StatusTemporaryRedirect, "/")
		return true
	}
	return false
}

func aheader() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.URL.Path == "/" {
			return
		}
		errMess = ""
		if checkElseredirect(c) {
			return
		}
	}
}

func cleanDeleteUser(ctx context.Context, client *ent.Client, user *ent.User) error {
	games, err := sql.GetUserGames(ctx, user)
	if err != nil {
		return err
	}
	for i := range games {
		err = sql.DeleteGame(ctx, client, games[i])
		if err != nil {
			return err
		}
	}
	err = sql.DeleteUser(ctx, client, user)
	if err != nil {
		return err
	}
	return nil
}
