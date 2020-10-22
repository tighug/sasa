package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Sasa",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("v0.0.0")
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
