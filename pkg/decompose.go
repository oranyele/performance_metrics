package pkg

import (
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const errorExitCode = 1

var rootCmd = &cobra.Command{
	Use:   "msm",
	Short: "Mobile Size Metrics",
	Long:  `Tool created by the Performance Team that extracts size metrics from Android and iOS artifacts.`,
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.SetHelpCommand(&cobra.Command{Hidden: true})
	rootCmd.CompletionOptions.DisableDefaultCmd = true
}

func initConfig() {
	viper.AddConfigPath("$HOME")
	viper.SetConfigName("config")
	viper.SetConfigType("json")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err)
		os.Exit(errorExitCode)
	}
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
		os.Exit(errorExitCode)
	}
}
