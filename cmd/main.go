package main

import (
	"makromusic/handler"
	"makromusic/pkg/database"
	"makromusic/pkg/service"
)

func main() {

	database.GetConnect()
	go service.Consume()
	handler.Handler()

}
