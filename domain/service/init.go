package service

import (
	"github.com/rs/zerolog/log"
)

const defaultConfig = (`SrcDir: src
EncodedDir: encoded
BuildDir: build
OutputDir: output
AnsFile: answer.txt
DBFile: db.csv`)

// Init ...
func Init(srcDir, ansFile, configFile string) error {
	EnsureDir(srcDir)
	EnsureFile(ansFile)

	if exists, err := AFs.Exists(configFile); err != nil {
		return err
	} else if !exists {
		file, err := AFs.Create(configFile)
		if err != nil {
			return err
		}
		if _, err := file.Write([]byte(defaultConfig)); err != nil {
			return err
		}
		log.Info().Msgf("Generated a default config file named %q.", configFile)
		return nil
	}
	log.Warn().Msgf("%q already exists.", configFile)
	return nil
}
