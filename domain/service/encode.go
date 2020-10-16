package service

import (
	"bufio"
	"io"
	"io/ioutil"
	"os"
	"strconv"
	"strings"

	"github.com/gogs/chardet"
	"github.com/tighug/sasa/domain/model"
	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
)

// EncodeFiles ...
func EncodeFiles(srcDir, outDir string) (model.Problems, error) {
	srcFileInfos, err := ioutil.ReadDir(srcDir)
	if err != nil {
		return nil, err
	}
	ensureDir(outDir)

	var probs model.Problems

	for _, srcFileInfo := range srcFileInfos {
		srcFileName := srcFileInfo.Name()
		id := srcFileName[:7]
		slice := strings.Split(srcFileName[8:], "_")
		name := slice[0]
		srcPath := strings.Join([]string{srcDir, srcFileName}, "")
		outFileName := strings.Join([]string{id, name}, "_")
		outPath := strings.Join([]string{outDir, outFileName, ".c"}, "")

		charset := "UTF-8"
		result, err := detectCharset(srcPath)
		if err != nil {
			return nil, err
		}

		srcFile, err := os.Open(srcPath)
		if err != nil {
			return nil, err
		}
		defer srcFile.Close()

		outFile, err := os.Create(outPath)
		if err != nil {
			return nil, err
		}
		defer outFile.Close()

		if result.Charset == charset {
			io.Copy(outFile, srcFile)
		} else {
			// TODO: Decrease denpedencies
			charset = "Shft_JIS"
			reader := transform.NewReader(srcFile, japanese.ShiftJIS.NewDecoder())
			tee := io.TeeReader(reader, outFile)
			s := bufio.NewScanner(tee)
			for s.Scan() {
			}
		}

		prob := model.Problem{}
		prob.Name = name
		prob.Charset = result.Charset
		prob.ID, err = strconv.Atoi(id)
		if err != nil {
			return nil, err
		}

		probs = append(probs, prob)
	}
	return probs, err
}

func detectCharset(filePath string) (*chardet.Result, error) {
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	textDetector := chardet.NewTextDetector()
	return textDetector.DetectBest(content)
}
