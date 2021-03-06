package service

import (
	"os"
	"os/exec"
	"path"
	"strings"

	"github.com/spf13/afero"
)

// RunFiles ...
func RunFiles(srcDir, destDir, inputFile string) error {
	srcFileInfos, err := AFs.ReadDir(srcDir)
	if err != nil {
		return err
	}

	EnsureDir(destDir)

	if err != nil {
		return err
	}

	for _, srcFileInfo := range srcFileInfos {
		srcFileName := srcFileInfo.Name()
		if ext := path.Ext(srcFileName); ext == ".log" {
			continue
		}
		if err != nil {
			return err
		}

		srcPath := strings.Join([]string{srcDir, "/", srcFileName}, "")
		destPath := strings.Join([]string{destDir, "/", srcFileName, ".log"}, "")

		cmd := exec.Command(srcPath)
		input, err := getInput(inputFile)
		if err != nil {
			return err
		}
		if input != nil {
			cmd.Stdin = input
		}
		output, _ := cmd.CombinedOutput()
		file, err := os.Create(destPath)
		if err != nil {
			return err
		}
		defer file.Close()
		file.Write(output)
	}
	return err
}

func getInput(inputFile string) (afero.File, error) {
	exists, err := AFs.Exists(inputFile)
	if err != nil {
		return nil, err
	} else if exists {
		input, err := AFs.OpenFile(inputFile, os.O_RDONLY, os.ModePerm)
		if err != nil {
			return nil, err
		}
		return input, nil
	}
	return nil, nil
}
