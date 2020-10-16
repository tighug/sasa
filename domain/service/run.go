package service

import (
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"strings"
)

// RunFiles ...
func RunFiles(srcDir, destDir string) error {
	srcFileInfos, err := ioutil.ReadDir(srcDir)
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

		srcPath := strings.Join([]string{srcDir, srcFileName}, "")
		destPath := strings.Join([]string{destDir, srcFileName, ".log"}, "")

		cmd := exec.Command(srcPath)
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
