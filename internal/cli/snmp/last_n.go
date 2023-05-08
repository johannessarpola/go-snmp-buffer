package snmp

import (
	"fmt"

	"github.com/dgraph-io/badger/v4"
	c "github.com/johannessarpola/go-network-buffer/internal/cli/common"
	m "github.com/johannessarpola/go-network-buffer/pkg/models"
	db "github.com/johannessarpola/go-network-buffer/pkg/snmp_db"
	u "github.com/johannessarpola/go-network-buffer/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const number string = "number"
const number_short string = "n"

var lastnCmd = &cobra.Command{
	Use:     "last",
	Aliases: []string{"l"},
	Short:   "latest n traps",
	//Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		path := u.GetDataFromFlagOrConf(cmd)
		fmt.Printf("Listing snmp data from database: %s\n", path)
		n := u.GetFlagOrConfInt(cmd, number)
		last_n(path, n)

	},
}

func init() {
	// Add flags to the root command
	lastnCmd.PersistentFlags().IntP(number, number_short, viper.GetInt(number), "Number of traps to fetch")
	// Bind the flags to the configuration parameters
	viper.BindPFlag(number, lastnCmd.PersistentFlags().Lookup(number))

}

func last_n(path string, n int) {
	arr := make([]m.StoredPacket, n)

	c.WithDatabase(path, func(d *badger.DB) error {
		return db.LastN(d, arr)
	})

	// TODO handle prefix better (store config)
	const prefix string = "snmp_"

	for i, spckt := range arr {
		pretty_k := u.PrettyPrintPrefixedKey([]byte(spckt.Key), len(prefix))
		fmt.Printf("Trap: %d (%s)-----\n", i, pretty_k)
		output_trap(&spckt.Packet)
	}

}

func output_trap(p *m.Packet) {
	fmt.Printf("community: %s\n", p.Community)
	for _, v := range p.Variables {
		u.PrintVars(v)
	}

}
