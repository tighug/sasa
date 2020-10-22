package cmd

import (
	"github.com/spf13/cobra"
	"github.com/tighug/sasa/interface/controller"
	"github.com/tighug/sasa/logger"
)

var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "Check output logs",
	Run: func(cmd *cobra.Command, args []string) {
		controller := controller.NewProblemController()
		if err := controller.Check(config.OutputDir, config.AnsFile); err != nil {
			logger.Error(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(checkCmd)
}
