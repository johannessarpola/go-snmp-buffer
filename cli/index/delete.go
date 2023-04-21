package index

import (
	"fmt"

	"github.com/dgraph-io/badger/v4"
	"github.com/johannessarpola/go-network-buffer/db"
	"github.com/johannessarpola/go-network-buffer/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var deleteCmd = &cobra.Command{
	Use:     "delete",
	Aliases: []string{"l"},
	Short:   "deletes idx",
	//Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// Read the values from the configuration file (if not overridden by input arguments)
		index := utils.GetFlagOrConfString(cmd, "index")
		path := utils.GetDataFromFlagOrConf(cmd)
		delete_idx(path, index) // TODO Index from parent
	},
}

func init() {
	// Add flags to the root command
	deleteCmd.PersistentFlags().StringP("index", "i", viper.GetString("index"), "idex name to delete")
	// Bind the flags to the configuration parameters
	viper.BindPFlag("index", deleteCmd.PersistentFlags().Lookup("index"))

}

func delete_idx(path string, key string) {
	if len(key) == 0 {
		fmt.Println("Please provide a key for index")
	} else {
		fmt.Printf("Deleting index %s\n", key)
		db.WithDatabase(path, func(d *badger.DB) error {
			return db.DeleteIndex(d, []byte(key))
		})
	}

}
