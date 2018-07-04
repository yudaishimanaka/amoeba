package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// API for Container
func fetchAllContainer(c *gin.Context){
	c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "Fetch all containers successfully."})
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
