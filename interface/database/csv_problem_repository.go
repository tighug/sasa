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
	defer file.Close()
	if err := file.Truncate(0); err != nil {
		return err
	}
	return gocsv.MarshalFile(&probs, file)
}

// FindAll ...
func (repo *ProblemRepository) FindAll() (model.Problems, error) {
	// TODO: Avoid hard coding
	file, err := os.OpenFile("./database.csv", os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	probs := []*model.Problem{}
	if err := gocsv.UnmarshalFile(file, &probs); err != nil {
		return nil, err
	}
	newProbs := model.Problems{}
	for _, prob := range probs {
		newProbs = append(newProbs, model.Problem{
			ID:         prob.ID,
			Name:       prob.Name,
			CanCompile: prob.CanCompile,
			Charset:    prob.Charset,
			Scores:     prob.Scores,
		})
	}
	return newProbs, err
}
