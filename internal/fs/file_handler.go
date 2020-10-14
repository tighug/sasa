package fs

import (
	"io/ioutil"
	"os"
	"os/exec"
)

// CompileFiles ...
func CompileFiles(srcDir, buildDir string) (err error) {
	srcFiles, err := readSrcFiles(srcDir)

	if err != nil {
		return err
	}

	if !dirExists(buildDir) {
		if err := makeDir(buildDir); err != nil {
			return err
		}
	}

	for _, srcFile := range srcFiles {
		srcFileName := srcFile.Name()
		srcPath := srcDir + srcFileName
		studentID := extractStudentID(srcFileName)
		buildPath := buildDir + studentID
		out, err := exec.Command("gcc", srcPath, "-o", buildPath).Output()

		if err != nil {
			return err
		}

		if out != nil {
			file, err := os.Create(buildPath + "_log")

			if err != nil {
				return err
			}

			defer file.Close()

			file.Write(out)
		}
	}

	return
}

func readSrcFiles(srcDir string) ([]os.FileInfo, error) {
	return ioutil.ReadDir(srcDir)
}

func makeDir(name string) error {
	return os.Mkdir(name, 0775)
}

func extractStudentID(s string) string {
	return s[0:7]
}

func fileExists(name string) bool {
	file, err := os.Stat(name)

	return err == nil && !file.IsDir()
}

func dirExists(name string) bool {
	file, err := os.Stat(name)

	return err == nil && file.IsDir()
}
