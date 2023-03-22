package main

import (
	"fastcrm/common"
	"fmt"

	log "github.com/sirupsen/logrus"
)

// main is the entry point for the application
func main() {

	// Prepare the database if needed
	common.GormMigrate()

	// Get the fiber app
	app := common.GetApp()

	// Map the routes
	common.MapRoutes(app, common.GetDb())

	// Start server
	addr := fmt.Sprintf(":%s", common.GetAppSettings().AppPort)

	log.Fatal(app.Listen(addr))
}
