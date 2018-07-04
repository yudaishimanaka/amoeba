package main

import (
	"github.com/gin-gonic/gin"
	"github.com/lxc/lxd/client"
	"net/http"
	"log"
)

// API for Container
func fetchAllContainer(c *gin.Context){
	c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "Fetch all containers successfully."})
}

func fetchSingleContainer(c *gin.Context){
	container, err := lxd.ConnectLXDUnix("", nil)
	if err != nil {
		log.Fatal(err)
	}

	list, err := container.GetContainers()
	if err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "list": list})
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
