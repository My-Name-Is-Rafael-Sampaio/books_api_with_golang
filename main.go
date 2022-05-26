package main

import (
	"github.com/My-Name-Is-Rafael-Sampaio/webapi-with-go/database"
	"github.com/My-Name-Is-Rafael-Sampaio/webapi-with-go/server"
)

func main() {
	database.StartDB()

	server := server.NewServer()

	server.Run()
}
