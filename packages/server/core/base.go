package core

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/PureML-Inc/PureML/server/config"
	"github.com/PureML-Inc/PureML/server/core/settings"
	"github.com/PureML-Inc/PureML/server/daos"
	"github.com/PureML-Inc/PureML/server/tools/filesystem"
)

func Cleanup(optDataDir ...string) error {
	// remove data dir and all its contents
	var dataDir string
	if len(optDataDir) == 0 || optDataDir[0] == "" {
		// fallback to the default test data directory
		dataDir = config.GetDataDir()
	} else {
		dataDir = optDataDir[0]
	}
	if err := os.RemoveAll(dataDir); err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

var _ App = (*BaseApp)(nil)

// BaseApp implements core.App and defines the base PureBackend app structure.
type BaseApp struct {
	// configurable parameters
	isDebug      bool
	dataDir      string
	databaseType string

	// internals
	settings *settings.Settings
	dao      *daos.Dao
}

// BaseAppConfig defines a BaseApp configuration option
type BaseAppConfig struct {
	DataDir      string
	IsDebug      bool
	DatabaseType string
}

// NewBaseApp creates and returns a new BaseApp instance
// configured with the provided arguments.
//
// To initialize the app, you need to call `app.Bootstrap()`.
func NewBaseApp(appConfig *BaseAppConfig) *BaseApp {
	app := &BaseApp{
		dataDir:      appConfig.DataDir,
		isDebug:      appConfig.IsDebug,
		settings:     settings.New(),
		databaseType: config.GetDatabaseType(),
	}

	return app
}

// IsBootstrapped checks if the application was initialized
// (aka. whether Bootstrap() was called).
func (app *BaseApp) IsBootstrapped() bool {
	return app.dao != nil && app.settings != nil
}

// Bootstrap initializes the application
// (aka. create data dir, open db connections, load settings, etc.).
//
// It will call ResetBootstrapState() if the application was already bootstrapped.
func (app *BaseApp) Bootstrap() error {
	// clear resources of previous core state (if any)
	if err := app.ResetBootstrapState(); err != nil {
		return err
	}

	// ensure that data dir exist
	if err := os.MkdirAll(app.DataDir(), os.ModePerm); err != nil {
		return err
	}

	if err := app.initDataDB(); err != nil {
		return err
	}

	return nil
}

// ResetBootstrapState takes care for releasing initialized app resources
// (eg. closing db connections).
func (app *BaseApp) ResetBootstrapState() error {
	if app.Dao() != nil {
		if err := app.Dao().Close(); err != nil {
			return err
		}
		if err := app.Dao().Close(); err != nil {
			return err
		}
	}

	app.dao = nil
	app.settings = nil

	return nil
}

// Dao returns the default app Dao instance.
func (app *BaseApp) Dao() *daos.Dao {
	return app.dao
}

// DataDir returns the app data directory path.
func (app *BaseApp) DataDir() string {
	return app.dataDir
}

// DatabaseType returns the app data directory path.
func (app *BaseApp) DatabaseType() string {
	return app.databaseType
}

// IsDebug returns whether the app is in debug mode
// (showing more detailed error logs, executed sql statements, etc.).
func (app *BaseApp) IsDebug() bool {
	return app.isDebug
}

// Settings returns the loaded app settings.
func (app *BaseApp) Settings() *settings.Settings {
	return app.settings
}

// NewFilesystem creates a new local or S3 filesystem instance
// based on the current app settings.
//
// NB! Make sure to call `Close()` on the returned result
// after you are done working with it.
func (app *BaseApp) NewFilesystem() (*filesystem.System, error) {
	if app.settings.S3.Enabled {
		return filesystem.NewS3(
			app.settings.S3.Bucket,
			app.settings.S3.Region,
			app.settings.S3.Endpoint,
			app.settings.S3.AccessKey,
			app.settings.S3.Secret,
			app.settings.S3.ForcePathStyle,
		)
	}

	// fallback to local filesystem
	return filesystem.NewLocal(filepath.Join(app.DataDir(), "storage"))
}

func (app *BaseApp) initDataDB() error {
	dao, err := daos.InitDB(app.DatabaseType())
	if err != nil {
		return err
	}
	app.dao = dao
	return nil
}
