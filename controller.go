package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

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
