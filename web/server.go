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
		c.HTML(http.StatusOK, "_dashboard.html", gin.H{
			"Usernum": 10,
			"Gamenum": 10,
			"Error":   errMess,
		})
	})

	// Game overviews
	router.GET("/game_overview", func(c *gin.Context) {
		users, err := sql.GetAllUsers(context.Background(), client)
		if err != nil {
			errMess = "Error happened"
		}
		games, err := sql.GetAllGames(context.Background(), client)
		if err != nil {
			errMess = "Error happened"
		}
		c.HTML(http.StatusOK, "_game_overview.html", gin.H{
			"Users":   users,
			"Games":   games,
			"Gamenum": 10,
			"Error":   errMess,
		})
	})

	// Adding Users
	router.GET("/addUser", func(c *gin.Context) {
		c.HTML(http.StatusOK, "_add.html", gin.H{
			"Error":   errMess,
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
		c.HTML(http.StatusOK, "_dashboard.html", gin.H{
			"Usernum": 10,
			"Gamenum": 10,
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
