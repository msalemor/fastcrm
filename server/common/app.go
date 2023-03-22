package common

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Settings struct {
	DbUser     string
	DbPassword string
	DbHostName string
	DbDatabase string
	DbPort     string
	AppPort    string
}

type App struct {
	FApp     *fiber.App
	Db       *gorm.DB
	Settings Settings
}

var app *App

// Get enviroment into settings
func (a *App) LoadEnvSettings() {
	godotenv.Load()
	a.Settings.DbUser = os.Getenv("DB_USER")
	a.Settings.DbPassword = os.Getenv("DB_PASSWORD")
	a.Settings.DbHostName = os.Getenv("DB_HOST")
	a.Settings.DbDatabase = os.Getenv("DB_DATABASE")
	a.Settings.DbPort = os.Getenv("DB_PORT")
	a.Settings.AppPort = os.Getenv("APP_PORT")
	log.Info("Environment variables loaded.")
}

// Get app instance
func getAppInstance() *App {
	if app == nil {
		app = &App{}

		app.Settings = Settings{}
		app.LoadEnvSettings()

		app.FApp = fiber.New()
		log.Info("Fiber app created.")

		// Mysql connection
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", app.Settings.DbUser, app.Settings.DbPassword, app.Settings.DbHostName, app.Settings.DbPort, app.Settings.DbDatabase)
		//log.Println(dsn)

		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

		if err != nil {
			log.Panic("Unable to connect to database.")
		}
		log.Info("Connected to MySql database.")
		app.Db = db

		log.Info("App initialized.")
	}
	return app
}

func GetApp() *fiber.App {
	return getAppInstance().FApp
}

func GetDb() *gorm.DB {
	return getAppInstance().Db
}

func GetAppSettings() Settings {
	return getAppInstance().Settings
}
