package service

import (
	"bufio"
	"strconv"
	"strings"

	"github.com/tighug/sasa/domain/model"
)

// CheckFiles ...
func CheckFiles(srcDir, ansFile string, probs model.Problems) (model.Problems, error) {
	srcFileInfos, err := AFs.ReadDir(srcDir)
	if err != nil {
		return nil, err
	}

	ansLines, err := readLine(ansFile)
	if err != nil {
		return nil, err
	}

	for _, srcFileInfo := range srcFileInfos {
		srcFileName := srcFileInfo.Name()
		srcPath := strings.Join([]string{srcDir, "/", srcFileName}, "")
		id, _ := strconv.Atoi(srcFileName[:7])

		probLines, err := readLine(srcPath)
		if err != nil {
			return nil, err
		}

		score := 0
		for i, pl := range ansLines {
			if pl == probLines[i] {
				score++
			}
		}

		for i, p := range probs {
			if p.ID == id {
				probs[i].Score = score
				break
			}
		}
	}
	return probs, err
}

func readLine(filePath string) ([]string, error) {
	file, err := AFs.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return lines, nil
}
