package main

import (
	"makromusic/handler"
	"makromusic/pkg/database"
)

func main() {

	database.GetConnect()
	handler.Handler()

}
