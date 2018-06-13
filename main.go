package main

import (
	"log"

	"github.com/lxc/lxd/client"
	"github.com/lxc/lxd/shared/api"
)

func main() {
	c, err := lxd.ConnectLXDUnix("", nil)
	if err != nil {
		log.Fatal(err)
	}

	req := api.ContainersPost{
		Name: "test-container",
		Source: api.ContainerSource{
			Type: "image",
			Alias: "ubuntu-16-04",
		},
	}

	op, err := c.CreateContainer(req)
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

	op, err = c.UpdateContainerState("test-container", reqState, "")
	if err != nil {
		log.Fatal(err)
	}

	err = op.Wait()
	if err != nil {
		log.Fatal(err)
	}

}