package app

import (
	"fmt"
	"os"

	"github.com/johannessarpola/go-network-buffer/internal/cli/listen"
	"github.com/johannessarpola/go-network-buffer/internal/cli/producer"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var appRootCmd = &cobra.Command{
	Use:   "app",
	Short: "app - CLI to run the apps",
	Long:  `description`,
}

func init() {
	viper.SetConfigName("app-config") // name of config file (without extension)
	viper.SetConfigType("yaml")       // type of config file
	viper.AddConfigPath(".")          // path to look for the config file in the working directory

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Using default values")
	} else {
		fmt.Printf("Using config file of: %s\n", viper.GetViper().ConfigFileUsed())
		fmt.Println(viper.AllSettings())
	}

	appRootCmd.AddCommand(listen.ListenCmd)
	appRootCmd.AddCommand(producer.ProducerCmd)
}

func Execute() {
	if err := appRootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "There was an error while executing your CLI '%s'", err)
		os.Exit(1)
	}
}
