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

var createCmd = &cobra.Command{
	Use:     "create",
	Aliases: []string{"c"},
	Short:   "creates index",
	Run: func(cmd *cobra.Command, args []string) {
		index := utils.GetFlagOrConfString(cmd, "index")
		path := utils.GetDataFromFlagOrConf(cmd)
		fmt.Printf("Sets index in database: %s\n", path)
		create_idx(path, index)
	},
}

func init() {
	createCmd.PersistentFlags().StringP("index", "i", viper.GetString("index"), "index name to set")
	viper.BindPFlag("index", deleteCmd.PersistentFlags().Lookup("index"))
}

func create_idx(path string, key string) {
	if len(key) == 0 {
		fmt.Println("Please provide a key for index")
	} else {
		err := c.WithDatabase(path, func(d *badger.DB) error {
			fmt.Printf("Creating index %s\n", key)
			err := db.CreateIndex(d, []byte(key))
			if err != nil {
				log.Fatal("Could not create index", err)
			}
			idx, _ := db.GetIndex(d, []byte(key)) // If it doesn't exist it should already return error in Set
			fmt.Printf("Created new index %s with initial value of %d", idx.Name, idx.Value)
			return nil
		})
		if err != nil {
			log.Fatal("Could not create index from database")
		}
	}
}
