package main

import (
	"github.com/gin-gonic/gin"
	"github.com/lxc/lxd/client"
	"net/http"
	"log"
)

// API for Container
func fetchAllContainer(c *gin.Context){
	// Fetch all container information.

	connection, err := lxd.ConnectLXDUnix("", nil)
	if err != nil {
		log.Fatal(err)
	}

	list, err := connection.GetContainers()
	if err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "list": list})
}

func fetchSingleContainer(c *gin.Context){

}

func createContainer(c *gin.Context){

}

func removeContainer(c *gin.Context){

}

func updateContainer(c *gin.Context){

}

// API for Images
func fetchImageList(c *gin.Context){

}

func createImage(c *gin.Context){

}

func updateImage(c *gin.Context){

}
