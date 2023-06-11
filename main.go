package main

import (
	"finalProject/database"
	"finalProject/router"
)

func main() {
	database.StartDB()

	var PORT = ":2009"

	router.StartServer().Run(PORT)
}
