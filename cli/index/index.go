package admin

import (
	"github.com/spf13/cobra"
)

var IndexCmd = &cobra.Command{
	Use:     "index",
	Aliases: []string{"idxs"},
	Short:   "commands related to indexes",
	//Args:    cobra.ExactArgs(1),
	// Run: func(cmd *cobra.Command, args []string) {
	// 	res := stringer.Reverse(args[0])
	// 	fmt.Println(res)
	// },
}

func init() {
	IndexCmd.AddCommand(getCmd)
	IndexCmd.AddCommand(listCmd)
}
