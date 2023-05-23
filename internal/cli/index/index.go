package index

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var dataIndexKey = "admincli.data.index"
var IndexCmd = &cobra.Command{
	Use:     "index",
	Aliases: []string{"idxs"},
	Short:   "commands related to indexes",
	Args:    cobra.ExactArgs(1),
	//	Run: func(cmd *cobra.Command, args []string) {
	//		fmt.Println("runnista")

	// },
}

func init() {

	getCmd.PersistentFlags().String("data.index", "_idxs", "folder to use for snmp data")
	viper.BindPFlag(dataIndexKey, getCmd.PersistentFlags().Lookup("data.index"))

	IndexCmd.AddCommand(getCmd)
	IndexCmd.AddCommand(setCmd)
	IndexCmd.AddCommand(listCmd)
	IndexCmd.AddCommand(deleteCmd)
	IndexCmd.AddCommand(createCmd)
}
