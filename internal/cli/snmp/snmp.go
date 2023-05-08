package snmp

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var SnmpCmd = &cobra.Command{
	Use:     "snmp",
	Aliases: []string{},
	Short:   "commands related to snmp database",
	Args:    cobra.ExactArgs(1),
	//	Run: func(cmd *cobra.Command, args []string) {
	//		fmt.Println("runnista")

	// },
}

func init() {

	SnmpCmd.PersistentFlags().StringP("data", "d", "", "folder of database")
	viper.BindPFlag("data", SnmpCmd.PersistentFlags().Lookup("data"))
	SnmpCmd.AddCommand(lastnCmd)

}
