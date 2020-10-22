package cmd

import (
	"github.com/spf13/cobra"
	"github.com/tighug/sasa/interface/controller"
	"github.com/tighug/sasa/logger"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Execute binary files",
	Run: func(cmd *cobra.Command, args []string) {
		controller := controller.NewProblemController()
		if err := controller.Run(config.BuildDir, config.OutputDir, config.InputFile); err != nil {
			logger.Error(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
}
