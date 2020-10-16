package database

import (
	"os"

	"github.com/gocarina/gocsv"
	"github.com/tighug/sasa/domain/model"
)

// ProblemRepository ...
type ProblemRepository struct{}

// SaveAll ...
func (repo *ProblemRepository) SaveAll(probs model.Problems) error {
	// TODO: Avoid hard coding
	file, err := os.OpenFile("./database.csv", os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		return err
	}
	return gocsv.MarshalFile(&probs, file)
}

// FindAll ...
func (repo *ProblemRepository) FindAll() (model.Problems, error) {
	// TODO: Avoid hard coding
	file, err := os.OpenFile("./database.csv", os.O_WRONLY|os.O_CREATE, os.ModePerm)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var probs model.Problems
	if err := gocsv.UnmarshalFile(file, &probs); err != nil {
		return probs, err
	}
	return probs, err
}
