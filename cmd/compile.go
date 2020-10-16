package cmd

import (
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/tighug/sasa/interface/controller"
)

// compileCmd represents the compile command
var compileCmd = &cobra.Command{
	Use:   "compile",
	Short: "Compile .c files",
	Run: func(cmd *cobra.Command, args []string) {
		controller := controller.NewProblemController()
		if err := controller.Compile("./encoded/", "./build/"); err != nil {
			log.Err(err).Msg("")
		}
	},
}

func init() {
	rootCmd.AddCommand(compileCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// compileCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// compileCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
