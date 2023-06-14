package admin

import (
	"fmt"
	"os"

	i "github.com/johannessarpola/go-network-buffer/internal/cli/index"
	s "github.com/johannessarpola/go-network-buffer/internal/cli/snmp"
	"github.com/johannessarpola/go-network-buffer/internal/cli/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var adminRootCmd = &cobra.Command{
	Use:   "admin",
	Short: "admin - CLI to view the database",
	Long: `description
stuff
stuff`,
	// Run: func(cmd *cobra.Command, args []string) {
	// 	fmt.Println("Started cli")
	// },
}

func init() {
	cobra.OnInitialize(func() {
		utils.InitConfig(adminRootCmd.Use, ".admin", adminRootCmd.Commands)
	})
	adminRootCmd.PersistentFlags().String("configs", "",
		`Path for config`)
	viper.BindPFlag("configs", adminRootCmd.PersistentFlags().Lookup("configs"))

	adminRootCmd.AddCommand(i.IndexCmd)
	adminRootCmd.AddCommand(s.SnmpCmd)
}

func Execute() {
	if err := adminRootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "There was an error while executing your CLI '%s'", err)
		os.Exit(1)
	}
}
