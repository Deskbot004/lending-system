package web

import (
	"context"
	"lending-system/ent"
	"lending-system/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
StartServer starts the web server

client: client that requested the database
*/
func StartServer(client *ent.Client) {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.Static("/assets", "./assets")
	var err error
	var errMess string

	// Landing Pate login later
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "_login.html", 1)
	})

	// Startpage _dashboard.html
	router.GET("/index", func(c *gin.Context) {
		users, games, err := getGamesAndUsers(context.Background(), client)
		if err != nil {
			errMess = "Error happened"
			log.Println(err)
		}
		c.HTML(http.StatusOK, "_dashboard.html", gin.H{
			"Usernum": len(users),
			"Gamenum": len(games),
			"Error":   errMess,
		})
	})

	// Game overviews
	router.GET("/game_overview", func(c *gin.Context) {
		users, games, err := getGamesAndUsers(context.Background(), client)
		if err != nil {
			errMess = "Error happened"
			log.Println(err)
		}
		c.HTML(http.StatusOK, "_game_overview.html", gin.H{
			"Users":   users,
			"Games":   games,
			"Gamenum": len(games),
			"Error":   errMess,
		})
	})

	router.GET("/game_overview/:name", func(c *gin.Context) {
		name := c.Param("name")
		user, err := sql.GetUserByName(context.Background(), client, name)
		if err != nil {
			errMess = "Error happened"
			log.Println(err)
		}
		users, _, err := getGamesAndUsers(context.Background(), client)
		if err != nil {
			errMess = "Error happened"
			log.Println(err)
		}
		games, err := sql.GetUserGames(context.Background(), user)
		if err != nil {
			errMess = "Error happened"
			log.Println(err)
		}
		// change games to relevant games with Gameinfo
		c.HTML(http.StatusOK, "_game_inner.html", gin.H{
			"Users":    users,
			"Username": user.Name,
			"Games":    games,
			"Gamenum":  len(games),
			"Error":    errMess,
		})
	})

	router.GET("/game_overview/:name/addGame", func(c *gin.Context) {
		name := c.Param("name")
		c.HTML(http.StatusOK, "_addGame.html", gin.H{
			"Username": name,
			"Error":    errMess,
		})
	})

	router.POST("/game_overview/:name/addGame", func(c *gin.Context) {
		username := c.Param("name")
		gametype := c.PostForm("type")
		gamename := c.PostForm("name")
		game := ent.Game{
			Name: gamename,
			Type: gametype,
			Ou:   username,
		}

		user, err := sql.GetUserByName(context.Background(), client, username)
		if err != nil {
			errMess = "Error happened"
			log.Println(err)
		}
		_, err = sql.AddGame(context.Background(), client, game, user)
		if err != nil {
			errMess = "Error happened"
			log.Println(err)
		}
		users, _, err := getGamesAndUsers(context.Background(), client)
		if err != nil {
			errMess = "Error happened"
			log.Println(err)
		}
		games, err := sql.GetUserGames(context.Background(), user)
		if err != nil {
			errMess = "Error happened"
			log.Println(err)
		}
		// change games to relevant games with Gameinfo
		c.HTML(http.StatusOK, "_game_inner.html", gin.H{
			"Users":    users,
			"Username": user.Name,
			"Games":    games,
			"Gamenum":  len(games),
			"Error":    errMess,
		})
	})

	router.GET("/deleteUser/:name", func(c *gin.Context) {
		//TODO delete all packages too
		name := c.Param("name")
		user, err := sql.GetUserByName(context.Background(), client, name)
		if err != nil {
			errMess = "Error happened"
			log.Println(err)
		}
		err = sql.DeleteUser(context.Background(), client, user)
		if err != nil {
			errMess = "Error happened"
			log.Println(err)
		}
		users, games, err := getGamesAndUsers(context.Background(), client)
		if err != nil {
			errMess = "Error happened"
			log.Println(err)
		}
		c.HTML(http.StatusOK, "_game_overview.html", gin.H{
			"Users":   users,
			"Games":   games,
			"Gamenum": len(games),
			"Error":   errMess,
		})
	})

	// Adding Users
	router.GET("/addUser", func(c *gin.Context) {
		c.HTML(http.StatusOK, "_addUser.html", gin.H{
			"Error": errMess,
		})
	})

	router.POST("/addUser", func(c *gin.Context) {
		username := c.PostForm("name")
		user := ent.User{
			Name: username,
		}
		sql.AddUser(context.Background(), client, user)
		// if fail
		/*
			c.HTML(http.StatusOK, "_add.html", gin.H{
				"Error":   errMess,
			})
		*/
		// if success
		users, games, err := getGamesAndUsers(context.Background(), client)
		if err != nil {
			errMess = "Error happened"
			log.Println(err)
		}
		c.HTML(http.StatusOK, "_dashboard.html", gin.H{
			"Usernum": len(users),
			"Gamenum": len(games),
			"Error":   errMess,
		})
	})

	// No route
	router.NoRoute(func(c *gin.Context) {
		c.HTML(http.StatusNotFound, "_404.html", 1)
	})

	err = router.Run() // listen and serve on 0.0.0.0:8080
	if err != nil {
		log.Fatalf("failed starting router: %v", err)
	}
}
