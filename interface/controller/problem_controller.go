package controller

import (
	"github.com/tighug/sasa/domain/service/fs"
	"github.com/tighug/sasa/usecase"
)

// ProblemController ...
type ProblemController struct {
	Interactor usecase.ProblemInteractor
}

// NewProblemController ...
func NewProblemController() *ProblemController {
	return &ProblemController{}
}

// Encode ...
func (controller *ProblemController) Encode() {
	fs.EncodeFiles("./src/", "./encoded/")
}

// Compile ...
func (controller *ProblemController) Compile() {
	controller.Interactor.FindAll()
}

// Run ...
func (controller *ProblemController) Run() {}
