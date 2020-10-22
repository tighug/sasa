package cmd

import (
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/tighug/sasa/interface/controller"
)

var allCmd = &cobra.Command{
	Use:   "all",
	Short: "Run a series of encode, build, run and check",
	Run: func(cmd *cobra.Command, args []string) {
		controller := controller.NewProblemController()
		if err := controller.Encode(config.SrcDir, config.EncodedDir); err != nil {
			log.Err(err).Msg("error on encoding")
		}
		if err := controller.Compile(config.EncodedDir, config.BuildDir); err != nil {
			log.Err(err).Msg("error on building")
		}
		if err := controller.Run(config.BuildDir, config.OutputDir, config.InputFile); err != nil {
			log.Err(err).Msg("error on running")
		}
		if err := controller.Check(config.OutputDir, config.AnsFile); err != nil {
			log.Err(err).Msg("error on checking")
		}
	},
}

func init() {
	rootCmd.AddCommand(allCmd)
}
