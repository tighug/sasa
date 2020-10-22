package cmd

import (
	"github.com/spf13/cobra"
	"github.com/tighug/sasa/interface/controller"
	"github.com/tighug/sasa/logger"
)

var encodeCmd = &cobra.Command{
	Use:   "encode",
	Short: "Encode source files",
	Run: func(cmd *cobra.Command, args []string) {
		controller := controller.NewProblemController()
		if err := controller.Encode(config.SrcDir, config.EncodedDir); err != nil {
			logger.Error(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(encodeCmd)
}
