package cmd

import (
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/tighug/sasa/interface/controller"
)

var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "Check output logs",
	Run: func(cmd *cobra.Command, args []string) {
		controller := controller.NewProblemController()
		if err := controller.Check(config.OutputDir, config.AnsFile); err != nil {
			log.Err(err).Msg("")
		}
	},
}

func init() {
	rootCmd.AddCommand(checkCmd)
}
