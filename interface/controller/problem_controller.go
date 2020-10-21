package controller

import (
	"github.com/tighug/sasa/domain/service"
	"github.com/tighug/sasa/interface/database"
	"github.com/tighug/sasa/usecase"
)

// ProblemController ...
type ProblemController struct {
	Interactor usecase.ProblemInteractor
}

// NewProblemController ...
func NewProblemController() *ProblemController {
	return &ProblemController{
		Interactor: usecase.ProblemInteractor{
			ProblemRepository: &database.ProblemRepository{},
		},
	}
}

// Init ...
func (controller *ProblemController) Init(srcDir, ansFile, configFile string) error {
	return service.Init(srcDir, ansFile, configFile)
}

// Encode ...
func (controller *ProblemController) Encode(srcDir, outDir string) error {
	probs, err := service.EncodeFiles(srcDir, outDir)
	if err != nil {
		return err
	}
	return controller.Interactor.SaveAll(probs)
}

// Compile ...
func (controller *ProblemController) Compile(srcDir, outDir string) error {
	probs, err := controller.Interactor.FindAll()
	if err != nil {
		return err
	}
	probs, err = service.CompileFiles(srcDir, outDir, probs)
	if err != nil {
		return err
	}
	return controller.Interactor.SaveAll(probs)
}

// Run ...
func (controller *ProblemController) Run(srcDir, outDir string) error {
	return service.RunFiles(srcDir, outDir)
}
