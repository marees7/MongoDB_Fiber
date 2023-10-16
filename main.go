package main

import (
	"mongo_fiber/app/server"

	log "github.com/sirupsen/logrus"
)

func main() {
	app, err := server.StartAPP()
	if err != nil {
		log.WithFields(log.Fields{
			"service": "serviceName",
			"err":     err,
		}).Error("failed to start the application")
		return
	}
	app.Listen(":3030")
}
