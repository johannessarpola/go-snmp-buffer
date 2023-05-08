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

var getCmd = &cobra.Command{
	Use:     "get",
	Aliases: []string{"g"},
	Short:   "gets index",
	//Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("--- get index from database")
		index := utils.GetFlagOrConfString(cmd, "index")
		path := utils.GetDataFromFlagOrConf(cmd)
		fmt.Printf("Gets index from database: %s\n", path)
		get_idx(path, index)
	},
}

func init() {
	getCmd.PersistentFlags().StringP("index", "i", viper.GetString("index"), "index to get")
	viper.BindPFlag("index", deleteCmd.PersistentFlags().Lookup("index"))
}

func get_idx(path string, key string) {
	if len(key) == 0 {
		fmt.Println("Please provide a key for index")
	} else {
		fmt.Printf("Get index %s\n", key)
		err := c.WithDatabase(path, func(d *badger.DB) error {
			idx, err := db.GetIndex(d, []byte(key))
			if err != nil {
				log.Fatal("Could not get index from database")
			}
			fmt.Printf("Index: %s with value %d\n", idx.Name, idx.Value)
			return err
		})
		if err != nil {
			log.Fatal("Could not get index from database")
		}
	}
}
