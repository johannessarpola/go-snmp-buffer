package snmp

import (
	"fmt"

	"github.com/dgraph-io/badger/v4"
	"github.com/johannessarpola/go-network-buffer/db"
	m "github.com/johannessarpola/go-network-buffer/models"
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

	db.WithDatabase(path, func(d *badger.DB) error {
		return db.LastN(d, arr)
	})

	// TODO handle prefix better (store config)
	const prefix string = "snmp_"

	for i, spckt := range arr {
		idx_arr := spckt.Key[len(prefix):]                                                                   // TODO Cleanup
		pretty_k := fmt.Sprintf("%s%d", string(spckt.Key[:len(prefix)]), u.ConvertToUint64([]byte(idx_arr))) // TODO Cleanup

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
