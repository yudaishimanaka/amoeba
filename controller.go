package main

import (
	"github.com/gin-gonic/gin"
	"github.com/lxc/lxd/client"
	"net/http"
	"log"
)

// Container
func FetchAllContainer(c *gin.Context){
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

func FetchSingleContainer(c *gin.Context) {
	// Fetch a container information.

	connection, err := lxd.ConnectLXDUnix("", nil)
	if err != nil {
		log.Fatal(err)
	}

	containerName := c.Param("containerName")

	info, _, err := connection.GetContainer(containerName)
	if err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "info": info})
}
