package fs

import (
	"io/ioutil"
	"os"
	"os/exec"

	"github.com/moxtsuan/go-nkf"
	"github.com/rs/zerolog/log"
)

// EncodeFiles ...
func EncodeFiles(srcDir, outDir string) error {
	srcFiles, err := readSrcFiles(srcDir)

	if err != nil {
		return err
	}

	if !dirExists(outDir) {
		if err := makeDir(outDir); err != nil {
			return err
		}
	}

	for _, srcFile := range srcFiles {
		srcFileName := srcFile.Name()
		srcPath := srcDir + srcFileName
		studentID := srcFileName[0:7]
		// buildPath := outDir + studentID + ".c"

		file, err := os.Open(srcPath)
		// content, err := ioutil.ReadFile(srcPath)

		if err != nil {
			return err
		}

		result, err := nkf.Guess(file)

		if err != nil {
			result = "ShiftJIS maybe"
		}

		log.Info().Msg(studentID + ":" + result)

		// defer file.Close()

		// checkCharset(content, studentID)

		// reader := transform.NewReader(file, japanese.ShiftJIS.NewDecoder())

		// utf8File, err := os.Create(buildPath)

		// if err != nil {
		// 	return err
		// }

		// defer utf8File.Close()

		// tee := io.TeeReader(reader, utf8File)
		// scanner := bufio.NewScanner(tee)

		// for scanner.Scan() {
		// }
	}

	return err
}

// func checkCharset(content []byte, id string) {

// }

// CompileFiles ...
func CompileFiles(srcDir, outDir string) error {
	srcFiles, err := readSrcFiles(srcDir)

	if err != nil {
		return err
	}

	if !dirExists(outDir) {
		if err := makeDir(outDir); err != nil {
			return err
		}
	}

	for _, srcFile := range srcFiles {
		srcFileName := srcFile.Name()
		srcPath := srcDir + srcFileName
		studentID := srcFileName[0:7]
		buildPath := outDir + studentID
		cmd := exec.Command("gcc", srcPath, "-o", buildPath)
		out, err := cmd.CombinedOutput()

		if len(out) != 0 {
			file, err := os.Create(buildPath + ".log")

			if err != nil {
				log.Err(err).Msg(studentID)
			}

			defer file.Close()

			file.Write(out)
		}

		if err != nil {
			log.Err(err).Msg(studentID)
		}
	}

	return err
}

func readSrcFiles(srcDir string) ([]os.FileInfo, error) {
	return ioutil.ReadDir(srcDir)
}

func makeDir(name string) error {
	return os.Mkdir(name, 0775)
}

func fileExists(name string) bool {
	file, err := os.Stat(name)

	return err == nil && !file.IsDir()
}

func dirExists(name string) bool {
	file, err := os.Stat(name)

	return err == nil && file.IsDir()
}
