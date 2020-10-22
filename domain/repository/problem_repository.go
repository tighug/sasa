package repository

import "github.com/tighug/sasa/domain/model"

// ProblemRepository ...
type ProblemRepository interface {
	SaveAll(model.Problems) error
	FindAll() (model.Problems, error)
}
