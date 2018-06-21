package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

type Config struct {
	DbConfig Database `json:"db"`
}

type Database struct {
	User     string `json:"user"`
	Password string `json:"password"`
	DbName   string `json:"db_name"`
}

type User struct {
	UserId   uint64 `xorm:"not null BIGINT pk autoincr 'user_id'"`
	UserName string `xorm:"not null unique 'user_name'"`
	Email    string `xorm:"not null TEXT 'email'"`
	IconPath string `xorm:"null TEXT 'icon_path'"`
}

func initDatabase(driver, user, password, dbname string, config Config) (e *xorm.Engine, err error) {
	engine, err := xorm.NewEngine(driver, user+":"+password+"@/")
	if err != nil {
		return nil, err
	}

	if _, err := engine.Exec("CREATE DATABASE " + dbname); err != nil {
		log.Printf("Database already exists.")
		return engine, nil
	} else {
		engine.Exec("USE " + dbname)
		engine.CreateTables(User{})
		log.Printf("Success initialize.")

		return engine, nil
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

	// Init database
	engine, err := initDatabase("mysql", config.DbConfig.User, config.DbConfig.Password, config.DbConfig.DbName, config)
	if err != nil {
		log.Fatal(engine)
	}

	// Gin start
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")

	r.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/container-list")
	})

	r.GET("/container-list", func(c *gin.Context) {
		c.HTML(http.StatusOK, "container-list.html", gin.H{"title": "Amoeba - Container-list"})
	})

	r.Static("/assets", "./assets")

	r.Run(":8080")
}
