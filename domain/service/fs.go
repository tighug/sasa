package service

import (
	"os"

	"github.com/rs/zerolog/log"
	"github.com/spf13/afero"
)

// AFs ...
var AFs = &afero.Afero{Fs: afero.NewOsFs()}

// EnsureDir ...
func EnsureDir(path string) error {
	if exists, err := AFs.DirExists(path); err != nil {
		return err
	} else if !exists {
		defer log.Info().Msgf("Created an empty directory named %q.", path)
		return AFs.MkdirAll(path, os.ModePerm)
	}
	log.Warn().Msgf("%q already exists.", path)
	return nil
}

// EnsureFile ...
func EnsureFile(path string) (afero.File, error) {
	if exists, err := AFs.Exists(path); err != nil {
		return nil, err
	} else if !exists {
		defer log.Info().Msgf("Created an empty file named %q.", path)
		return AFs.Create(path)
	}
	log.Warn().Msgf("%q already exists.", path)
	return nil, nil
}
