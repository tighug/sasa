package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/tighug/sasa/logger"
)

// Config ...
type Config struct {
	SrcDir     string
	EncodedDir string
	BuildDir   string
	OutputDir  string
	AnsFile    string
	DBFile     string
	InputFile  string
}

var config Config

var rootCmd = &cobra.Command{
	Use:   "sasa",
	Short: "Sasa is an Auntomatic Scoring Application for PandA-1",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		logger.Error(err)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
}

func initConfig() {
	viper.AddConfigPath(".")
	viper.SetConfigName(".sasarc")
	viper.SetConfigType("yaml")

	viper.SetDefault("SrcDir", "src")
	viper.SetDefault("EncodedDir", "encoded")
	viper.SetDefault("BuildDir", "build")
	viper.SetDefault("OutputDir", "output")
	viper.SetDefault("AnsFile", "answer.txt")
	viper.SetDefault("DBFile", "database.csv")
	viper.SetDefault("InputFile", "input.txt")

	viper.AutomaticEnv()
	viper.Unmarshal(&config)
}
