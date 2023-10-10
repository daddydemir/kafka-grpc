package main

import (
	"producerProject/handler"
	"producerProject/pkg/database"
)

func main() {

	database.GetConnect()
	handler.Handler()

}
