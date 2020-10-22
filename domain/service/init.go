package service

import (
	"github.com/tighug/sasa/logger"
)

const defaultConfig = (`SrcDir: src
EncodedDir: encoded
BuildDir: build
OutputDir: output
AnsFile: answer.txt
DBFile: db.csv
InputFile: input.txt`)

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
		logger.Info("The config file \"" + configFile + "\" has been created")
		return nil
	}
	logger.Warn("The config file \"" + configFile + "\" already exists")
	return nil
}
