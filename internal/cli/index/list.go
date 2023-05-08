package index

import (
	"fmt"
	"log"

	"github.com/dgraph-io/badger/v4"
	c "github.com/johannessarpola/go-network-buffer/internal/cli/common"
	db "github.com/johannessarpola/go-network-buffer/pkg/index_db"
	"github.com/johannessarpola/go-network-buffer/utils"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"l"},
	Short:   "lists indexes",
	//Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Indexes in database: ")
		path := utils.GetDataFromFlagOrConf(cmd)
		fmt.Printf("Listing indexes in database: %s\n", path)
		c.WithDatabase(path, cli_list_idx)
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
