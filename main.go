package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/multitemplate"
)

func initRender() multitemplate.Render {
	r := multitemplate.New()
	r.AddFromFiles("board", "templates/base.html", "templates/container-board.html")
	r.AddFromFiles("list", "templates/base.html", "templates/container-list.html")
	r.AddFromFiles("snapshot-lit", "templates/base.html", "templates/snapshot-list.html")
	return r
}

func main() {
	// Gin start
	r := gin.Default()
	r.LoadHTMLGlob("templates/*.html")
	r.HTMLRender = initRender()

	r.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/board")
	})

	r.GET("/board", func(c *gin.Context) {
		c.HTML(http.StatusOK, "board", FetchAllContainer)
	})

	r.GET("/list", func(c *gin.Context) {
		c.HTML(http.StatusOK, "list", gin.H{})
	})

	container := r.Group("/container")
	{
		container.GET("/all", FetchAllContainer)
	}

	r.Static("/assets", "./assets")

	r.Run(":8080")
}
