package web

import (
	"lending-system/ent"
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
		c.HTML(http.StatusOK, "_dashboard.html", gin.H{
			"Projectnum": 10,
			"Packagenum": 10,
			"Error":      errMess,
		})
	})

	// Startpage _dashboard.html
	router.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "_dashboard.html", gin.H{
			"Projectnum": 10,
			"Packagenum": 10,
			"Error":      errMess,
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
