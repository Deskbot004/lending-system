package web

import (
	"context"
	"fmt"
	"lending-system/ent"
	"lending-system/sql"

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

func checkElseredirect(c *gin.Context, r *gin.Engine) bool {
	session := sessions.Default(c)
	if session.Get("login") != "true" {
		c.Request.URL.Path = "/"
		r.HandleContext(c)
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
		if checkElseredirect(c, router) {
			return
		}
	}
}