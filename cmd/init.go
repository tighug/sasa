package cmd

import (
	"github.com/spf13/cobra"
	"github.com/tighug/sasa/interface/controller"
	"github.com/tighug/sasa/logger"
)

const configFile = ".sasarc.yaml"

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize the current directory",
	Run: func(cmd *cobra.Command, args []string) {
		controller := controller.NewProblemController()
		if err := controller.Init(config.SrcDir, config.AnsFile, configFile); err != nil {
			logger.Error(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
