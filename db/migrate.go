package main

import (
	"github.com/go-xorm/xorm"
	_ "github.com/go-sql-driver/mysql"
	"../models"
	"io/ioutil"
	"log"
	"encoding/json"
)

type Config struct {
	DbConfig Database `json:"db"`
}

type Database struct {
	User     string `json:"user"`
	Password string `json:"password"`
	DbName   string `json:"db_name"`
}

func initDatabase(driver, user, password, dbname string) (e *xorm.Engine, err error) {
	engine, err := xorm.NewEngine(driver, user+":"+password+"@/")
	if err != nil {
		return nil, err
	}

	if _, err := engine.Exec("CREATE DATABASE " + dbname); err != nil {
		log.Printf("Database already exists.")
		return engine, nil
	} else {
		engine.Exec("USE " + dbname)
		engine.CreateTables(models.Container{})
		log.Printf("Success initialize.")

		return engine, nil
	}
}

func main(){
	// Unmarshal config.json
	file, err := ioutil.ReadFile("../config.json")
	if err != nil {
		log.Fatal(err)
	}

	var config Config
	json.Unmarshal(file, &config)

	// Init database
	engine, err := initDatabase("mysql", config.DbConfig.User, config.DbConfig.Password, config.DbConfig.DbName)
	if err != nil {
		log.Fatal(engine)
	}
}
