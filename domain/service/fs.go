package service

import (
	"os"

	"github.com/spf13/afero"
	"github.com/tighug/sasa/logger"
)

// AFs ...
var AFs = &afero.Afero{Fs: afero.NewOsFs()}

// EnsureDir ...
func EnsureDir(path string) error {
	if exists, err := AFs.DirExists(path); err != nil {
		return err
	} else if !exists {
		defer logger.Info("The directory \"" + path + "\" has been created.")
		return AFs.MkdirAll(path, os.ModePerm)
	}
	logger.Warn("The directory \"" + path + "\" already exists.")
	return nil
}

// EnsureFile ...
func EnsureFile(path string) (afero.File, error) {
	if exists, err := AFs.Exists(path); err != nil {
		return nil, err
	} else if !exists {
		defer logger.Info("The file \"" + path + "\" has been created")
		return AFs.Create(path)
	}
	logger.Warn("The directory \"" + path + "\" already exists")
	return nil, nil
}
