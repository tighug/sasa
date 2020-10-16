package service

import (
	"io/ioutil"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/tighug/sasa/domain/model"
)

// CompileFiles ...
func CompileFiles(srcDir, outDir string, probs model.Problems) (model.Problems, error) { // TODO: Change val passing to ref passing
	srcFileInfos, err := ioutil.ReadDir(srcDir)
	if err != nil {
		return nil, err
	}
	EnsureDir(outDir)

	for _, srcFileInfo := range srcFileInfos {
		srcFileName := srcFileInfo.Name()
		outFileName := strings.Split(srcFileName, ".")[0]
		srcPath := strings.Join([]string{srcDir, srcFileName}, "")
		outPath := strings.Join([]string{outDir, outFileName}, "")
		canCompile := true
		id, _ := strconv.Atoi(srcFileName[:7])

		cmd := exec.Command("gcc", srcPath, "-o", outPath)
		output, err := cmd.CombinedOutput()
		if len(output) != 0 {
			file, err := os.Create(strings.Join([]string{outPath, ".log"}, ""))
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
