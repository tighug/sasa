package cmd

import (
	"os"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Create src, resource and answer.txt",
	Run: func(cmd *cobra.Command, args []string) {
		if err := os.Mkdir("src", 0775); err != nil {
			log.Err(err).Msg("Cannot create src.")
		} else {
			log.Info().Msg("Created src.")
		}

		if err := os.Mkdir("resource", 0775); err != nil {
			log.Err(err).Msg("Cannot create resource.")
		} else {
			log.Info().Msg("Created resource.")
		}

		if file, err := os.Create("answer.txt"); err != nil {
			log.Err(err).Msg("Cannot create " + file.Name())
		} else {
			log.Info().Msg("Created " + file.Name())
		}
	},
}

func init() {
	rootCmd.AddCommand(initCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
