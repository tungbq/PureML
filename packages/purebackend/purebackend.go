package purebackend

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/PureML-Inc/PureML/purebackend/apis"
	"github.com/PureML-Inc/PureML/purebackend/core"
	"github.com/PureML-Inc/PureML/purebackend/core/settings"
)

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	contact@pureml.com

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@securityDefinitions.apikey	ApiKeyAuth
//	@in							header
//	@name						Authorization
//	@description				Header for logged in user format: Bearer {token}

var _ core.App = (*PureBackend)(nil)

// Version of PureBackend
var Version = "(untracked)"

// appWrapper serves as a private core.App instance wrapper.
type appWrapper struct {
	core.App
}

// PureBackend defines a PureBackend app launcher.
//
// It implements [core.App] via embedding and all of the app interface methods
// could be accessed directly through the instance (eg. PureBackend.DataDir()).
type PureBackend struct {
	*appWrapper

	debugFlag       bool
	dataDirFlag     string
	hideStartBanner bool
}

// Config is the PureBackend initialization config struct.
type Config struct {
	// optional default values for the console flags
	DefaultDebug   bool
	DefaultDataDir string // if not set, it will fallback to "./pureml_data"
	DatabaseType   string // if not set, it will fallback to "sqlite3"
	DatabaseUrl    string // if not set, it will fallback to "file:./pureml_data/pureml.db"
	Settings       *settings.Settings

	// hide the default console server info on app startup
	HideStartBanner bool
}

// New creates a new PureBackend instance with the default configuration.
// Use [NewWithConfig()] if you want to provide a custom configuration.
//
// Note that the application will not be initialized/bootstrapped yet,
// aka. DB connections, migrations, app settings, etc. will not be accessible.
// Everything will be initialized when [Start()] is executed.
// If you want to initialize the application before calling [Start()],
// then you'll have to manually call [Bootstrap()].
func New() *PureBackend {
	_, isUsingGoRun := inspectRuntime()

	return NewWithConfig(&Config{
		DefaultDebug: isUsingGoRun,
	})
}

// NewWithConfig creates a new PureBackend instance with the provided config.
//
// Note that the application will not be initialized/bootstrapped yet,
// aka. DB connections, migrations, app settings, etc. will not be accessible.
// Everything will be initialized when [Start()] is executed.
// If you want to initialize the application before calling [Start()],
// then you'll have to manually call [Bootstrap()].
func NewWithConfig(config *Config) *PureBackend {
	if config == nil {
		panic("missing config")
	}

	// initialize a default data directory based on the executable baseDir
	if config.DefaultDataDir == "" {
		baseDir, _ := inspectRuntime()
		config.DefaultDataDir = filepath.Join(baseDir, "pureml_data")
	}

	pb := &PureBackend{
		debugFlag:       config.DefaultDebug,
		dataDirFlag:     config.DefaultDataDir,
		hideStartBanner: config.HideStartBanner,
	}

	// initialize the app instance
	pb.appWrapper = &appWrapper{core.NewBaseApp(&core.BaseAppConfig{
		DataDir:      pb.dataDirFlag,
		IsDebug:      pb.debugFlag,
		DatabaseType: config.DatabaseType,
		DatabaseUrl:  config.DatabaseUrl,
		Settings:     config.Settings,
	})}

	return pb
}

// Start starts the application
func (pb *PureBackend) Start() error {
	return apis.Serve(pb, pb.hideStartBanner)
}

// inspectRuntime tries to find the base executable directory and how it was run.
func inspectRuntime() (baseDir string, withGoRun bool) {
	if strings.HasPrefix(os.Args[0], os.TempDir()) {
		// probably ran with go run
		withGoRun = true
		baseDir, _ = os.Getwd()
	} else {
		// probably ran with go build
		withGoRun = false
		baseDir = filepath.Dir(os.Args[0])
	}
	return
}
