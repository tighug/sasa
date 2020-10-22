package service

import (
	"bufio"
	"io"
	"strconv"
	"strings"

	"github.com/gogs/chardet"
	"github.com/tighug/sasa/domain/model"
	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
)

// EncodeFiles ...
func EncodeFiles(srcDir, destDir string) (model.Problems, error) {
	srcFileInfos, err := AFs.ReadDir(srcDir)
	if err != nil {
		return nil, err
	}
	if err := EnsureDir(destDir); err != nil {
		return nil, err
	}

	var probs model.Problems

	for _, srcFileInfo := range srcFileInfos {
		srcFileName := srcFileInfo.Name()
		id := srcFileName[:7]
		slice := strings.Split(srcFileName[8:], "_")
		name := slice[0]
		srcPath := strings.Join([]string{srcDir, "/", srcFileName}, "")
		destFileName := strings.Join([]string{id, name}, "_")
		destPath := strings.Join([]string{destDir, "/", destFileName, ".c"}, "")

		charset := "UTF-8"
		result, err := detectCharset(srcPath)
		if err != nil {
			return nil, err
		}

		srcFile, err := AFs.Open(srcPath)
		if err != nil {
			return nil, err
		}
		defer srcFile.Close()

		destFile, err := AFs.Create(destPath)
		if err != nil {
			return nil, err
		}
		defer destFile.Close()

		if result.Charset == charset {
			io.Copy(destFile, srcFile)
		} else {
			// TODO: Decrease denpedencies
			charset = "Shft_JIS"
			reader := transform.NewReader(srcFile, japanese.ShiftJIS.NewDecoder())
			tee := io.TeeReader(reader, destFile)
			scanner := bufio.NewScanner(tee)
			for scanner.Scan() {
			}
		}

		var prob model.Problem

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
	content, err := AFs.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	textDetector := chardet.NewTextDetector()
	return textDetector.DetectBest(content)
}
