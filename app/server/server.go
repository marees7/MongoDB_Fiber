package server

import (
	"mongo_fiber/app/constant"
	"mongo_fiber/app/database"
	h "mongo_fiber/app/handler"

	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
)

func StartAPP() (*fiber.App, error) {
	db, err := database.NewDb()
	if err != nil {
		log.WithFields(log.Fields{
			"service": "serviceName",
			"err":     err,
		}).Error("failed to create db")
		return nil, err
	}
	// Create new router
	app := fiber.New()

	// Get all the routes from handler.
	_, err = h.GetRoutes(app, db)
	if err != nil {
		log.WithFields(log.Fields{
			"service": constant.ServiceName,
			"err":     err,
		}).Error("failed to get routes")
		return nil, err
	}
	return app, nil
}
