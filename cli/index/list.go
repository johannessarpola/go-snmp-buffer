package admin

import (
	"fmt"
	"log"

	"github.com/johannessarpola/go-network-buffer/db"
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
		list_indexes("../../_idxs")
	},
}

// "../../_idxs"
func list_indexes(path string) {
	fs, err := utils.NewFileStore(path)
	if err != nil {
		log.Fatal("could not open index filestore")
	}

	c := db.ListIndexes(fs)
	for _, idx := range c {
		if idx != nil {
			fmt.Printf("Index: %s with value %d\n", idx.Name, idx.Value)
		}
	}
}
