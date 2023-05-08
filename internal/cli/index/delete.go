package index

import (
	"fmt"

	"github.com/dgraph-io/badger/v4"
	u "github.com/johannessarpola/go-network-buffer/internal/cli/utils"
	db "github.com/johannessarpola/go-network-buffer/pkg/indexdb"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var deleteCmd = &cobra.Command{
	Use:     "delete",
	Aliases: []string{"d"},
	Short:   "deletes idx",
	//Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// Read the values from the configuration file (if not overridden by input arguments)
		index := u.GetFlagOrConfString(cmd, "index")
		path := u.GetDataFromFlagOrConf(cmd)
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
		u.WithDatabase(path, func(d *badger.DB) error {
			return db.DeleteIndex(d, []byte(key))
		})
	}

}
