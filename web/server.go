package web

import (
	"context"
	"lending-system/ent"
	"lending-system/sql"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

var errMess string

/*
StartServer starts the web server

client: client that requested the database
*/
func StartServer(client *ent.Client) {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.Static("/assets", "./assets")
	router.Static("/db", "./db")
	router.Static("/tmp", "./tmp")
	router.MaxMultipartMemory = 8 << 20

	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("mysession", store))

	var err error
	password := "SeanJonas"
	adminpw := "Deskbot002"
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
		if givenPassword == adminpw {
			session := sessions.Default(c)
			session.Set("login", "true")
			session.Set("admin", "true")
			session.Save()
			c.Redirect(http.StatusFound, "/index")
		} else if givenPassword != password {
			errMess = "Falsches Passwort"
			c.HTML(http.StatusOK, "_login.html", gin.H{
				"Error": errMess,
			})
		} else {
			session := sessions.Default(c)
			session.Set("login", "true")
			session.Set("admin", "false")
			session.Save()
			c.Redirect(http.StatusFound, "/index")
		}
	})

	// logout
	router.GET("/logout", func(c *gin.Context) {
		session := sessions.Default(c)
		session.Set("login", "false")
		session.Set("admin", "false")
		session.Save()
		c.Redirect(http.StatusFound, "/")
	})

	// Startpage _dashboard.html
	router.GET("/index", func(c *gin.Context) {
		users, games, err := getGamesAndUsers(context.Background(), client)
		if err != nil {
			errMess = "Error happened"
			log.Println(err)
		}
		c.HTML(http.StatusOK, "_dashboard.html", gin.H{
			"Admin":   sessions.Default(c).Get("admin"),
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
		sortArrays(users, games)
		c.HTML(http.StatusOK, "_game_overview.html", gin.H{
			"Admin":   sessions.Default(c).Get("admin"),
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
		sortArrays(users, games)
		c.HTML(http.StatusOK, "_game_inner.html", gin.H{
			"Admin":    sessions.Default(c).Get("admin"),
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
			"Admin":    sessions.Default(c).Get("admin"),
			"Username": name,
			"Error":    errMess,
		})
	})

	router.POST("/game_overview/:name/addGame", func(c *gin.Context) {
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
		c.Redirect(http.StatusFound, "/game_overview/"+username)
	})

	// config views
	router.GET("/game_overview/:name/edit", func(c *gin.Context) {
		name := c.Param("name")
		user, err := sql.GetUserByName(context.Background(), client, name)
		if err != nil {
			errMess = "Error happened"
			log.Println(err)
		}
		c.HTML(http.StatusOK, "_configUser.html", gin.H{
			"Admin": sessions.Default(c).Get("admin"),
			"User":  user,
			"Error": errMess,
		})
	})

	router.POST("/game_overview/:name/edit", func(c *gin.Context) {
		name := c.Param("name")
		newname := c.PostForm("name")

		user, err := sql.GetUserByName(context.Background(), client, name)
		if err != nil {
			errMess = "Error happened"
			log.Println(err)
		}

		picname, err := changePicture(c, user.Picture)
		if err != nil {
			errMess = "Error happened"
			log.Println(err)
		}

		usernew := ent.User{
			Name:    newname,
			Picture: picname,
		}
		err = sql.UpdateUser(context.Background(), user, usernew)
		if err != nil {
			errMess = "Error happened"
			log.Println(err)
		}
		c.Redirect(http.StatusFound, "/game_overview/"+newname)
	})

	router.GET("/game_overview/:name/:gameid/edit", func(c *gin.Context) {
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
			"Admin":    sessions.Default(c).Get("admin"),
			"Username": name,
			"Users":    users,
			"Game":     game,
			"Error":    errMess,
		})
	})

	router.POST("/game_overview/:name/:gameid/edit", func(c *gin.Context) {
		name := c.Param("name")
		gameid := c.Param("gameid")
		gameidInt, _ := strconv.Atoi(gameid)
		game, err := sql.GetGameByID(context.Background(), client, gameidInt)
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
		c.Redirect(http.StatusFound, "/game_overview/"+name)
	})

	router.GET("/deleteUser/:name", func(c *gin.Context) {
		name := c.Param("name")
		user, err := sql.GetUserByName(context.Background(), client, name)
		if err != nil {
			errMess = "Error happened"
			log.Println(err)
		}
		err = cleanDeleteUser(context.Background(), client, user)
		if err != nil {
			errMess = "Error happened"
			log.Println(err)
		}
		c.Redirect(http.StatusFound, "/game_overview")
	})

	router.GET("/deleteGame/:gameid", func(c *gin.Context) {
		gameid := c.Param("gameid")
		gameidInt, _ := strconv.Atoi(gameid)
		game, err := sql.GetGameByID(context.Background(), client, gameidInt)
		if err != nil {
			errMess = "Error happened"
			log.Println(err)
		}
		err = sql.DeleteGame(context.Background(), client, game)
		if err != nil {
			errMess = "Error happened"
			log.Println(err)
		}
		c.Redirect(http.StatusFound, "/game_overview")
	})

	// Adding Users
	router.GET("/addUser", func(c *gin.Context) {
		c.HTML(http.StatusOK, "_addUser.html", gin.H{
			"Admin": sessions.Default(c).Get("admin"),
			"Error": errMess,
			"Name":  "",
		})
	})

	router.POST("/addUser", func(c *gin.Context) {
		username := c.PostForm("name")
		user := ent.User{
			Name: username,
		}
		_, err = sql.AddUser(context.Background(), client, user)
		if err != nil {
			if strings.Contains(err.Error(), "users.name") {
				errMess = "Ein Nutzer mit dem Namen existiert bereits, bitte gebe einen anderen Namen ein"
			} else {
				errMess = "Error happened"
			}
			c.HTML(http.StatusOK, "_addUser.html", gin.H{
				"Admin": sessions.Default(c).Get("admin"),
				"Error": errMess,
				"Name":  username,
			})
		}
		c.Redirect(http.StatusFound, "/game_overview")
	})

	// Ausgeliehen
	router.GET("/lended", func(c *gin.Context) {
		users, games, err := getGamesAndUsers(context.Background(), client)
		if err != nil {
			errMess = "Error happened"
			log.Println(err)
		}
		games = removeUnlended(games)
		sortArrays(users, games)
		c.HTML(http.StatusOK, "_lended_overview.html", gin.H{
			"Admin":   sessions.Default(c).Get("admin"),
			"Users":   users,
			"Games":   games,
			"Gamenum": len(games),
			"Error":   errMess,
		})
	})

	// Backup
	router.GET("/backup", func(c *gin.Context) {
		users, err := sql.GetAllUsers(context.Background(), client)
		if err != nil {
			errMess = "Error happened"
			log.Println(err)
		}
		err = zipper(users)
		if err != nil {
			errMess = "Error happened"
			log.Println(err)
		}
		c.FileAttachment("./tmp/backup.zip", "backup.zip")
		c.Redirect(http.StatusFound, "/index")
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
