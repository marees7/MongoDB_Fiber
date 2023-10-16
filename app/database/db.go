package database

import (
	"mongo_fiber/app/constant"
	"os"

	"github.com/eaciit/dbox"
	_ "github.com/eaciit/dbox/dbc/mongo"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

func init() {
	log.Info("Starting server...")
	err := godotenv.Load()
	if err != nil {
		log.WithFields(log.Fields{
			"service": constant.ServiceName,
			"err":     err,
		}).Warn("failed to load env")
	}
}

func DSN() *dbox.ConnectionInfo {
	dbHost := os.Getenv("DB_HOST")
	dbUsername := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")
	ci := &dbox.ConnectionInfo{
		Host:     dbHost,
		Database: dbName,
		UserName: dbUsername,
		Password: dbPassword,
		Settings: nil,
	}
	return ci
}

func NewDb() (dbox.IConnection, error) {
	conn, err := dbox.NewConnection("mongo", DSN())
	if err != nil {
		panic("Connect Failed") // Change with your error handling
	}

	err = conn.Connect()
	if err != nil {
		panic("Connect Failed") // Change with your error handling
	}
	return conn, nil
}
