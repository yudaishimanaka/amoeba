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

	r.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/container-list")
	})

	r.GET("/container-board", func(c *gin.Context) {
		html := template.Must(template.ParseFiles("templates/base.html", "templates/container-board.html"))
		r.SetHTMLTemplate(html)
		c.HTML(http.StatusOK, "base.html", gin.H{})
	})

	r.GET("/container-list", func(c *gin.Context) {
		html := template.Must(template.ParseFiles("templates/base.html", "templates/container-list.html"))
		r.SetHTMLTemplate(html)
		c.HTML(http.StatusOK, "base.html", gin.H{})
	})

	r.GET("/snapshot-list", func(c *gin.Context) {
		html := template.Must(template.ParseFiles("templates/base.html", "templates/snapshot-list.html"))
		r.SetHTMLTemplate(html)
		c.HTML(http.StatusOK, "base.html", gin.H{})
	})

	r.Static("/assets", "./assets")

	r.Run(":8080")
}
