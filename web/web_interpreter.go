package web

import (
	"context"
	"fmt"
	"io"
	"lending-system/ent"
	"lending-system/sql"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"sort"
	"strings"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type Form struct {
	File *multipart.FileHeader `form:"file" binding:"required"`
}

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

func changePicture(c *gin.Context, oldpic string) (string, error) {
	file, header, err := c.Request.FormFile("file")
	if file == nil {
		return oldpic, nil
	}
	if err != nil {
		return "default.png", err
	}
	filename := header.Filename
	if !strings.HasSuffix(filename, ".png") && !strings.HasSuffix(filename, ".jpg") {
		log.Println(strings.HasSuffix(filename, ".png"))
		log.Println(filename)
		return "default.png", fmt.Errorf("file is not an image")
	}

	defer file.Close()
	out, err := os.Create("./assets/images/user/" + filename)
	if err != nil {
		return "default.png", err
	}

	defer out.Close()
	_, err = io.Copy(out, file)
	if err != nil {
		return "default.png", err
	}

	if oldpic != "default.png" {
		err = os.Remove("./assets/images/user/" + oldpic)
		if err != nil {
			log.Println(err)
		}
	}

	return filename, err
}

func sortArrays(users []*ent.User, games []*ent.Game) {
	sort.Slice(users, func(i, j int) bool {
		return users[i].Name < users[j].Name
	})
	sort.Slice(games, func(i, j int) bool {
		return games[i].Name < games[j].Name
	})
}

func removeUnlended(games []*ent.Game) []*ent.Game{
	gameNum := len(games)
	for i := gameNum - 1; i >= 0; i-- {
		if games[i].Cu == games[i].Ou {
			games = remove(games, i)
		}
	}
	return games
}

func remove(s []*ent.Game, i int) []*ent.Game {
    s[i] = s[len(s)-1]
    return s[:len(s)-1]
}