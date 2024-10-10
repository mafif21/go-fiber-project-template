package main

import (
	"fmt"
	"go-fiber-project-template/internal/app"
	"go-fiber-project-template/internal/config"
)

func main() {
	viperConf := app.NewViper()
	logConf := app.NewLogger(viperConf)
	dbConf := config.NewDB(viperConf, logConf)
	appConf := app.NewFiber(viperConf)
	validatorConf := app.NewValidator(viperConf)

	serverConf := &app.StartAppConfig{
		App:       appConf,
		DB:        dbConf,
		Validator: validatorConf,
		Config:    viperConf,
		Log:       logConf,
	}

	app.StartApp(serverConf)

	webPort := viperConf.GetInt("PORT")
	err := appConf.Listen(fmt.Sprintf(":%d", webPort))
	if err != nil {
		logConf.Fatalf("Failed to start server: %v", err)
	}
}
