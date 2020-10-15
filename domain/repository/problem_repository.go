package repository

import "github.com/tighug/sasa/domain/model"

// ProblemRepository ...
type ProblemRepository interface {
	SaveAll([]model.Problem) error
	FindAll() ([]model.Problem, error)
}
