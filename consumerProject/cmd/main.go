package main

import (
	"consumerProject/pkg/database"
	"consumerProject/pkg/service"
)

func main() {

	database.GetConnect()
	service.Consume()
}
