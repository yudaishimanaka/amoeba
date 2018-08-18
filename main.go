package main

import (
	"net/http"
	"encoding/json"
	"io/ioutil"
	"log"
	"strconv"

	"./models"

	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/multitemplate"
	"github.com/go-xorm/xorm"
	"github.com/lxc/lxd/client"
	"github.com/lxc/lxd/shared/api"
	_ "github.com/go-sql-driver/mysql"
)

type Config struct {
	DbConfig Database `json:"db"`
}

type Database struct {
	User     string `json:"user"`
	Password string `json:"password"`
	DbName   string `json:"db_name"`
}

func initRender() multitemplate.Render {
	r := multitemplate.New()
	r.AddFromFiles("board", "templates/base.html", "templates/container-board.html")
	r.AddFromFiles("list", "templates/base.html", "templates/container-list.html")
	r.AddFromFiles("snapshot-lit", "templates/base.html", "templates/snapshot-list.html")
	return r
}

func FetchAllContainer(c *gin.Context) {
	lxdConn, err := lxd.ConnectLXDUnix("", nil)
	if err != nil {
		log.Fatal(err)
	}

	containers, err := lxdConn.GetContainers()
	if err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusCreated, containers)
}

func FetchSingleContainer(c *gin.Context) {
	lxdConn, err := lxd.ConnectLXDUnix("", nil)
	if err != nil {
		log.Fatal(err)
	}

	containerName := c.Params.ByName("containerName")

	container, _, err := lxdConn.GetContainer(containerName)
	if err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusCreated, container)

}

func CreateContainer(c *gin.Context) {
	lxdConn, err := lxd.ConnectLXDUnix("", nil)
	if err != nil {
		log.Fatal(err)
	}

	dbConn, ok := c.MustGet("databaseConn").(xorm.Engine)
	if !ok {
		log.Fatal(ok)
	}

	imageId, _ := strconv.Atoi(c.Params.ByName("imageId"))
	osType := c.Params.ByName("osType")
	cpu, _ := strconv.Atoi(c.Params.ByName("cpu"))
	memory, _ := strconv.Atoi(c.Params.ByName("memory"))
	disk, _ := strconv.Atoi(c.Params.ByName("disk"))

	var container = models.Container{ImageId: imageId, OsType: osType, Cpu: cpu, Memory: memory, Disk: disk}
	dbConn.Insert(&container)

	req := api.ContainersPost{
		Name: c.Params.ByName("containerName"),
		Source: api.ContainerSource{
			Type: "image",
			Alias: "alias", // TODO get imageName from imageId.
		},
	}

	op, err := lxdConn.CreateContainer(req)
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

	op, err = lxdConn.UpdateContainerState("imageName", reqState, "") // TODO use imageName
	if err != nil {
		log.Fatal(err)
	}

	// Wait for the operation to complete
	err = op.Wait()
	if err != nil {
		log.Fatal(err)
	}
	c.JSON(http.StatusCreated, "successfully")
}

func ApiMiddleware(engine *xorm.Engine) gin.HandlerFunc {
	return func(c *gin.Context) {
			c.Set("databaseConn", engine)
			c.Next()
	}
}

func main() {
	// Unmarshal config.json
	file, err := ioutil.ReadFile("config.json")
	if err != nil {
		log.Fatal(err)
	}

	var config Config
	json.Unmarshal(file, &config)
	connectString := config.DbConfig.User+":"+config.DbConfig.Password+"@/amoeba"
	engine, err := xorm.NewEngine("mysql", connectString)
	if err != nil {
		log.Fatal(err)
	}

	defer engine.Close()

	// Gin start
	r := gin.Default()
	r.LoadHTMLGlob("templates/*.html")
	r.HTMLRender = initRender()
	r.Use(ApiMiddleware(engine))

	r.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/board")
	})

	r.GET("/board", func(c *gin.Context) {
		c.HTML(http.StatusOK, "board", gin.H{})
	})

	r.GET("/list", func(c *gin.Context) {
		c.HTML(http.StatusOK, "list", gin.H{})
	})

	apis := r.Group("/api")
	{
		apis.GET("/containers", FetchAllContainer)
		apis.GET("/container/:name", FetchSingleContainer)
		apis.POST("/container/create", CreateContainer)
	}

	r.Static("/assets", "./assets")

	r.Run(":8080")
}
