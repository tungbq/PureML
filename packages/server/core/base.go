package core

import (
	"os"

	"github.com/PureML-Inc/PureML/server/config"
)

func Bootstrap(optDataDir ...string) error {
	// ensure that data dir exist
	var dataDir string
	if len(optDataDir) == 0 || optDataDir[0] == "" {
		// fallback to the default test data directory
		dataDir = config.GetDataDir()
	} else {
		dataDir = optDataDir[0]
	}
	if err := os.MkdirAll(dataDir, os.ModePerm); err != nil {
		return err
	}
	return nil
}

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
		return err
	}
	return nil
}