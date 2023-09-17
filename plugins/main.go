package main

import (
	"plugins/plugins/ccdata"

	_ "github.com/rs/zerolog/log"
)

func main() {
	// Create a server, expose plugin as grpc methods which invokes db service to write data
	// Use redis to store count of plugin executions in last 1hour, separate db service
	// github integration plugin and chesscom data plugin
	// Run the main logic of the application
	ccdata.Main()
}
