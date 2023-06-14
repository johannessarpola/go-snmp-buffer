package app

import (
	"fmt"
	"os"

	"github.com/johannessarpola/go-network-buffer/internal/cli/listen"
	"github.com/johannessarpola/go-network-buffer/internal/cli/producer"
	"github.com/johannessarpola/go-network-buffer/internal/cli/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var appRootCmd = &cobra.Command{
	Use:   "app",
	Short: "app - CLI to run the apps",
	Long:  `description`,
}

func init() {

	cobra.OnInitialize(func() {
		utils.InitConfig(appRootCmd.Use, ".app", appRootCmd.Commands)
	})
	appRootCmd.PersistentFlags().String("configs", "",
		`Path for config`)
	viper.BindPFlag("configs", appRootCmd.PersistentFlags().Lookup("configs"))

	appRootCmd.AddCommand(listen.ListenCmd)
	appRootCmd.AddCommand(producer.ProducerCmd)
}

func Execute() {
	if err := appRootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "There was an error while executing your CLI '%s'", err)
		os.Exit(1)
	}
}
