package cmd

import (
	"github.com/spf13/cobra"
	"github.com/tighug/sasa/interface/controller"
	"github.com/tighug/sasa/logger"
)

var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "Compile encoded source files",
	Run: func(cmd *cobra.Command, args []string) {
		controller := controller.NewProblemController()
		if err := controller.Compile(config.EncodedDir, config.BuildDir); err != nil {
			logger.Error(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(buildCmd)
}
