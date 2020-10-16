package usecase

import (
	"github.com/tighug/sasa/domain/model"
	"github.com/tighug/sasa/domain/repository"
)

// ProblemInteractor ...
type ProblemInteractor struct {
	ProblemRepository repository.ProblemRepository
}

// SaveAll ...
func (interactor *ProblemInteractor) SaveAll(probs model.Problems) error {
	return interactor.ProblemRepository.SaveAll(probs)
}

// FindAll ...
func (interactor *ProblemInteractor) FindAll() (model.Problems, error) {
	return interactor.ProblemRepository.FindAll()
}
