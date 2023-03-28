package main

import (
	"log"
	"os"

	"github.com/PureMLHQ/PureML/packages/purebackend"
	"github.com/PureMLHQ/PureML/packages/purebackend/core/settings"
)

func main() {
	app := purebackend.NewWithConfig(&purebackend.Config{
		DefaultDebug: true,
		DatabaseType: "sqlite3",
		// DatabaseUrl:  If using postgres, set this to the postgres connection string
		Settings: &settings.Settings{
			// set the default settings
			S3: settings.S3Config{
				Enabled:   false,
				Bucket:    os.Getenv("PURE_S3_BUCKET"),
				Region:    os.Getenv("PURE_S3_REGION"),
				Endpoint:  os.Getenv("PURE_S3_ENDPOINT"),
				AccessKey: os.Getenv("PURE_S3_ACCESS_KEY"),
				Secret:    os.Getenv("PURE_S3_SECRET"),
			},
			R2: settings.R2Config{
				Enabled:   false,
				AccountId: os.Getenv("PURE_R2_ACCOUNT_ID"),
				Bucket:    os.Getenv("PURE_R2_BUCKET"),
				Endpoint:  os.Getenv("PURE_R2_ENDPOINT"),
				AccessKey: os.Getenv("PURE_R2_ACCESS_KEY"),
				Secret:    os.Getenv("PURE_R2_SECRET"),
			},
			Search: settings.SearchConfig{
				Enabled: false,
			},
			Site: settings.SiteConfig{
				BaseURL: os.Getenv("PURE_SITE_BASE_URL"),
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
	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
