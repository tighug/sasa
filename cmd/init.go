package cmd

import (
	"github.com/spf13/cobra"
	"github.com/tighug/sasa/interface/controller"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize the current directory",
	Run: func(cmd *cobra.Command, args []string) {
		controller := controller.NewProblemController()
		controller.Init(config.SrcDir, config.AnsFile, configFile)
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
