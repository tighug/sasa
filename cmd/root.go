package cmd

import (
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// Config ...
type Config struct {
	SrcDir     string
	EncodedDir string
	BuildDir   string
	OutputDir  string
	AnsFile    string
	DBFile     string
}

const configFile = ".sasarc.yaml"

var config Config

var rootCmd = &cobra.Command{
	Use:   "sasa",
	Short: "Sasa is an Auntomatic Scoring Application for PandA-1",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Err(err)
	}
}

func init() {
	initLogger()
	cobra.OnInitialize(initConfig)
}

func initLogger() {
	output := zerolog.ConsoleWriter{Out: os.Stdout}
	output.FormatTimestamp = func(i interface{}) string {
		return ""
	}
	log.Logger = log.Output(output)
}

func initConfig() {
	viper.AddConfigPath(".")
	viper.SetConfigName(configFile)

	viper.SetDefault("SrcDir", "src")
	viper.SetDefault("EncodedDir", "encoded")
	viper.SetDefault("BuildDir", "build")
	viper.SetDefault("OutputDir", "output")
	viper.SetDefault("AnsFile", "answer.txt")
	viper.SetDefault("DBFile", "db.csv")

	viper.AutomaticEnv() // read in environment variables that match

	if err := viper.ReadInConfig(); err == nil {
		log.Debug().Msgf("A config file is found. Sasa use %q.", viper.ConfigFileUsed())
	} else {
		log.Debug().Msg("No config files in the current directory. Sasa use default.")
	}

	viper.Unmarshal(&config)
}
