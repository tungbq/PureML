package core

import (
	"os"

	"github.com/PureML-Inc/PureML/server/config"
)

func Bootstrap() error {
	// ensure that data dir exist
	dataDir := config.GetDataDir()
	if err := os.MkdirAll(dataDir, os.ModePerm); err != nil {
		return err
	}
	return nil
}