package index

import (
	"fmt"
	"log"

	"github.com/dgraph-io/badger/v4"
	u "github.com/johannessarpola/go-network-buffer/internal/cli/utils"
	db "github.com/johannessarpola/go-network-buffer/pkg/indexdb"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var listCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"l"},
	Short:   "lists indexes",
	//Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Indexes in database: ")
		path := viper.GetString(dataIndexKey)
		fmt.Printf("Listing indexes in database: %s\n", path)
		u.WithDatabase(path, cli_list_idx)
	},
}

func cli_list_idx(dbi *badger.DB) error {
	c, err := db.ListIndexes(dbi)
	if err != nil {
		log.Fatal("Error with listing", err)
	}
	for _, idx := range c {
		if idx != nil {
			fmt.Printf("Index: %s with value %d\n", idx.Name, idx.Value)
		}
	}
	return nil
}
