package cmd

import (
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/tighug/sasa/interface/controller"
)

// encodeCmd represents the encode command
var encodeCmd = &cobra.Command{
	Use:   "encode",
	Short: "Encode source files",
	Run: func(cmd *cobra.Command, args []string) {
		controller := controller.NewProblemController()
		if err := controller.Encode(config.SrcDir, config.EncodedDir); err != nil {
			log.Err(err).Msg("")
		}
	},
}

func init() {
	rootCmd.AddCommand(encodeCmd)
}
