package main

import (
	"GoTasks/Task8/internal/rest/app"
	"GoTasks/task8/internal/rest/config"
	"log"
)

func main() {
	if err := app.Run(config.NewConfig()); err != nil {
		log.Println("error app.Run(): ", err)
	}
}
