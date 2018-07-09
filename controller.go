package main

import (
	"github.com/gin-gonic/gin"
	"github.com/lxc/lxd/client"
	"net/http"
	"log"
	"github.com/lxc/lxd/shared/api"
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
	// Fetch a container information.

	connection, err := lxd.ConnectLXDUnix("", nil)
	if err != nil {
		log.Fatal(err)
	}

	container, _, err := connection.GetContainer(c.Params.ByName("name"))

	c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "container": container})
}

func createContainer(c *gin.Context){
	// Create a container.

	connection, err := lxd.ConnectLXDUnix("", nil)
	if err != nil {
		log.Fatal(err)
	}

	req := api.ContainersPost{
		Name: "test",
		Source: api.ContainerSource{
			Type: "image",
			Alias: "ubuntu-16-04",
		},
	}

	op, err := connection.CreateContainer(req)
	if err != nil {
		log.Fatal(err)
	}

	// Wait op to complete.
	opErr := op.Wait()
	if opErr != nil {
		log.Fatal(opErr)
	}

	reqState := api.ContainerStatePut{
		Action: "start",
		Timeout: -1,
	}

	op, err = connection.UpdateContainerState("test", reqState, "" )
	if err != nil {
		log.Fatal(err)
	}

	// Wait op to complete.
	opErr = op.Wait()
	if opErr != nil {
		log.Fatal(opErr)
	}
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
