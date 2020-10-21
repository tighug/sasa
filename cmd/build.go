package cmd

import (
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/tighug/sasa/interface/controller"
)

// compileCmd represents the compile command
var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "Compile encoded source files",
	Run: func(cmd *cobra.Command, args []string) {
		controller := controller.NewProblemController()
		if err := controller.Compile(config.EncodedDir, config.BuildDir); err != nil {
			log.Err(err).Msg("")
		}
	},
}

func init() {
	rootCmd.AddCommand(buildCmd)
}
