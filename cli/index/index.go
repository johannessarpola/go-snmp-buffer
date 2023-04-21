package index

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

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

	IndexCmd.PersistentFlags().StringP("data", "d", "", "folder of database")
	viper.BindPFlag("data", IndexCmd.PersistentFlags().Lookup("data"))

	IndexCmd.AddCommand(getCmd)
	IndexCmd.AddCommand(setCmd)
	IndexCmd.AddCommand(listCmd)
	IndexCmd.AddCommand(deleteCmd)
	IndexCmd.AddCommand(createCmd)
}
