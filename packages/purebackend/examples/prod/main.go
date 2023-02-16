package main

import (
	"log"
	"os"

	"github.com/PureML-Inc/PureML/purebackend"
	"github.com/PureML-Inc/PureML/purebackend/config"
	"github.com/PureML-Inc/PureML/purebackend/core/settings"
)

func main() {
	app := purebackend.NewWithConfig(&purebackend.Config{
		DatabaseType: config.GetDatabaseType(),
		DatabaseUrl:  config.GetDatabaseURL(),
		Settings: &settings.Settings{
			R2: settings.R2Config{
				Enabled:   true,
				AccountId: os.Getenv("PURE_R2_ACCOUNT_ID"),
				Bucket:    os.Getenv("PURE_R2_BUCKET"),
				Endpoint:  os.Getenv("PURE_R2_ENDPOINT"),
				AccessKey: os.Getenv("PURE_R2_ACCESS_KEY"),
				Secret:    os.Getenv("PURE_R2_SECRET"),
			},
			AdminAuthToken: settings.TokenConfig{
				Secret: os.Getenv("PURE_ADMIN_AUTH_TOKEN_SECRET"),
			},
		},
	})
	if err := app.Bootstrap(); err != nil {
		log.Fatal(err)
	}
	// Manually seed db if admin does not exist
	// Set admin details in .env file
	//
	// Refer:= URL_TO_README_GOES_HERE
	app.Dao().Datastore().SeedAdminIfNotExists()
	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
