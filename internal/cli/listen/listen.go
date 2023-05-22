package listen

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var listenCmd = &cobra.Command{
	Use:     "listen",
	Aliases: []string{"l"},
	Short:   "commands related to indexes",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("runnista")

	},
}

func init() {

	listenCmd.PersistentFlags().Int("port", 10161, "port to listen to")
	viper.BindPFlag("listen.port", listenCmd.PersistentFlags().Lookup("port"))

	listenCmd.PersistentFlags().String("host", "0.0.0.0", "host to listen")
	viper.BindPFlag("listen.host", listenCmd.PersistentFlags().Lookup("port"))

	listenCmd.PersistentFlags().String("data.snmp", "_snmp", "folder to use for snmp data")
	viper.BindPFlag("listen.data.snmp", listenCmd.PersistentFlags().Lookup("data.snmp"))

	listenCmd.PersistentFlags().String("data.index", "_idxs", "folder to use for snmp data")
	viper.BindPFlag("listen.data.index", listenCmd.PersistentFlags().Lookup("data.index"))
}
