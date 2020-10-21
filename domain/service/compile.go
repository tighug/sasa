package service

import (
	"os/exec"
	"strconv"
	"strings"

	"github.com/tighug/sasa/domain/model"
)

// CompileFiles ...
func CompileFiles(srcDir, destDir string, probs model.Problems) (model.Problems, error) { // TODO: Change val passing to ref passing
	srcFileInfos, err := AFs.ReadDir(srcDir)
	if err != nil {
		return nil, err
	}
	if err := EnsureDir(destDir); err != nil {
		return nil, err
	}

	for _, srcFileInfo := range srcFileInfos {
		srcFileName := srcFileInfo.Name()
		destFileName := strings.Split(srcFileName, ".")[0]
		srcPath := strings.Join([]string{srcDir, "/", srcFileName}, "")
		destPath := strings.Join([]string{destDir, "/", destFileName}, "")
		canCompile := true
		id, _ := strconv.Atoi(srcFileName[:7])

		cmd := exec.Command("gcc", srcPath, "-o", destPath)
		output, err := cmd.CombinedOutput()
		if len(output) != 0 {
			file, err := AFs.Create(strings.Join([]string{destPath, ".log"}, ""))
			if err != nil {
				return nil, err
			}
			defer file.Close()
			file.Write(output)
		}
		if err != nil {
			canCompile = false
		}

		for i, p := range probs {
			if p.ID == id {
				probs[i].CanCompile = canCompile
				break
			}
		}
	}

	return probs, err
}
