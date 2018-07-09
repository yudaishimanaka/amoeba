package main

import (
	"net/http"
	"html/template"

	"github.com/gin-gonic/gin"
)

func main() {
	// Gin start
	r := gin.Default()
	r.LoadHTMLGlob("templates/*.html")

	assets := r.Group("/")
	{
		assets.GET("/", func(c *gin.Context) {
			c.Redirect(http.StatusMovedPermanently, "/container-board")
		})

		assets.GET("/container-board", func(c *gin.Context) {
			html := template.Must(template.ParseFiles("templates/base.html", "templates/container-board.html"))
			r.SetHTMLTemplate(html)
			c.HTML(http.StatusOK, "base.html", gin.H{"containers": "aaaaa"})
		})

		assets.GET("/container-list", func(c *gin.Context) {
			html := template.Must(template.ParseFiles("templates/base.html", "templates/container-list.html"))
			r.SetHTMLTemplate(html)
			c.HTML(http.StatusOK, "base.html", gin.H{})
		})

		assets.GET("/snapshot-list", func(c *gin.Context) {
			html := template.Must(template.ParseFiles("templates/base.html", "templates/snapshot-list.html"))
			r.SetHTMLTemplate(html)
			c.HTML(http.StatusOK, "base.html", gin.H{})
		})
	}

	v1 := r.Group("/api/v1")
	{
		v1.GET("/fetchAllContainer", fetchAllContainer)
		v1.GET("/fetchSingleContainer", fetchSingleContainer)
		v1.POST("/createContainer/:name", createContainer)
		v1.PUT("/updateContainer/:name", updateContainer)
		v1.DELETE("/removeContainer/:name", removeContainer)
	}

	r.Static("/assets", "./assets")

	r.Run(":8080")
}
