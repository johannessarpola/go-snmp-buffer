package index

import (
	"fmt"

	"github.com/johannessarpola/go-network-buffer/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var setCmd = &cobra.Command{
	Use:     "set",
	Aliases: []string{"s"},
	Short:   "sets index",
	//Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("--- get index from database")
		index := utils.GetFlagOrConf(cmd, "index")
		path := utils.GetDataFromFlagOrConf(cmd)
		fmt.Printf("Sets index in database: %s\n", path)
		set_idx(path, index)
	},
}

func init() {
	getCmd.PersistentFlags().StringP("index", "i", viper.GetString("index"), "index name to set")
	getCmd.PersistentFlags().Int64P("value", "v", viper.GetInt64("value"), "Value to set the index to")
	viper.BindPFlag("index", deleteCmd.PersistentFlags().Lookup("index"))
}

func set_idx(path string, key string) {
	// if len(key) == 0 {
	// 	fmt.Println("Please provide a key for index")
	// } else {
	// 	fmt.Printf("Get index %s\n", key)
	// 	err := db.WithDatabase(path, func(d *badger.DB) error {
	// 		idx, err := db.GetIndex(d, []byte(key))
	// 		if err != nil {
	// 			log.Fatal("Could not get index from database")
	// 		}
	// 		fmt.Printf("Index: %s with value %d\n", idx.Name, idx.Value)
	// 		return err
	// 	})
	// 	if err != nil {
	// 		log.Fatal("Could not get index from database")
	// 	}
	// }
}