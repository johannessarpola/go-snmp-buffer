package index

import (
	"fmt"
	"log"

	"github.com/dgraph-io/badger/v4"
	c "github.com/johannessarpola/go-network-buffer/internal/cli/common"
	db "github.com/johannessarpola/go-network-buffer/pkg/index_db"
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
		index := utils.GetFlagOrConfString(cmd, "index")
		val := utils.GetFlagOrConfUint(cmd, "value")
		path := utils.GetDataFromFlagOrConf(cmd)
		fmt.Printf("Sets index in database: %s\n", path)
		set_idx(path, index, val)
	},
}

func init() {
	setCmd.PersistentFlags().StringP("index", "i", viper.GetString("index"), "index name to set")
	setCmd.PersistentFlags().Uint64P("value", "v", viper.GetUint64("value"), "Value to set the index to")
	viper.BindPFlag("index", deleteCmd.PersistentFlags().Lookup("index"))
}

func set_idx(path string, key string, value uint64) {

	if len(key) == 0 {
		fmt.Println("Please provide a key for index")
	} else {
		err := c.WithDatabase(path, func(d *badger.DB) error {
			fmt.Printf("Setting index %s to %d\n", key, value)
			err := db.SetIndex(d, []byte(key), value)
			if err != nil {
				log.Fatal("Could not set index", err)
			}
			idx, _ := db.GetIndex(d, []byte(key)) // If it doesn't exist it should already return error in Set
			fmt.Printf("New for index %s is %d", idx.Name, idx.Value)
			return nil
		})
		if err != nil {
			log.Fatal("Could not set index from database")
		}
	}
}
