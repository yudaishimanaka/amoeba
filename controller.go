package main

import (
	"net/http"
	"log"

	"github.com/lxc/lxd/shared/api"
	"github.com/gin-gonic/gin"
	"github.com/lxc/lxd/client"
)

type CreateRequest struct {
	ContainerName string `json:"container_name"`
	AliasName	  string `json:"alias_name"`
}

// Container
func CreateContainer(c *gin.Context) {
	// Create a container

	connection, err := lxd.ConnectLXDUnix("", nil)
	if err != nil {
		log.Fatal(err)
	}

	var json CreateRequest
	ret := c.Bind(&json)
	log.Println(ret)

	req := api.ContainersPost{
		Name: json.ContainerName,
		Source: api.ContainerSource{
			Type: "image",
			Alias: json.AliasName,
		},
	}

	op, err := connection.CreateContainer(req)
	if err != nil {
		log.Fatal(err)
	}

	err = op.Wait()
	if err != nil {
		log.Fatal(err)
	}

	reqState := api.ContainerStatePut{
		Action: "start",
		Timeout: -1,
	}

	op, err = connection.UpdateContainerState(json.ContainerName, reqState, "")
	if err != nil {
		log.Fatal(err)
	}

	err = op.Wait()
	if err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusCreated, gin.H{})
}

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
