package web

import (
	"context"
	"lending-system/ent"
	"lending-system/sql"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)
var router *gin.Engine
var errMess string

/*
StartServer starts the web server

client: client that requested the database
*/
func StartServer(client *ent.Client) {
	router = gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.Static("/assets", "./assets")
	router.Static("/db", "./db")

	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("mysession", store))

	var err error
	password := "SeanJonas"
	errMess = ""

	router.Use(aheader())

	// Landing Pate login later
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "_login.html", gin.H{
			"Error": errMess,
		})
	})

	router.POST("/", func(c *gin.Context) {
		givenPassword := c.PostForm("pw")
		if givenPassword != password {
			errMess = "Falsches Passwort"
			c.HTML(http.StatusOK, "_login.html", gin.H{
				"Error": errMess,
			})
		} else {
			users, games, err := getGamesAndUsers(context.Background(), client)
			if err != nil {
				errMess = "Error happened"
				log.Println(err)
			}
			session := sessions.Default(c)
			session.Set("login", "true")
			session.Save()
			c.HTML(http.StatusOK, "_dashboard.html", gin.H{
				"Usernum": len(users),
				"Gamenum": len(games),
				"Error":   errMess,
			})
		}
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
		errMess = ""
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
		c.HTML(http.StatusOK, "_game_inner.html", gin.H{
			"Users":    users,
			"Username": user.Name,
			"Games":    games,
			"Gamenum":  len(games),
			"Error":    errMess,
		})
	})

	router.GET("/game_overview/:name/addGame", func(c *gin.Context) {
		errMess = ""
		name := c.Param("name")
		c.HTML(http.StatusOK, "_addGame.html", gin.H{
			"Username": name,
			"Error":    errMess,
		})
	})

	router.POST("/game_overview/:name/addGame", func(c *gin.Context) {
		errMess = ""
		username := c.Param("name")
		gametype := c.PostForm("gametype")
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
		c.HTML(http.StatusOK, "_game_inner.html", gin.H{
			"Users":    users,
			"Username": user.Name,
			"Games":    games,
			"Gamenum":  len(games),
			"Error":    errMess,
		})
	})

	// config views
	router.GET("/game_overview/:name/edit", func(c *gin.Context) {
		errMess = ""
		name := c.Param("name")
		user, err := sql.GetUserByName(context.Background(), client, name)
		if err != nil {
			errMess = "Error happened"
			log.Println(err)
		}
		c.HTML(http.StatusOK, "_configUser.html", gin.H{
			"User":  user,
			"Error": errMess,
		})
	})

	router.POST("/game_overview/:name/edit", func(c *gin.Context) {
		errMess = ""
		name := c.Param("name")
		newname := c.PostForm("name")
		user, err := sql.GetUserByName(context.Background(), client, name)
		if err != nil {
			errMess = "Error happened"
			log.Println(err)
		}

		usernew := ent.User{
			Name: newname,
		}
		err = sql.UpdateUser(context.Background(), user, usernew)
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
		c.HTML(http.StatusOK, "_game_inner.html", gin.H{
			"Users":    users,
			"Username": newname,
			"Games":    games,
			"Gamenum":  len(games),
			"Error":    errMess,
		})
	})

	router.GET("/game_overview/:name/:gameid/edit", func(c *gin.Context) {
		errMess = ""
		name := c.Param("name")
		gameid := c.Param("gameid")
		gameidInt, _ := strconv.Atoi(gameid)
		game, err := sql.GetGameByID(context.Background(), client, gameidInt)
		if err != nil {
			errMess = "Error happened"
			log.Println(err)
		}
		users, _, err := getGamesAndUsers(context.Background(), client)
		if err != nil {
			errMess = "Error happened"
			log.Println(err)
		}
		c.HTML(http.StatusOK, "_configGame.html", gin.H{
			"Username": name,
			"Users":    users,
			"Game":     game,
			"Error":    errMess,
		})
	})

	router.POST("/game_overview/:name/:gameid/edit", func(c *gin.Context) {
		errMess = ""
		name := c.Param("name")
		gameid := c.Param("gameid")
		gameidInt, _ := strconv.Atoi(gameid)
		game, err := sql.GetGameByID(context.Background(), client, gameidInt)
		if err != nil {
			errMess = "Error happened"
			log.Println(err)
		}

		user, err := sql.GetUserByName(context.Background(), client, name)
		if err != nil {
			errMess = "Error happened"
			log.Println(err)
		}
		gamenew := ent.Game{
			Name:  c.PostForm("name"),
			Type:  c.PostForm("gametype"),
			Cu:    c.PostForm("gamecu"),
			Notes: c.PostForm("notes"),
		}
		err = sql.UpdateGame(context.Background(), game, gamenew)
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
		c.HTML(http.StatusOK, "_game_inner.html", gin.H{
			"Users":    users,
			"Username": name,
			"Games":    games,
			"Gamenum":  len(games),
			"Error":    errMess,
		})
	})

	router.GET("/deleteUser/:name", func(c *gin.Context) {
		errMess = ""
		//TODO delete all Games too
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
		errMess = ""
		c.HTML(http.StatusOK, "_addUser.html", gin.H{
			"Error": errMess,
		})
	})

	router.POST("/addUser", func(c *gin.Context) {
		errMess = ""
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
		errMess = ""
		c.HTML(http.StatusNotFound, "_404.html", 1)
	})

	err = router.Run() // listen and serve on 0.0.0.0:8080
	if err != nil {
		log.Fatalf("failed starting router: %v", err)
	}
}
