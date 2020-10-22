package cmd

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/tighug/sasa/interface/controller"
	"github.com/tighug/sasa/logger"
)

var allCmd = &cobra.Command{
	Use:   "all",
	Short: "Run a series of encode, build, run and check",
	Run: func(cmd *cobra.Command, args []string) {
		controller := controller.NewProblemController()

		bold := color.New(color.Bold).SprintFunc()
		step := color.New(color.FgHiBlack).SprintFunc()
		fmt.Printf("\n%s  %s\n", step("[1/4]"), bold("Encoding..."))
		if err := controller.Encode(config.SrcDir, config.EncodedDir); err != nil {
			logger.Error(err)
		}
		fmt.Printf("\n%s  %s\n", step("[2/4]"), bold("Building..."))
		if err := controller.Compile(config.EncodedDir, config.BuildDir); err != nil {
			logger.Error(err)
		}
		fmt.Printf("\n%s  %s\n", step("[3/4]"), bold("Running..."))
		if err := controller.Run(config.BuildDir, config.OutputDir, config.InputFile); err != nil {
			logger.Error(err)
		}
		fmt.Printf("\n%s  %s\n", step("[4/4]"), bold("Checking..."))
		if err := controller.Check(config.OutputDir, config.AnsFile); err != nil {
			logger.Error(err)
		}

		logger.Success("")
	},
}

func init() {
	rootCmd.AddCommand(allCmd)
}
