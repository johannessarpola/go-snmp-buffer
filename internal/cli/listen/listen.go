package listen

import (
	"fmt"

	sdb "github.com/johannessarpola/go-network-buffer/pkg/snmpdb"

	"github.com/johannessarpola/go-network-buffer/pkg/snmp"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var ListenCmd = &cobra.Command{
	Use:     "listen",
	Aliases: []string{"l"},
	Short:   "commands related to indexes",
	Run: func(cmd *cobra.Command, args []string) {
		port := viper.GetInt("listen.port")
		host := viper.GetString("listen.host")

		f_snmp := viper.GetString("listen.data.snmp.folder")
		p_snmp := viper.GetString("listen.data.snmp.prefix")
		f_idxs := viper.GetString("listen.data.index.folder")
		fmt.Printf("Using SNMP databse on folder %s\n", f_snmp)
		fmt.Printf("Using Index databse on folder %s\n", f_idxs)

		snmp_data, err := sdb.OpenDatabases(f_snmp, p_snmp, f_idxs)
		if err != nil {
			logrus.Fatal("Could not open databases", err)
		}
		defer snmp_data.Close()
		fmt.Printf("Listening to %s:%d\n", host, port)
		snmp.ListenSnmp(port, host, snmp_data)
	},
}

func init() {

	ListenCmd.PersistentFlags().Int("port", 10161, "port to listen to")
	viper.BindPFlag("listen.port", ListenCmd.PersistentFlags().Lookup("port"))

	ListenCmd.PersistentFlags().String("host", "0.0.0.0", "host to listen")
	viper.BindPFlag("listen.host", ListenCmd.PersistentFlags().Lookup("port"))

	ListenCmd.PersistentFlags().String("data.snmp.folder", "_snmp", "folder to use for snmp data")
	viper.BindPFlag("listen.data.snmp.folder", ListenCmd.PersistentFlags().Lookup("data.snmp.folder"))

	ListenCmd.PersistentFlags().String("data.snmp.prefix", "snmp_", "folder to use for snmp data")
	viper.BindPFlag("listen.data.snmp.prefix", ListenCmd.PersistentFlags().Lookup("data.snmp.prefix"))

	ListenCmd.PersistentFlags().String("data.index.folder", "_idxs", "folder to use for snmp data")
	viper.BindPFlag("listen.data.index.folder", ListenCmd.PersistentFlags().Lookup("data.index.folder"))
}
