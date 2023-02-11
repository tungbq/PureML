package core

import (
	"github.com/PureML-Inc/PureML/server/core/settings"
	"github.com/PureML-Inc/PureML/server/daos"
	"github.com/PureML-Inc/PureML/server/tools/filesystem"
)

// App defines the main PureBackend app interface.
type App interface {
	// Dao returns the default app Dao instance.
	//
	// This Dao could operate only on the tables and models
	// associated with the default app database. For example,
	// trying to access the request logs table will result in error.
	Dao() *daos.Dao

	// DataDir returns the app data directory path.
	DataDir() string

	// IsDebug returns whether the app is in debug mode
	// (showing more detailed error logs, executed sql statements, etc.).
	IsDebug() bool

	// Settings returns the loaded app settings.
	Settings() *settings.Settings

	// NewFilesystem creates and returns a configured filesystem.System instance.
	//
	// NB! Make sure to call `Close()` on the returned result
	// after you are done working with it.
	NewFilesystem() (*filesystem.System, error)
	
	// UploadFile uploads a file to the app storage.
	UploadFile(file *filesystem.File, basePath string) (string, error)

	// IsBootstrapped checks if the application was initialized
	// (aka. whether Bootstrap() was called).
	IsBootstrapped() bool

	// Bootstrap takes care for initializing the application
	// (open db connections, load settings, etc.).
	//
	// It will call ResetBootstrapState() if the application was already bootstrapped.
	Bootstrap() error

	// ResetBootstrapState takes care for releasing initialized app resources
	// (eg. closing db connections).
	ResetBootstrapState() error
}
