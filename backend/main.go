package main

import (
	"Todolist/config"
	"Todolist/database"
	"Todolist/router"
)

func main() {

	config.InitConfig()
	database.InitDB()

	router := router.SetupRouter()

	Port := config.AppConfig.App.Port
	router.Run(Port)

}
